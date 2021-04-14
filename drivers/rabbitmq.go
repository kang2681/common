package drivers

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/kang2681/common/log"

	"github.com/streadway/amqp"
)

var reconnectWaitTime = 3 * time.Second
var consumeErrWaitTime = 5 * time.Second

type RabbitMQConfig struct {
	Host     string // host
	Username string // 用户名
	Password string // 密码
	VHost    string // vhost
}

type RabbitMQClient struct {
	config       *RabbitMQConfig
	conn         atomic.Value
	log          *log.Logger
	mx           sync.RWMutex
	reConnecting bool
}

func NewRabbitMQClient(config *RabbitMQConfig) (*RabbitMQClient, error) {
	client := &RabbitMQClient{
		config: config,
		log:    log.NewWithUUID().WithField("struct", "RabbitMQClient"),
	}
	if err := client.connect(); err != nil {
		client.log.Debugf("client connect error %s", err.Error())
		return nil, err
	}
	return client, nil
}

// rabbitmq 连接
func (c *RabbitMQClient) connect() error {
	addr := fmt.Sprintf("amqp://%s:%s@%s", c.config.Username, c.config.Password, c.config.Host)
	if c.config.VHost != "" {
		addr = fmt.Sprintf("%s/%s", addr, c.config.VHost)
	}
	c.log.Debugf(" connect addr :%s", addr)
	conn, err := amqp.Dial(addr)
	if err != nil {
		c.log.Debugf("connect addr: %s, error %s", addr, err.Error())
		return err
	}
	c.log.Debugf("connect addr: %s success", addr)
	c.conn.Store(conn)
	return nil
}

// get channel
func (c *RabbitMQClient) getChan() (*amqp.Channel, error) {
	conn := c.conn.Load().(*amqp.Connection)
	mqchan, err := conn.Channel()
	if err != nil {
		go c.reConnect(err)
		c.log.Errorf("get chan error: %s", err.Error())
		return nil, err
	}
	return mqchan, nil
}

// queue inspect
func (c *RabbitMQClient) QueueInspect(queueName string) (amqp.Queue, error) {
	mqchan, err := c.getChan()
	if err != nil {
		c.log.Errorf("get chan error:%s", err.Error())
		return amqp.Queue{}, err
	}
	defer mqchan.Close()
	return mqchan.QueueInspect(queueName)
}

// 发布消息
func (c *RabbitMQClient) Publish(exchangeName, routingKey, message string) error {
	mqchan, err := c.getChan()
	if err != nil {
		c.log.Errorf("get chan error:%s", err.Error())
		return err
	}
	defer mqchan.Close()
	mqMsg := amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(message),
		DeliveryMode: amqp.Persistent,
	}
	return mqchan.Publish(exchangeName, routingKey, false, false, mqMsg)
}

// 消费消息
func (c *RabbitMQClient) getConsume(queueName string, prefetchCount int) (*amqp.Channel, <-chan amqp.Delivery, error) {
	mqchan, err := c.getChan()
	if err != nil {
		c.log.Errorf("get chan error:%s", err.Error())
		return nil, nil, err
	}
	c.log.Debug("getConsume success")
	err = mqchan.Qos(prefetchCount, 0, false)
	if err != nil {
		mqchan.Close()
		c.log.Errorf("chan Qos error:%s", err.Error())
		return nil, nil, err
	}
	c.log.Debugf("getConsume Qos success")
	msgs, err := mqchan.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		mqchan.Close()
		c.log.Errorf("Consume error:%s", err.Error())
		return nil, nil, err
	}
	c.log.Debug("getConsume Consume success")
	return mqchan, msgs, nil
}

// consume message
func (c *RabbitMQClient) Consume(queueName string, prefetchCount int) <-chan amqp.Delivery {
	deliveryChan := make(chan amqp.Delivery)
	go func(deliveryChan chan amqp.Delivery) {
		for {
			// get consume
			mqchan, msgs, err := c.getConsume(queueName, prefetchCount)
			if err != nil {
				c.log.Errorf("get consume error:%s", err.Error())
				continue
			}
			// delivery message
			for v := range msgs {
				deliveryChan <- v
			}
			mqchan.Close()
		}
	}(deliveryChan)
	return (<-chan amqp.Delivery)(deliveryChan)
}

// set reconn flag
func (c *RabbitMQClient) LockReConn() bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	if c.reConnecting {
		return false
	}
	c.reConnecting = true
	return true
}

// get reconn flag
func (c *RabbitMQClient) UnlockReConn() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.reConnecting = false
}

// rabbit reconnect
func (c *RabbitMQClient) reConnect(err error) {
	// check err is connect close
	if err != amqp.ErrClosed {
		c.log.Debugf("skip reConnect error:%s", err.Error())
		return
	}
	// ensure one thread to visited
	if !c.LockReConn() {
		c.log.Debugf("RabbitMQ is reconning....")
		return
	}
	c.log.Debugf("start to RabbitMQ reconning")

	defer c.UnlockReConn()
	conn := c.conn.Load().(*amqp.Connection)
	conn.Close()
	for {
		c.log.Debugf("RabbitMQ is doing reconning....")
		err := c.connect()
		if err == nil {
			break
		}
		c.log.Debugf("reconnect mq error:%s", err.Error())
		time.Sleep(reconnectWaitTime)
	}
	c.log.Debugf("RabbitMQ reconning success")
}

// 关闭MQ连接
func (c *RabbitMQClient) Close() error {
	c.log.Debugf(" call close")
	connVal := c.conn.Load()
	if connVal == nil {
		return nil
	}
	conn := connVal.(*amqp.Connection)
	err := conn.Close()
	if err != nil {
		if err == amqp.ErrClosed {
			return nil
		}
		c.log.Debugf("close error %s", err.Error())
		return err
	}
	return nil
}

package drivers

import (
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	ip, user, pwd, vhost string
	conn                 *amqp.Connection
	log                  *logrus.Logger
	lock                 sync.RWMutex
	reConnFlagLock       sync.RWMutex
	isReConn             bool
}

var reconnectWaitTime = time.Second * 3
var consumeErrWaitTime = time.Second * 5

func NewRabbitMQConnect(ip, user, pwd, vhost string, log *logrus.Logger) (*RabbitMQ, error) {
	mq := &RabbitMQ{ip: ip, user: user, pwd: pwd, vhost: vhost, log: log}
	err := mq.rabbitMqConn()
	if err != nil {
		log.Errorf("<mq> rabbitMqConn error:%s", err.Error())
		return nil, err
	}
	return mq, nil
}

// rabbit mq connect
func (this *RabbitMQ) rabbitMqConn() error {
	if this.ip == "" || this.user == "" || this.pwd == "" {
		this.log.Error("<mq> rabbit connect param error,param:%+v", this)
		return fmt.Errorf("<mq> rabbit connect param error")
	}
	addr := "amqp://" + this.user + ":" + this.pwd + "@" + this.ip
	if this.vhost != "" {
		addr += "/" + this.vhost
	}
	this.log.Debugf("<mq> rabbitMqConn addr:%s", addr)
	conn, err := amqp.Dial(addr)
	if err != nil {
		this.log.Errorf("<mq> amqp dial error:%s,addr:%s", err.Error(), addr)
		return err
	}
	this.conn = conn
	return nil
}

// set reconn flag
func (this *RabbitMQ) setReConnFlag(flag bool) {
	this.reConnFlagLock.Lock()
	defer this.reConnFlagLock.Unlock()
	this.isReConn = flag
}

// get reconn flag
func (this *RabbitMQ) getReConnFlag() bool {
	this.reConnFlagLock.Lock()
	defer this.reConnFlagLock.Unlock()
	return this.isReConn
}

// rabbit reconnect
func (this *RabbitMQ) reConnect(err error) {
	// check err is connect close
	if err != amqp.ErrClosed {
		this.log.Debugf("<mq> skip reConnect error:%s not le closed error")
		return
	}
	// ensure one thread to visited
	if this.getReConnFlag() {
		this.log.Debugf("<mq> RabbitMQ is reconning")
		return
	}
	this.setReConnFlag(true)
	this.log.Debugf("<mq> start to RabbitMQ reconning")

	this.lock.Lock()
	defer this.lock.Unlock()
	this.conn.Close()
	for {
		err := this.rabbitMqConn()
		if err == nil {
			break
		}
		this.log.Debugf("<mq> reconnect mq error:%s", err.Error())
		time.Sleep(reconnectWaitTime)
	}
	this.setReConnFlag(false)
	this.log.Debugf("<mq> RabbitMQ reconning success")
}

// queue inspect
func (this *RabbitMQ) QueueInspect(queueName string) (amqp.Queue, error) {
	mqchan, err := this.getChan()
	if err != nil {
		this.log.Errorf("<mq> get chan error:%s", err.Error())
		return amqp.Queue{}, err
	}
	defer mqchan.Close()
	return mqchan.QueueInspect(queueName)
}

// publish message
func (this *RabbitMQ) PublishMsg(exchangeName, routingKey, message string) error {
	mqchan, err := this.getChan()
	if err != nil {
		this.log.Errorf("<mq> get chan error:%s", err.Error())
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

// consume message
func (this *RabbitMQ) Consume(queueName string, prefetchCount int) <-chan amqp.Delivery {
	deliveryChan := make(chan amqp.Delivery)
	go func(deliveryChan chan amqp.Delivery) {
		for {
			// check is reconnecting
			if this.getReConnFlag() {
				time.Sleep(consumeErrWaitTime)
				continue
			}
			// get consume
			mqchan, msgs, err := this.getConsume(queueName, prefetchCount)
			if err != nil {
				if mqchan != nil {
					mqchan.Close()
				}
				this.log.Errorf("<mq> get consume error:%s", err.Error())
				continue
			}
			// delivery message
			for v := range msgs {
				deliveryChan <- v
			}
		}
	}(deliveryChan)
	return (<-chan amqp.Delivery)(deliveryChan)
}

// get channel
func (this *RabbitMQ) getChan() (*amqp.Channel, error) {
	mqchan, err := this.conn.Channel()
	if err != nil {
		go this.reConnect(err)
		this.log.Errorf("<mq> get chan error:%s", err.Error())
		return nil, err
	}
	return mqchan, err
}

// get consume msg chan
func (this *RabbitMQ) getConsume(queueName string, prefetchCount int) (*amqp.Channel, <-chan amqp.Delivery, error) {
	mqchan, err := this.getChan()
	if err != nil {
		this.log.Errorf("<mq> get chan error:%s", err.Error())
		return nil, nil, err
	}
	this.log.Debug("<mq> getConsume success")
	err = mqchan.Qos(prefetchCount, 0, false)
	if err != nil {
		mqchan.Close()
		this.log.Errorf("<mq> chan Qos error:%s", err.Error())
		return nil, nil, err
	}
	this.log.Debugf("<mq> getConsume Qos success")
	msgs, err := mqchan.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		mqchan.Close()
		this.log.Errorf("<mq> Consume error:%s", err.Error())
		return nil, nil, err
	}
	this.log.Debug("<mq> getConsume Consume success")
	return mqchan, msgs, nil
}

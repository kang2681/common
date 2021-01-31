package sockets

import (
	"io"
	"net"
	"time"
)

//NewSocketConn 创建一个SOCKET连接，并发送content到目标 超时单位：秒
func NewSocketConn(server, content string, timeout int64) (string, int, error) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		return "", 0, err
	}
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	_, err = conn.Write([]byte(content))
	if err != nil {
		return "", 0, err
	}
	buf := make([]byte, 1024)
	result := ""
	resLen := 0
	for {
		n, errs := conn.Read(buf)
		if errs != nil {
			if errs == io.EOF {
				err = nil
			} else {
				err = errs
			}
			break
		}
		if n == 0 {
			break
		}
		if n < 1024 {
			buf = buf[0:n]
		}
		result += string(buf)
		resLen += n
	}
	return result, resLen, err
}

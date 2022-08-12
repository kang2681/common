package qywx

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRobot_Send(t *testing.T) {
	Convey("TestRobot_Send", t, func() {
		Convey("false", func() {
			r := NewRobot("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=111")
			msg := robotMessage{
				MsgType: "text",
				Text: TextMessage{
					Content: "text",
				},
			}
			r.Send(msg)
		})
	})
}

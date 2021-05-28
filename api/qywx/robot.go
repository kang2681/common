package qywx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/kang2681/common/httpext"
	"github.com/kang2681/common/log"
)

type Robot struct {
	log *log.Logger
	URL string
}

func NewRobot(l *log.Logger, url string) *Robot {
	return &Robot{log: l, URL: url}
}

type robotMessage struct {
	MsgType  string          `json:"msgtype"` // text markdown image news file
	Text     TextMessage     `json:"text,omitempty"`
	Markdown MarkdownMessage `json:"markdown"`
	Image    ImageMessage    `json:"image"`
	News     NewsMessage     `json:"news"`
}

// {"errcode":0,"errmsg":"ok"}
type robotResp struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MedisId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

func (r *Robot) Send(msg robotMessage) (*robotResp, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	req := httpext.PostRequest{
		Body: bytes.NewBuffer(data),
	}
	req.URL = r.URL
	resp, err := httpext.NewDefaultClient().PostRaw(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	rs := robotResp{}
	if err := resp.Json(&rs); err != nil {
		return nil, err
	}
	return &rs, nil
}

type TextMessage struct {
	Content             string   `json:"content"`                         // 文本内容，最长不超过2048个字节，必须是utf8编码
	MentionedList       []string `json:"mentioned_list,omitempty"`        // userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"` // 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
}

func (r *Robot) SendText(msg TextMessage) error {
	rs, err := r.Send(robotMessage{MsgType: "text", Text: msg})
	if err != nil {
		return err
	}
	if rs.ErrCode != 0 {
		return fmt.Errorf("%s", rs.ErrMsg)
	}
	return nil
}

type MarkdownMessage struct {
	Content string `json:"content"` // markdown内容，最长不超过4096个字节，必须是utf8编码
}

func (r *Robot) SendMarkdown(msg MarkdownMessage) error {
	rs, err := r.Send(robotMessage{MsgType: "markdown", Markdown: msg})
	if err != nil {
		return err
	}
	if rs.ErrCode != 0 {
		return fmt.Errorf("%s", rs.ErrMsg)
	}
	return nil
}

type ImageMessage struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

func (r *Robot) SendImage(msg ImageMessage) error {
	rs, err := r.Send(robotMessage{MsgType: "image", Image: msg})
	if err != nil {
		return err
	}
	if rs.ErrCode != 0 {
		return fmt.Errorf("%s", rs.ErrMsg)
	}
	return nil
}

// msgtype	是	消息类型，此时固定为image
// base64	是	图片内容的base64编码
// md5	是	图片内容（base64编码前）的md5值
type NewsMessage struct {
	Articles []NewsAtriclesMessage `json:"articles"` // 是	图文消息，一个图文消息支持1到8条图文
}

type NewsAtriclesMessage struct {
	Title       string `json:"title"`       // 是	标题，不超过128个字节，超过会自动截断
	Description string `json:"description"` // 否	描述，不超过512个字节，超过会自动截断
	Url         string `json:"url"`         // 是	点击后跳转的链接。
	Picurl      string `json:"picurl"`      // 否	图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150。
}

func (r *Robot) SendNews(msg NewsMessage) error {
	rs, err := r.Send(robotMessage{MsgType: "news", News: msg})
	if err != nil {
		return err
	}
	if rs.ErrCode != 0 {
		return fmt.Errorf("%s", rs.ErrMsg)
	}
	return nil
}

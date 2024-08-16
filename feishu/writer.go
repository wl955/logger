package feishu

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var Writer = &feishu{}

var client = &http.Client{}

type feishu struct {
	token string
}

func (o *feishu) With(token string) {
	o.token = token
}

func (o *feishu) Write(p []byte) (n int, e error) {
	type Content struct {
		Text string `json:"text"`
	}
	type Msg struct {
		MsgType string  `json:"msg_type"`
		Content Content `json:"content"`
	}
	msg := Msg{
		MsgType: "text",
		Content: Content{
			Text: string(p[:]),
		},
	}
	bs, _ := json.Marshal(msg)

	req, e := http.NewRequest("POST",
		"https://open.feishu.cn/open-apis/bot/v2/hook/"+o.token,
		bytes.NewReader(bs),
	)
	if e != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, e := client.Do(req)
	if e != nil {
		return
	}
	defer res.Body.Close()

	return
}

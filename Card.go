package GLarkBot

import (
	"github.com/xiehengjian/GRequests"
)

type Card struct {
	Config   Config
	Header   Header
	Elements []interface{}
}
type Config struct {
	WideScreenMode bool
}

type Header struct {
	Title    Text
	Template string
}

type Module struct {
}
type Text struct {
	Tag     string
	Content string
}
type Div struct {
	Tag    string
	Text   Text
	Fields Field
}
type Field struct {
	IsShort bool
	Text    Text
}

type Image struct {
}

type Action struct {
	Tag     string
	Actions []interface{}
}

type Button struct {
	Tag  string
	Text Text
	Url  string
}
type SelectMenu struct {
}

type OverFlow struct {
}

type DatePicker struct {
}

func (this *Bot) SendCardWithOpenID(OpenID string) {
	card := Card{
		Config: Config{WideScreenMode: false},
		Header: Header{Title: Text{
			Tag:     "plain_text",
			Content: "登录通知",
		}, Template: "wathet"},
		Elements: []interface{}{
			Div{Tag: "div", Text: Text{Tag: "plain_text", Content: "您好,该操作需要您登录系统,请您点击登录按钮进行登录"}},
			Button{
				Tag:  "button",
				Url:  "https://open.feishu.cn/open-apis/authen/v1/index?redirect_uri=http://127.0.0.1:5000/login&app_id=",
				Text: Text{Tag: "plain_text", Content: "登录"},
			},
		},
	}

	url := "https://open.feishu.cn/open-apis/message/v4/send/"
	data := map[string]interface{}{
		"open_id":OpenID,
		"card":card,
	}
	GRequests.Post(url,this.TenantAccessHeader,data)


}

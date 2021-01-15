package GLarkBot

import (
	"github.com/xiehengjian/GRequests"
)

type content struct {
	Text string
}
type SendText struct {
	Open_id  string
	Msg_type string
	Content  content
}

func (this *Bot) SendTextWithOpenID(openID string, text string) {
	url := "https://open.feishu.cn/open-apis/message/v4/send/"
	data := SendText{
		Open_id:  openID,
		Msg_type: "text",
		Content: content{
			Text: text,
		},
	}
	GRequests.Post(url,this.TenantAccessHeader,data)
}

func (this *Bot) SendTextWithUserID(userID string, text string) {
	url := "https://open.feishu.cn/open-apis/message/v4/send/"
	data :=map[string]interface{}{
		"user_id":userID,
		"msg_type":"text",
		"content":content{Text: text},
	}
	GRequests.Post(url,this.TenantAccessHeader,data)
}

func (this *Bot) SendTextWithChatID(chatID string, text string) {
	url := "https://open.feishu.cn/open-apis/message/v4/send/"
	data := map[string]interface{}{
		"chat_id":chatID,
		"msg_type":"text",
		"content":content{Text: text},
	}
	GRequests.Post(url,this.TenantAccessHeader,data)
}

func (this *Bot) ReplyTextWithChatID(chatID string, openMessageID string, text string) {
	url := "https://open.feishu.cn/open-apis/message/v4/send/"
	data := map[string]interface{}{
		"chat_id":chatID,
		"msg_type":"text",
		"root_id":openMessageID,
		"content": content{Text: text},
	}
	GRequests.Post(url,this.TenantAccessHeader,data)
}

package GLarkBot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	client := http.Client{}
	url := "https://open.feishu.cn/open-apis/message/v4/send/"
	data := SendText{
		Open_id:  openID,
		Msg_type: "text",
		Content: content{
			Text: text,
		},
	}
	bytesData, _ := json.Marshal(data)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	request.Header.Set("Authorization", "Bearer "+this.TenantAccessToken)
	request.Header.Set("Content-Type", "application/json")
	client.Do(request)
	response,_:=client.Do(request)
	bytes,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func (this *Bot) SendTextWithUserID(userID string, text string) {
	client := http.Client{}
	url := "https://open.feishu.cn/open-apis/message/v4/send/"
	data := make(map[string]interface{})
	data["user_id"] = userID
	data["msg_type"] = "text"
	data["content"] = content{Text: text}
	bytesData, _ := json.Marshal(data)
	fmt.Println(string(bytesData))
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	request.Header.Set("Authorization", "Bearer "+this.TenantAccessToken)
	request.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(request)
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func (this *Bot) SendTextWithChatID(chatID string, text string) {
	client := http.Client{}
	url := "https://open.feishu.cn/open-apis/message/v4/send/"
	data := make(map[string]interface{})
	data["chat_id"] = chatID
	data["msg_type"] = "text"
	data["content"] = content{Text: text}
	bytesData, _ := json.Marshal(data)
	fmt.Println(bytesData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	request.Header.Set("Authorization", "Bearer "+this.TenantAccessToken)
	request.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(request)
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func (this *Bot) ReplyTextWithChatID(chatID string, openMessageID string, text string) {
	client := http.Client{}
	url := "https://open.feishu.cn/open-apis/message/v4/send/"
	data := make(map[string]interface{})
	data["chat_id"] = chatID
	data["msg_type"] = "text"
	data["root_id"] = openMessageID
	data["content"] = content{Text: text}
	bytesData, _ := json.Marshal(data)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	request.Header.Set("Authorization", "Bearer "+this.TenantAccessToken)
	request.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(request)
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

package GLarkBot

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GroupInfoResponse struct {
	Code int
	Msg  string
	Data GroupInfoData
}
type GroupInfoData struct {
	Avatar        string
	Description   string
	Chat_id       string
	Name          string
	Owner_open_id string
	Owner_user_id string
}

type GroupListResponse struct {
	Code int
	Msg  string
	Data GroupListData
}
type GroupListData struct {
	Has_more   bool
	Page_token string
	Groups     []GroupInfoData
}

func (this *Bot) GetGroupInfo(ChatID string) GroupInfoData {
	client := http.Client{}
	url := "https://open.feishu.cn/open-apis/chat/v4/info?chat_id=" + ChatID
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer "+this.TenantAccessToken)
	response, _ := client.Do(request)
	body := GroupInfoResponse{}
	bytes, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(bytes, &body)
	return body.Data
}

func (this *Bot) GetGroupList() []GroupInfoData {
	client := http.Client{}
	url := "https://open.feishu.cn/open-apis/chat/v4/list"
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer "+this.TenantAccessToken)
	response, _ := client.Do(request)
	body := GroupListResponse{}
	bytes, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(bytes, &body)
	return body.Data.Groups

}

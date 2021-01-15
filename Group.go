package GLarkBot

import (
	"github.com/xiehengjian/GRequests"
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
	url := "https://open.feishu.cn/open-apis/chat/v4/info?chat_id=" + ChatID
	body := GroupInfoResponse{}
	GRequests.Unmarshal(GRequests.Get(url,this.TenantAccessHeader,nil).Body,&body)
	return body.Data
}

func (this *Bot) GetGroupList() []GroupInfoData {
	url := "https://open.feishu.cn/open-apis/chat/v4/list"
	body := GroupListResponse{}
	GRequests.Unmarshal(GRequests.Get(url,this.TenantAccessHeader,nil).Body,&body)
	return body.Data.Groups
}

package GLarkBot

import (
	"github.com/xiehengjian/GRequests"
)

type UserListResponse struct {
	Code int
	Data UserListResponseData
}
type UserListResponseData struct {
	Has_more bool
	Items    []UserInfo
}

type UserInfo struct {
	Union_id          string
	User_id           string
	Open_id           string
	Name              string
	En_name           string
	Email             string
	Mobile            string
	Mobile_visible    string
	Gender            string
	Avatar            AvatarInfo
	Status            UserStatus
	Department_ids    []string
	Leader_user_id    string
	City              string
	Country           string
	Work_station      string
	Join_time         int
	Is_tenant_manager bool
	Employee_no       string
	Employee_type     int
}

type AvatarInfo struct {
	Avatar_72     string
	Avatar_240    string
	Avatar_640    string
	Avatar_origin string
}

type UserStatus struct {
	Is_frozen    bool
	Is_resigned  bool
	Is_Activated bool
}

type UserInfoResponse struct {
	Code int
	Msg  string
	Data UserInfoData
}
type UserInfoData struct {
	User UserInfo
}

type GetUserIDInfoWithMobilesResponse struct {
	Code int
	Msg string
	Data GetUserIDInfoWithMobilesResponseData
}
type GetUserIDInfoWithMobilesResponseData struct {
	Mobile_users map[string][]UserIDInfoData
}
type UserIDInfoData struct {
	Open_id string
	User_id string
}
func (this *Bot) GetUserList() {
	url := "https://open.feishu.cn/open-apis/contact/v3/users"
	body:=make(map[string]interface{})
	GRequests.Unmarshal(GRequests.Post(url,this.TenantAccessHeader,nil).Body,&body)
}

func (this *Bot) GetUserInfoWithOpenID(openID string) UserInfo {
	url := "https://open.feishu.cn/open-apis/contact/v3/users/" + openID
	body := UserInfoResponse{}
	GRequests.Unmarshal(GRequests.Post(url,this.TenantAccessHeader,nil).Body,&body)
	return body.Data.User
}


func (this *Bot) GetUserIDInfoWithMobiles(mobiles string) UserIDInfoData{
	url:="https://open.feishu.cn/open-apis/user/v1/batch_get_id?mobiles="+mobiles
	body := GetUserIDInfoWithMobilesResponse{}
	GRequests.Unmarshal(GRequests.Post(url,this.TenantAccessHeader,nil).Body,&body)
	return body.Data.Mobile_users[mobiles][0]
}

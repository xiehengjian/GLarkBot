package GLarkBot

import (
	"fmt"
	"github.com/xiehengjian/GRequests"
)

type Tenant struct {
	Code int
	//Msg string
	Tenant_access_token string
	//Expire int
}
type data struct {
	Access_token  string
	Name          string
	Refresh_token string
}

type User struct {
	Code int
	Msg  string
	Data data
}

func (this *Bot) GetTenantAccessToken() {
	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
	//定义请求主题
	data := map[string]interface{}{
		"app_id":this.AppID,
		"app_secret": this.AppSecret,
	}
	header:=map[string]string{
		"Content-Type":"application/json",
	}
	body := Tenant{}
	GRequests.Unmarshal(GRequests.Post(url,header,data).Body,&body)
	this.TenantAccessToken = body.Tenant_access_token
}

func (this *Bot) GetUserAccessToken(code string) {
	url := "https://open.feishu.cn/open-apis/authen/v1/access_token"
	//定义请求主题
	data := map[string]interface{}{
		"app_access_token":this.TenantAccessToken,
		"grant_type":"authorization_code",
		"code":code,
	}
	header:=map[string]string{
		"Content-Type":"application/json",
	}
	body := User{}
	GRequests.Unmarshal(GRequests.Post(url,header,data).Body,&body)

	this.UserAccessToken = body.Data.Access_token
	this.RefreshToken = body.Data.Refresh_token
	//暂存UserToken
	this.SaveUserAccessToken()
	this.SaveRefreshToken()

}

func (this *Bot) RefreshUserAccessToken() {
	url := "https://open.feishu.cn/open-apis/authen/v1/refresh_access_token"
	//定义请求主题
	data := map[string]interface{}{
		"app_access_token":this.TenantAccessToken,
		"grant_type":"refresh_token",
		"refresh_token": this.RefreshToken,

	}
	header:=map[string]string{
		"Content-Type":"application/json",
	}
	body := User{}
	GRequests.Unmarshal(GRequests.Post(url,header,data).Body,&body)
	this.UserAccessToken = body.Data.Access_token
	this.RefreshToken = body.Data.Refresh_token
	//暂存UserToken
	//暂存UserToken
	this.SaveUserAccessToken()
	this.SaveRefreshToken()

}

func (this *Bot) GetAuthorization(redirectUri string ,openID string) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/authen/v1/index?redirect_uri=%s/login&app_id=%s",redirectUri,this.AppID)
	this.SendTextWithOpenID(openID, url)
}

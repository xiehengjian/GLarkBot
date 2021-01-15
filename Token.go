package GLarkBot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	//先创建一个客户端
	client := http.Client{}
	//定义请求地址
	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
	//定义请求主题
	data := make(map[string]interface{})
	data["app_id"] = this.AppID
	data["app_secret"] = this.AppSecret
	bytesData, _ := json.Marshal(data)
	//创建请求
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	//设置请求头
	request.Header.Set("Content-Type", "application/json")
	//执行请求
	r, _ := client.Do(request)
	body := Tenant{}
	bytes, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(bytes, &body)
	this.TenantAccessToken = body.Tenant_access_token

}

func (this *Bot) GetUserAccessToken(code string) {
	//先创建一个客户端
	client := http.Client{}
	//定义请求地址
	url := "https://open.feishu.cn/open-apis/authen/v1/access_token"
	//定义请求主题
	data := make(map[string]interface{})
	data["app_access_token"] = this.TenantAccessToken
	data["grant_type"] = "authorization_code"
	data["code"] = code
	bytesData, _ := json.Marshal(data)
	//创建请求
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	//设置请求头
	request.Header.Set("Content-Type", "application/json")
	//执行请求
	r, _ := client.Do(request)
	bytes, _ := ioutil.ReadAll(r.Body)
	body := User{}
	json.Unmarshal(bytes, &body)

	this.UserAccessToken = body.Data.Access_token
	this.RefreshToken = body.Data.Refresh_token
	//暂存UserToken
	this.SaveUserAccessToken()
	this.SaveRefreshToken()

}

func (this *Bot) RefreshUserAccessToken() {
	//先创建一个客户端
	client := http.Client{}
	//定义请求地址
	url := "https://open.feishu.cn/open-apis/authen/v1/refresh_access_token"
	//定义请求主题
	data := make(map[string]interface{})
	data["app_access_token"] = this.TenantAccessToken
	data["grant_type"] = "refresh_token"
	data["refresh_token"] = this.RefreshToken
	bytesData, _ := json.Marshal(data)
	//创建请求
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	//设置请求头
	request.Header.Set("Content-Type", "application/json")
	//执行请求
	r, _ := client.Do(request)
	bytes, _ := ioutil.ReadAll(r.Body)
	body := User{}
	json.Unmarshal(bytes, &body)

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

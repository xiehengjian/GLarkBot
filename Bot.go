package GLarkBot

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Bot struct {
	AppID             string
	AppSecret         string
	TenantAccessToken string
	UserAccessToken   string
	AppAccessToken    string
	RefreshToken      string
}

func NewBot(AppID string, AppSecret string) Bot {
	bot := Bot{AppID: AppID, AppSecret: AppSecret}
	bot.GetTenantAccessToken()
	bot.ReadUserAccessToken()
	bot.ReadRefreshToken()
	return bot
}

func (this *Bot) SaveUserAccessToken() {
	f, err := os.Create("UserAccessToken")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		f.Write([]byte(this.UserAccessToken))

	}
}

func (this *Bot) ReadUserAccessToken() {
	f, err := os.OpenFile("UserAccessToken", os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		contentByte, _ := ioutil.ReadAll(f)
		this.UserAccessToken = string(contentByte)
	}

}

func (this *Bot) SaveRefreshToken() {
	f, err := os.Create("RefreshToken")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		f.Write([]byte(this.RefreshToken))

	}
}

func (this *Bot) ReadRefreshToken() {
	f, err := os.OpenFile("RefreshToken", os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		contentByte, _ := ioutil.ReadAll(f)
		this.RefreshToken = string(contentByte)
	}

}

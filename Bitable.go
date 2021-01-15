package GLarkBot

import (
	"fmt"
	"github.com/xiehengjian/GRequests"
)

type Record struct {
	Fields map[string]string
	Id string
}

type CreateRecordResponse struct {
	Code int
	Data Record `json:"data"`
}

type GetRecordResponse struct {
	Code int
	Data Record
}
type Person struct {
	Id string `json:"id"`
}

type ListRecordsResponse struct {
	Code int
	Data ListRecordsResponseData


}
type ListRecordsResponseData struct {
	Has_more bool
	Items []Record
	Page_token string
}
func (this *Bot) CreateRecord(appToken string, tableID string, fields map[string]interface{}) CreateRecordResponse {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables/%s/records", appToken, tableID)
	data := map[string]map[string]interface{}{
		"fields":fields,
	}
	body := CreateRecordResponse{}
	GRequests.Unmarshal(GRequests.Post(url,this.UserAccessHeader,data).Body,&body)
	return body
}


func (this *Bot) ListRecords(appToken string, tableID string,viewID string)ListRecordsResponse{
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables/%s/records?view_id=%s", appToken, tableID,viewID)
	body := ListRecordsResponse{}
	GRequests.Unmarshal(GRequests.Get(url,this.UserAccessHeader,nil).Body,&body)
	return body
}

func (this *Bot) GetRecords(appToken string, tableID string,recordID string) GetRecordResponse{
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables/%s/records/%s", appToken, tableID,recordID)
	body := GetRecordResponse{}
	GRequests.Unmarshal(GRequests.Get(url,this.UserAccessHeader,nil).Body,&body)
	return body
}
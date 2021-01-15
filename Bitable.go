package GLarkBot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	client := http.Client{}
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables/%s/records", appToken, tableID)
	data := make(map[string]map[string]interface{})
	data["fields"] = fields
	bytesData, _ := json.Marshal(data)

	request, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	request.Header.Set("Authorization", "Bearer "+this.UserAccessToken)
	request.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(request)
	bytes, _ := ioutil.ReadAll(response.Body)

	body := CreateRecordResponse{}
	json.Unmarshal(bytes, &body)
	return body
}


func (this *Bot) ListRecords(appToken string, tableID string,viewID string)ListRecordsResponse{
	client := http.Client{}
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables/%s/records?view_id=%s", appToken, tableID,viewID)
	request, _ := http.NewRequest("GET", url,nil)
	request.Header.Set("Authorization", "Bearer "+this.UserAccessToken)
	request.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(request)
	bytes, _ := ioutil.ReadAll(response.Body)
	body := ListRecordsResponse{}
	json.Unmarshal(bytes, &body)
	//fmt.Println(string(bytes))
	return body
}

func (this *Bot) GetRecords(appToken string, tableID string,recordID string) GetRecordResponse{
	client := http.Client{}
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables/%s/records/%s", appToken, tableID,recordID)
	request, _ := http.NewRequest("GET", url,nil)
	request.Header.Set("Authorization", "Bearer "+this.UserAccessToken)
	request.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(request)
	bytes, _ := ioutil.ReadAll(response.Body)

	body := GetRecordResponse{}
	json.Unmarshal(bytes, &body)


	return body
}
package bulkvs2go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func NewClient(username string, password string) *BulkVS2GoClient {
	return &BulkVS2GoClient{
		APIURL:      "https://portal.bulkvs.com/api/v1.0/",
		APIUsername: username,
		APIPassword: password,
	}
}

func (c *BulkVS2GoClient) GetAccountDetail() (*AccountDetailResponse, error) {
	var ret AccountDetailResponse
	err := c.getJson("/accountDetail", &ret)
	return &ret, err
}

func (c *BulkVS2GoClient) GetTnRecord(r *GETTNRecordRequest) (*TNRecordResponse, error) {
	var ret TNRecordResponse
	err := c.getJson("/tnRecord", ret)
	return &ret, err
}

func (c *BulkVS2GoClient) GetTrunkGroups(tg string) (*TrunkGroupsResponse, error) {
	var ret TrunkGroupsResponse
	err := c.getJson(fmt.Sprintf("/trunkGroups%+v", url.Values{"TrunkGroup": {tg}}), ret)
	return &ret, err
}

func (c *BulkVS2GoClient) PostMessageSend(m *MessageSendRequest) (*MessageSendResponse, error) {
	var ret MessageSendResponse
	requestData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf("%+v/messageSend", c.APIURL), "application/json",bytes.NewBuffer(requestData) )
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(resp.Body).Decode(&ret)
	return &ret, err
}
func (c *BulkVS2GoClient) GetWebhooks(w string) (*WebhooksResponse, error) {
	var ret WebhooksResponse
	err := c.getJson("/webHooks", ret)
	return &ret, err
}

func (c *BulkVS2GoClient) getJson(endpoint string, target interface{}) error {
    r, err := http.Get(fmt.Sprintf("%+v/%+v", c.APIURL, endpoint))
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}
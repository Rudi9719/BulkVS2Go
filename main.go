package bulkvs2go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func NewClient(username string, password string) *BulkVS2GoClient {
	return &BulkVS2GoClient{
		APIURL:      "https://portal.bulkvs.com/api/v1.0",
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
	err = c.postJson("/messageSend", "application/json", bytes.NewBuffer(requestData), &ret)
	if err != nil {
		return nil, err
	}
	return &ret, err
}
func (c *BulkVS2GoClient) GetWebhooks(w string) (*WebhooksResponse, error) {
	var ret WebhooksResponse
	err := c.getJson("/webHooks", ret)
	return &ret, err
}

func (c *BulkVS2GoClient) postJson(endpoint string, contentType string, body io.Reader, target interface{}) error {
	client := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%+v%+v", c.APIURL, endpoint), body)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.APIUsername, c.APIPassword)
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (c *BulkVS2GoClient) getJson(endpoint string, target interface{}) error {
	client := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%+v%+v", c.APIURL, endpoint), http.NoBody)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.APIUsername, c.APIPassword)
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

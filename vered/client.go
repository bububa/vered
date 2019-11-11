package vered

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/afterpay/sdk/vered/util"
)

const GATEWAY = "http://114.242.193.198:50022/los/"

type Client struct {
	c          *http.Client
	publicKey  []byte
	privateKey []byte
	secret     []byte
	partnerNo  string
	channel    string
	mock       bool
}

func NewClient(publicKey string, privateKey string, secret string, partnerNo string, channel string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	pub, _ := base64.StdEncoding.DecodeString(publicKey)
	priv, _ := base64.StdEncoding.DecodeString(privateKey)
	secKey, _ := base64.StdEncoding.DecodeString(secret)

	return &Client{
		c:          client,
		publicKey:  pub,
		privateKey: priv,
		secret:     util.FixAesKeyLength(secKey),
		partnerNo:  partnerNo,
		channel:    channel,
	}
}

func (this *Client) SetMock(flag bool) {
	this.mock = flag
}

func (this *Client) Post(req Request) (json.RawMessage, error) {
	if req.GetPartnerNo() == "" {
		req.SetPartnerNo(this.partnerNo)
	}
	if req.GetChannel() == "" {
		req.SetChannel(this.channel)
	}
	signedReq, err := SignRequest(req, this.publicKey, this.privateKey, this.secret)
	if err != nil {
		return nil, err
	}
	var uri string
	switch req.(type) {
	case *DownloadFileRequest:
		uri = fmt.Sprintf("%svfc-intf-partner-gw.downloadFile", GATEWAY)
	default:
		if req.GetFileBytes() == nil {
			endpoint := "service"
			if this.mock {
				endpoint = "mock"
			}
			uri = fmt.Sprintf("%svfc-intf-partner-gw.%s", GATEWAY, endpoint)
		} else {
			uri = fmt.Sprintf("%svfc-intf-partner-gw.uploadFile", GATEWAY)
		}
	}

	buf, contentType, err := signedReq.Reader()
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST", uri, buf)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", contentType)
	resp, err := this.c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	var encryptedResp EncryptedResponse
	err = json.Unmarshal(body, &encryptedResp)
	if err != nil {
		return nil, err
	}
	if !encryptedResp.Success() {
		return nil, encryptedResp
	}
	data, err := encryptedResp.Decrypt(this.publicKey, this.secret)
	if err != nil {
		return nil, err
	}
	if req.GetFileBytes() != nil {
		return json.RawMessage(data), nil
	}
	var bizResponse BizResponse
	err = json.Unmarshal(data, &bizResponse)
	if err != nil {
		return nil, err
	}
	if !bizResponse.Success() {
		return nil, bizResponse
	}
	return bizResponse.Model, nil
}

func (this *Client) Upload(fileBytes []byte) (string, string, error) {
	req := &BaseRequest{
		FileBytes: fileBytes,
	}
	if req.GetPartnerNo() == "" {
		req.SetPartnerNo(this.partnerNo)
	}
	if req.GetChannel() == "" {
		req.SetChannel(this.channel)
	}
	signedReq, err := SignRequest(req, this.publicKey, this.privateKey, this.secret)
	if err != nil {
		return "", "", err
	}
	uri := fmt.Sprintf("%svfc-intf-partner-gw.uploadFile", GATEWAY)

	buf, contentType, err := signedReq.Reader()
	if err != nil {
		return "", contentType, err
	}
	httpReq, err := http.NewRequest("POST", uri, buf)
	if err != nil {
		return "", contentType, err
	}
	httpReq.Header.Set("Content-Type", contentType)
	resp, err := this.c.Do(httpReq)
	if err != nil {
		return "", contentType, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", contentType, err
	}
	fmt.Println(string(body))
	var encryptedResp EncryptedResponse
	err = json.Unmarshal(body, &encryptedResp)
	if err != nil {
		return "", contentType, err
	}
	if !encryptedResp.Success() {
		return "", contentType, encryptedResp
	}
	data, err := encryptedResp.Decrypt(this.publicKey, this.secret)
	if err != nil {
		return "", contentType, err
	}
	var bizResponse BizResponse
	err = json.Unmarshal(data, &bizResponse)
	if err != nil {
		return "", contentType, err
	}
	return bizResponse.GwFileKey, contentType, nil
}

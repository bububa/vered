package vered

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/afterpay/sdk/vered/internal/debug"
	"github.com/afterpay/sdk/vered/util"
)

type Client struct {
	c          *http.Client
	gateway    string
	publicKey  []byte
	privateKey []byte
	secret     []byte
	partnerNo  string
	channel    string
	mock       bool
}

func NewClient(gateway string, publicKey string, privateKey string, secret string, partnerNo string, channel string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	pub, _ := base64.StdEncoding.DecodeString(publicKey)
	priv, _ := base64.StdEncoding.DecodeString(privateKey)
	secKey, _ := base64.StdEncoding.DecodeString(secret)

	return &Client{
		c:          client,
		gateway:    gateway,
		publicKey:  pub,
		privateKey: priv,
		secret:     util.FixAesKeyLength(secKey),
		partnerNo:  partnerNo,
		channel:    channel,
	}
}

func NewClientWithCert(gateway string, publicKey string, certFile string, keyFile string, secret string, partnerNo string, channel string) (*Client, error) {
	client, err := NewTLSHttpClient(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	keyPEMBlock, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	keyDERBlock, _ := pem.Decode(keyPEMBlock)
	if keyDERBlock == nil {
		return nil, errors.New("missing key file")
	}
	pub, _ := base64.StdEncoding.DecodeString(publicKey)
	secKey, _ := base64.StdEncoding.DecodeString(secret)
	return &Client{
		c:          client,
		gateway:    gateway,
		publicKey:  pub,
		privateKey: keyDERBlock.Bytes,
		secret:     util.FixAesKeyLength(secKey),
		partnerNo:  partnerNo,
		channel:    channel,
	}, nil
}

func NewTLSHttpClient(certFile, keyFile string) (httpClient *http.Client, err error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	return newTLSHttpClient(tlsConfig)
}

func newTLSHttpClient(tlsConfig *tls.Config) (*http.Client, error) {
	dialTLS := func(network, addr string) (net.Conn, error) {
		return tls.DialWithDialer(&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}, network, addr, tlsConfig)
	}
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			DialTLS:               dialTLS,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}, nil
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
	reqBytes, _ := json.Marshal(req)
	signedReq, err := SignRequest(req, this.publicKey, this.privateKey, this.secret)
	if err != nil {
		return nil, err
	}
	var uri string
	switch req.(type) {
	case *DownloadFileRequest:
		uri = fmt.Sprintf("%svfc-intf-partner-gw.downloadFile", this.gateway)
	default:
		if req.GetFileBytes() == nil {
			endpoint := "service"
			if this.mock {
				endpoint = "mock"
			}
			uri = fmt.Sprintf("%svfc-intf-partner-gw.%s", this.gateway, endpoint)
		} else {
			uri = fmt.Sprintf("%svfc-intf-partner-gw.uploadFile", this.gateway)
		}
	}
	debug.DebugPrintPostJSONRequest(uri, reqBytes)
	dataBytes, buf, contentType, err := signedReq.Reader()
	if err != nil {
		return nil, err
	}
	debug.DebugPrintPostJSONRequest(uri, dataBytes)
	httpReq, err := http.NewRequest("POST", uri, buf)
	if err != nil {
		debug.DebugPrintError(err)
		return nil, err
	}
	httpReq.Header.Set("Content-Type", contentType)
	resp, err := this.c.Do(httpReq)
	if err != nil {
		debug.DebugPrintError(err)
		return nil, err
	}
	defer resp.Body.Close()
	var encryptedResp EncryptedResponse
	err = debug.DecodeJSONHttpResponse(resp.Body, &encryptedResp)
	if err != nil {
		debug.DebugPrintError(err)
		return nil, err
	}
	if !encryptedResp.Success() {
		debug.DebugPrintError(encryptedResp)
		return nil, encryptedResp
	}
	data, err := encryptedResp.Decrypt(this.publicKey, this.secret)
	if err != nil {
		debug.DebugPrintError(err)
		return nil, err
	}
	if req.GetFileBytes() != nil {
		return json.RawMessage(data), nil
	}
	var bizResponse BizResponse
	err = debug.DecodeJSONHttpResponse(bytes.NewReader(data), &bizResponse)
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
	uri := fmt.Sprintf("%svfc-intf-partner-gw.uploadFile", this.gateway)

	dataBytes, buf, contentType, err := signedReq.Reader()
	if err != nil {
		return "", contentType, err
	}
	debug.DebugPrintPostMultipartRequest(uri, dataBytes)
	httpReq, err := http.NewRequest("POST", uri, buf)
	if err != nil {
		return "", contentType, err
	}
	httpReq.Header.Set("Content-Type", contentType)
	resp, err := this.c.Do(httpReq)
	if err != nil {
		debug.DebugPrintError(err)
		return "", contentType, err
	}
	defer resp.Body.Close()
	var encryptedResp EncryptedResponse
	err = debug.DecodeJSONHttpResponse(resp.Body, &encryptedResp)
	if err != nil {
		debug.DebugPrintError(err)
		return "", contentType, err
	}
	if !encryptedResp.Success() {
		debug.DebugPrintError(encryptedResp)
		return "", contentType, encryptedResp
	}
	data, err := encryptedResp.Decrypt(this.publicKey, this.secret)
	if err != nil {
		debug.DebugPrintError(err)
		return "", contentType, err
	}
	var bizResponse BizResponse
	err = debug.DecodeJSONHttpResponse(bytes.NewReader(data), &bizResponse)
	if err != nil {
		debug.DebugPrintError(err)
		return "", contentType, err
	}
	return bizResponse.GwFileKey, contentType, nil
}

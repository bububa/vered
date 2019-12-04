package vered

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"github.com/afterpay/sdk/vered/util"
)

type SignedRequest struct {
	PartnerNo string            `json:"partnerNo"`         // 合作伙伴编号，由金融云分配
	EncKey    string            `json:"encKey"`            // 对称加密密钥的密文
	EncReq    string            `json:"encReq"`            // 业务报文的密文
	EncFile   map[string][]byte `json:"encFile,omitempty"` // 使用multipart，文件使用一个part，field名称为encFile
	ReqSign   string            `json:"reqSign"`           // 业务报文的签名
}

type Request interface {
	Method() string
	GetPartnerNo() string
	SetPartnerNo(string)
	GetChannel() string
	SetChannel(string)
	GetFileBytes() []byte
	GetGwTransCode() string
	SetGwTransCode(string)
	GetGwTransTime() string
	SetGwTransTime(string)
	SetGwFileHash(string)
	SetChannelJnlNo(string)
	SetChannelDate(string)
}

type BaseRequest struct {
	PartnerNo    string `json:"-"`                     // 合作伙伴编号，由金融云分配
	FileBytes    []byte `json:"-"`                     // 文件字节码
	Channel      string `json:"channel"`               // 渠道编号
	ChannelDate  string `json:"channelDate"`           // 渠道日期 yyyyMMdd
	ChannelJnlNo string `json:"channelJnlNo"`          // 渠道流水号
	GwTransCode  string `json:"gwTransCode,omitempty"` // 业务接口名称，由业务接口文档指定
	GwTransTime  string `json:"gwTransTime"`           // 交易发起时间，格式为yyyyMMdd-HHmmss-SSS 建议使用真实时间，服务端会检查时间合理性。
	GwFileHash   string `json:"gwFileHash,omitempty"`  // 文件明文的哈希的Base64字符串
}

func (this *BaseRequest) Method() string {
	return ""
}

func (this *BaseRequest) GetFileBytes() []byte {
	return this.FileBytes
}

func (this *BaseRequest) GetPartnerNo() string {
	return this.PartnerNo
}

func (this *BaseRequest) SetPartnerNo(str string) {
	this.PartnerNo = str
}

func (this *BaseRequest) GetChannel() string {
	return this.Channel
}

func (this *BaseRequest) SetChannel(channel string) {
	this.Channel = channel
}

func (this *BaseRequest) GetGwTransCode() string {
	return this.GwTransCode
}

func (this *BaseRequest) SetGwTransCode(code string) {
	this.GwTransCode = code
}

func (this *BaseRequest) GetGwTransTime() string {
	return this.GwTransTime
}

func (this *BaseRequest) SetGwTransTime(t string) {
	this.GwTransTime = t
}

func (this *BaseRequest) SetGwFileHash(h string) {
	this.GwFileHash = h
}

func (this *BaseRequest) SetChannelJnlNo(str string) {
	this.ChannelJnlNo = str
}

func (this *BaseRequest) SetChannelDate(date string) {
	this.ChannelDate = date
}

type DownloadFileRequest struct {
	BaseRequest
	Format     string `json:"gwDownloadFileFormat"`     // 文件格式
	FssReadKey string `json:"gwDownloadFileFssReadKey"` // 文件定位符
	Name       string `json:"gwDownloadFileName"`       // 文件名
}

func (this *DownloadFileRequest) Method() string {
	return this.GwTransCode
}

func SignRequest(req Request, publicKey []byte, privateKey []byte, secret []byte) (*SignedRequest, error) {
	req.SetGwTransTime(time.Now().Format("20060102-150405-000"))
	req.SetChannelDate(time.Now().Format("20060102"))
	jnlNo, _ := util.Salt()
	req.SetChannelJnlNo(jnlNo)
	fileBytes := req.GetFileBytes()
	if fileBytes == nil {
		req.SetGwTransCode(req.Method())
	} else {
		req.SetGwFileHash(util.Sha1Base64(fileBytes))
	}
	encKey, err := util.RSAEncryptBase64(publicKey, secret)
	if err != nil {
		return nil, err
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	encReq, err := util.AESEncryptBase64(secret, reqBytes)
	if err != nil {
		return nil, err
	}
	reqSign, err := util.RSASignWithSHA1Base64(privateKey, util.Sha1Sum(reqBytes))
	if err != nil {
		return nil, err
	}
	signedReq := &SignedRequest{
		PartnerNo: req.GetPartnerNo(),
		EncKey:    encKey,
		EncReq:    encReq,
		ReqSign:   reqSign,
	}
	if fileBytes != nil {
		encryptFile, err := util.AESEncrypt(secret, fileBytes)
		if err != nil {
			return nil, err
		}
		signedReq.EncFile = map[string][]byte{
			"fileMap": encryptFile,
		}
	}
	return signedReq, nil
}

func (this *SignedRequest) Reader() ([]byte, io.Reader, string, error) {
	if this.EncFile == nil {
		bytesData, err := json.Marshal(this)
		if err != nil {
			return nil, nil, "", err
		}
		contentType := "application/json;charset=UTF-8"
		return bytesData, bytes.NewReader(bytesData), contentType, nil
	}
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	w.WriteField("partnerNo", this.PartnerNo)
	w.WriteField("encKey", this.EncKey)
	w.WriteField("encReq", this.EncReq)
	w.WriteField("reqSign", this.ReqSign)
	fw, err := w.CreateFormFile("encFile", "file")
	if err != nil {
		return nil, nil, "", err
	}
	_, err = fw.Write(this.EncFile["fileMap"])
	if err != nil {
		return nil, nil, "", err
	}
	contentType := w.FormDataContentType()
	w.Close()
	return buf.Bytes(), buf, contentType, nil
}

type BizResponse struct {
	GwTransTime  string          `json:"gwTransTime"`           // 将请求方上送的时间字符串返回
	GwFileKey    string          `json:"gwFileKey,omitempty"`   // 文件读取授权码，小于64位
	ResponseCode string          `json:"responseCode"`          // 技术响应码
	ResponseMsg  string          `json:"responseMsg,omitempty"` // 技术响应消息
	Model        json.RawMessage `json:"model,omitempty`        // 业务报文的key值
}

func (this *BizResponse) Success() bool {
	return this.ResponseCode == "000000"
}

func (this BizResponse) Error() string {
	return fmt.Sprintf("CODE: %s, MESSAGE: %s", this.ResponseCode, this.ResponseMsg)
}

type EncryptedResponse struct {
	GwTransTime  string                  `json:"gwTransTime"`           // 将请求方上送的时间字符串返回
	GwFileKey    string                  `json:"gwFileKey,omitempty"`   // 文件读取授权码，小于64位
	ResponseCode string                  `json:"responseCode"`          // 技术响应码
	ResponseMsg  string                  `json:"responseMsg,omitempty"` // 技术响应消息
	Model        *EncryptedResponseModel `json:"model,omitempty"`       // 技术响应体
}

type EncryptedResponseModel struct {
	EncResp  string `json:"encResp"`  // 业务报文的密文
	RespSign string `json:"respSign"` // 业务报文的签名
}

func (this *EncryptedResponse) Success() bool {
	return this.ResponseCode == "000000"
}

func (this EncryptedResponse) Error() string {
	return fmt.Sprintf("CODE: %s, MESSAGE: %s", this.ResponseCode, this.ResponseMsg)
}

func (this *EncryptedResponse) Decrypt(publicKey []byte, secret []byte) ([]byte, error) {
	if !this.Success() {
		return nil, this
	}
	data, err := util.AESDecryptBase64(secret, this.Model.EncResp)
	if err != nil {
		return nil, err
	}
	err = util.RSAVerifySignWithSha1(publicKey, util.Sha1Sum(data), this.Model.RespSign)
	if err != nil {
		return nil, err
	}
	return data, nil
}

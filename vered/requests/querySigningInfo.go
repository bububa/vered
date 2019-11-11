package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type QuerySigningInfoRequest struct {
	vered.BaseRequest
	PmtCmpNo    string `json:"pmtCmpNo"`             // 支付公司编号, T003：宝付
	PrjNum      string `json:"prjNum"`               // 项目编号
	AcctNo      string `json:"signAcctNo"`           // 签约账号, 必填，1-32位
	BindPhone   string `json:"signBindPhone"`        // 绑定手机号, 必填，11位
	IdNo        string `json:"signIdNo"`             // 签约证件号码
	IdType      string `json:"signIdType"`           // 签约证件类型
	InnerBankNo string `json:"signInnerBankNo"`      // 签约蔷薇内部银行行号, 必填，7位，蔷薇内部银行行号
	PathNo      string `json:"signPathNo,omitempty"` // 签约通道编码, 需要发送短信签约填写, T003S00001：宝付签约；交通银行不支持协议支付，该字段不传
}

func (this *QuerySigningInfoRequest) Method() string {
	return "vfc-intf-partner-myy.qrySigningInfo"
}

type QuerySigningInfoResponse struct {
	BindingJnl string `json:"bindingJnl,omitempty"` // 签约协议编号, 交行宝付通道目前无该字段
	Reason     string `json:"failCause,omitempty"`  // 失败原因
	Status     string `json:"operStatus"`           // 签约状态, 非必填，（2-成功，3-失败）
	AcctNo     string `json:"signAcctNo"`           // 签约账号
}

func QuerySigningInfo(clt *vered.Client, req *QuerySigningInfoRequest) (*QuerySigningInfoResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &QuerySigningInfoResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

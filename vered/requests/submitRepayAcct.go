package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type SubmitRepayAcctRequest struct {
	vered.BaseRequest
	ApplicationNum string `json:"applicationNum"`  // 申请编号,不超过64位
	PmtCmpNo       string `json:"pmtCmpNo"`        // 支付公司编号, T003：宝付
	AcctNo         string `json:"signAcctNo"`      // 签约账号, 必填，1-32位
	BindPhone      string `json:"signBindPhone"`   // 绑定手机号, 必填，11位
	InnerBankNo    string `json:"signInnerBankNo"` // 签约蔷薇内部银行行号, 必填，7位，蔷薇内部银行行号
	OperToken      string `json:"signOperToken"`   // 签约操作令牌, 必填，签约验证需要，客户端保证唯一性
	PathNo         string `json:"signPathNo"`      // 签约通道编码, 需要发送短信签约填写
}

func (this *SubmitRepayAcctRequest) Method() string {
	return "vfc-intf-partner-myy.submitRepayAcct"
}

type SubmitRepayAcctResponse struct {
	Reason string `json:"failCause,omitempty"` // 失败原因
	Status string `json:"operStatus"`          // 签约状态, 非必填，（2-成功，3-失败）
	AcctNo string `json:"signAcctNo"`          // 签约账号
}

func SubmitRepayAcct(clt *vered.Client, req *SubmitRepayAcctRequest) (*SubmitRepayAcctResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &SubmitRepayAcctResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

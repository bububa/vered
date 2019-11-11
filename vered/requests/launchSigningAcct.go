package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type LaunchSigningAcctRequest struct {
	vered.BaseRequest
	ForceSignFlag string `json:"forceSignFlag,omitempty"` // 强制签约标识,默认不强制（1-是 0-否）
	PmtCmpNo      string `json:"pmtCmpNo"`                // 支付公司编号
	PrjNum        string `json:"prjNum"`                  // 项目编号
	AcctNo        string `json:"signAcctNo"`              // 签约账号, 必填，1-32位
	BindPhone     string `json:"signBindPhone"`           // 绑定手机号, 必填，11位
	IdNo          string `json:"signIdNo"`                // 签约证件号码
	IdType        string `json:"signIdType"`              // 签约证件类型
	InnerBankNo   string `json:"signInnerBankNo"`         // 签约蔷薇内部银行行号, 必填，7位，蔷薇内部银行行号
	Name          string `json:"signName"`                // 签约账户姓名
	OperToken     string `json:"signOperToken"`           // 签约操作令牌, 必填，签约验证需要，客户端保证唯一性
	PathNo        string `json:"signPathNo"`              // 签约通道编码, 需要发送短信签约填写
}

func (this *LaunchSigningAcctRequest) Method() string {
	return "vfc-intf-partner-myy.launchSigningAcct"
}

type LaunchSigningAcctResponse struct {
	Reason string `json:"failCause,omitempty"` // 失败原因
	Status string `json:"operStatus"`          // 签约状态, 非必填，（2-成功，3-失败）
	AcctNo string `json:"signAcctNo"`          // 签约账号
}

func LaunchSigningAcct(clt *vered.Client, req *LaunchSigningAcctRequest) (*LaunchSigningAcctResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &LaunchSigningAcctResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

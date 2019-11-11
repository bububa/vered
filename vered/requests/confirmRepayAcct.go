package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type ConfirmRepayAcctRequest struct {
	vered.BaseRequest
	ApplicationNum string `json:"applicationNum"`        // 申请编号
	ValidationCode string `json:"messageValidationCode"` // 短信验证码
	OperToken      string `json:"signOperToken"`         // 签约操作令牌, 使用“发起银行卡账户签约”的值
}

func (this *ConfirmRepayAcctRequest) Method() string {
	return "vfc-intf-partner-myy.confirmRepayAcct"
}

type ConfirmRepayAcctResponse struct {
	BindingJnl string `json:"bindingJnl"` // 签约协议编号
	Status     string `json:"operStatus"` // 签约状态, 非必填，（2-成功，3-失败）
	AcctNo     string `json:"signAcctNo"` // 签约账号
}

func ConfirmRepayAcct(clt *vered.Client, req *ConfirmRepayAcctRequest) (*ConfirmRepayAcctResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &ConfirmRepayAcctResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

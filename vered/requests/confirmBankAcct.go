package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type ConfirmBankAcctRequest struct {
	vered.BaseRequest
	ValidationCode string `json:"messageValidationCode"` // 短信验证码
	OperToken      string `json:"signOperToken"`         // 签约操作令牌, 使用“发起银行卡账户签约”的值
}

func (this *ConfirmBankAcctRequest) Method() string {
	return "vfc-intf-partner-myy.confirmBankAcct"
}

type ConfirmBankAcctResponse struct {
	BindingJnl string `json:"bindingJnl"` // 签约协议编号
	Status     string `json:"operStatus"` // 签约状态, 非必填，（2-成功，3-失败）
	AcctNo     string `json:"signAcctNo"` // 签约账号
}

func ConfirmBankAcct(clt *vered.Client, req *ConfirmBankAcctRequest) (*ConfirmBankAcctResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &ConfirmBankAcctResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

package requests

import (
	"github.com/afterpay/sdk/vered"
)

type QueryConfirmRepayAcctResultRequest struct {
	vered.BaseRequest
	ApplicationNum string `json:"applicationNum"` // 申请编号
	AcctNo         string `json:"signAcctNo"`     // 签约账号
}

func (this *QueryConfirmRepayAcctResultRequest) Method() string {
	return "vfc-intf-partner-myy.qryConfirmRepayAcctResult"
}

func QueryConfirmRepayAcctResult(clt *vered.Client, req *QueryConfirmRepayAcctResultRequest) error {
	_, err := clt.Post(req)
	return err
}

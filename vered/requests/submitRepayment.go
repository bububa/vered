package requests

import (
	"github.com/afterpay/sdk/vered"
)

type SubmitRepaymentRequest struct {
	vered.BaseRequest
	LoanNum   string `json:"loanNum"`   // 借据编号
	RepayMode string `json:"repayMode"` // 还款模式, 01：提前结清，03：强制提前结清
}

func (this *SubmitRepaymentRequest) Method() string {
	return "vfc-intf-partner-myy.submitRepaymentReq"
}

func SubmitRepayment(clt *vered.Client, req *SubmitRepaymentRequest) error {
	_, err := clt.Post(req)
	return err
}

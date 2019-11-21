package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type CalcRepaymentRequest struct {
	vered.BaseRequest
	LoanNum      string `json:"loanNum"`                // 借据编号
	RepayMode    string `json:"repayMode"`              // 还款模式, 01：提前结清，03：强制提前结清
	ScheduleDate string `json:"scheduleDate,omitempty"` // ignore
}

func (this *CalcRepaymentRequest) Method() string {
	return "vfc-intf-partner-myy.calcRepayment"
}

type CalcRepaymentResponse struct {
	Info      *RepaymentAmountInfo `json:"amountInfo"` // 金额信息
	LoanNum   string               `json:"loanNum"`    // 借据编号
	RepayMode string               `json:"repayMode"`  // 还款模式, 01：提前结清，03：强制提前结清
}

type RepaymentAmountInfo struct {
	ChargeAmt    float64 `json:"chargeAmt"`    // 费用金额
	InterestAmt  float64 `json:"interestAmt"`  // 利息金额
	PrincipalAmt float64 `json:"principalAmt"` // 本金金额
	TotalAmt     float64 `json:"totalAmt"`     // 总还款金额
}

func CalcRepayment(clt *vered.Client, req *CalcRepaymentRequest) (*CalcRepaymentResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &CalcRepaymentResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

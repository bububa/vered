package requests

import (
	"github.com/afterpay/sdk/vered"
)

type SubmitWithholdingRequest struct {
	vered.BaseRequest
	loanNum           string               `json:"loanNum"`                     // 借据编号
	PmtCmpNo          string               `json:"pmtCmpNo"`                    // 支付公司编号, T003：宝付
	Term              uint                 `json:"termNum"`                     // 还款期次
	WithHolding       *WithdrawHoldingInfo `json:"withholdingInfo,omitempty"`   // 划扣信息
	WithholdingMethod string               `json:"withholdingMethod,omitempty"` // 划扣分账方式
}

func (this *SubmitWithholdingRequest) Method() string {
	return "vfc-intf-partner-myy.submitWithholding"
}

type WithdrawHoldingInfo struct {
	ChnlAmt        float64 `json:"chnlAmt"`        // 渠道原始收款金额, 必输，渠道的实际收款金额为本金额减去渠道营销抵扣金额
	ChnlDeductAmt  float64 `json:"chnlDeductAmt"`  // 渠道营销抵扣金额
	CustInstallAmt float64 `json:"custInstallAmt"` // 客户期供还款金额
	FundAmt        float64 `json:"fundAmt"`        // 贷微赢原始收款金额
	FundDeductAmt  float64 `json:"fundDeductAmt"`  // 贷微赢营销抵扣金额
}

func SubmitWithholding(clt *vered.Client, req *SubmitWithholdingRequest) error {
	_, err := clt.Post(req)
	return err
}

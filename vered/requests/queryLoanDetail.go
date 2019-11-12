package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type ReceivableType = string

const (
	CHARGE                     ReceivableType = "CHARGE"
	INSTALLMENT                ReceivableType = "INSTALLMENT"
	EARLY_SETTLEMENT_PRINCIPAL ReceivableType = "EARLY_SETTLEMENT_PRINCIPAL"
)

type QueryLoanDetailRequest struct {
	vered.BaseRequest
	LoanNum string `json:"loanNum"` // 借据编号
}

func (this *QueryLoanDetailRequest) Method() string {
	return "vfc-intf-partner-myy.qryLoanDetail"
}

type QueryLoanDetailResponse struct {
	Info    *LoanDetailInfo `json:"basicInfo"`          // 基础借款信息
	Plans   []RepaymentPlan `json:"receivablePlanList"` // 还款计划列表
	Summary *LoanSummary    `json:"summary"`            // 借据概要信息
}

type LoanDetailInfo struct {
	IntFloatPercent       float64       `json:"intFloatPercent"`       // 利率浮动比例
	IntFloatType          string        `json:"intFloatType"`          // 利率浮动类型, 01浮动比例 02 浮动值 03非浮动利率
	IntFloatValue         float64       `json:"intFloatValue"`         // 利率浮动值
	InterestRate          float64       `json:"interestRate"`          // 借款利率
	PenaltyIntRate        float64       `json:"penaltyIntRate"`        // 罚息利率
	RateFloatTypeInd      string        `json:"rateFloatTypeInd"`      // 是否浮动利率, Y:是 / N:否
	LoanAmount            float64       `json:"loanAmount"`            // 借据金额
	RepaymentDate         string        `json:"repaymentDate"`         // 还款日
	RepaymentDateRuleCode RepayDateRule `json:"repaymentDateRuleCode"` // 还款日规则, TRANCHE_DAY：放款日对日 FIX_DAY：固定日历日
	RepaymentMethodCode   string        `json:"repaymentMethodCode"`   // 还款方式,  01: 等额本息(ANNUITY) 02: 等额本金(EQUAL_PRINCIPAL) 03: 利随本清(CLEAN_PRINCIPAL_AND_INTEREST) 04: 按期付息,到期还本(PAYMENT_AT_MONTH_INTEREST_DUE) 05: 等本等息(EQUAL_PRINCIPAL_AND_INTEREST)
	Term                  uint          `json:"term"`                  // 期限
	TermUnitCode          string        `json:"termUnitCode"`          // 期限单位, 01: 天(DAY) 02: 月(MONTH) 03: 年(YEAR)
	ValueDate             string        `json:"valueDate"`             // 起息日
	ValueDateRuleCode     RepayDateRule `json:"valueDateRuleCode"`     // 起息日规则
}

type RepaymentPlan struct {
	ExemptPenalty     float64        `json:"exemptPenalty"`     // 减免罚息
	GrossRcvbAmt      float64        `json:"grossRcvbAmt"`      // 应还总额
	OtsdPrincipal     float64        `json:"otsdPrincipal"`     // 剩余本金
	OutstandingAmt    float64        `json:"outstandingAmt"`    // 未还总额
	OverdueState      string         `json:"overdueState"`      // 逾期标志, N:正常 Y:逾期
	RcvbInterest      float64        `json:"rcvbInterest"`      // 应还利息
	RcvbPenalty       float64        `json:"rcvbPenalty"`       // 应还罚息
	RcvbPrincipal     float64        `json:"rcvbPrincipal"`     // 应还本金
	ReceivableDate    string         `json:"receivableDate"`    // 应还日期
	ReceivableType    ReceivableType `json:"receivableType"`    // 费用类型, 费用(CHARGE); 期供(INSTALLMENT);提前结清（EARLY_SETTLEMENT_PRINCIPAL）
	ReceivedAmt       float64        `json:"receivedAmt"`       // 已还总额
	ReceivedInterest  float64        `json:"receivedInterest"`  // 已还利息
	ReceivedPenalty   float64        `json:"receivedPenalty"`   // 已还罚息
	ReceivedPrincipal float64        `json:"receivedPrincipal"` // 已还本金
	RepaymentCd       string         `json:"repaymentCd"`       // 还款标志, NONE: 未收取; PART_COLL: 部分收取; ALL_RCVD:全部收取
	Term              uint           `json:"term"`              // 款项期次
}

type LoanSummary struct {
	ExpireDate     string     `json:"expireDate"`     // 借据到期日
	GrossRcvbAmt   float64    `json:"grossRcvbAmt"`   // 应还总额
	LoanNum        string     `json:"loanNum"`        // 借据编号
	OtsdInterest   float64    `json:"otsdInterest"`   // 未还利息
	OtsdPenalty    float64    `json:"otsdPenalty"`    // 未还罚息
	OtsdPrincipal  float64    `json:"otsdPrincipal"`  // 未还本金
	OutstandingAmt float64    `json:"outstandingAmt"` // 未还总额
	Period         uint       `json:"period"`         // 借款期限
	PeriodUnit     string     `json:"periodUnit"`     // 借款期限单位, 01:天(DAY)；02:月(MONTH)；03:年(YEAR)
	RcvbInterest   float64    `json:"rcvbInterest"`   // 应还利息
	RcvbPenalty    float64    `json:"rcvbPenalty"`    // 应还罚息
	RcvbPrincipal  float64    `json:"rcvbPrincipal"`  // 应还本金
	StartDate      string     `json:"startDate"`      // 借据起始日
	Status         LoanStatus `json:"status"`         // 借据状态
	TrancheAmt     float64    `json:"trancheAmt"`     // 放款金额
}

func QueryLoanDetail(clt *vered.Client, req *QueryLoanDetailRequest) (*QueryLoanDetailResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &QueryLoanDetailResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

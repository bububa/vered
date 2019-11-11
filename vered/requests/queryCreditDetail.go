package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type RepayDateRule = string

const (
	TRANCHE_DAY RepayDateRule = "TRANCHE_DAY"
	FIX_DAY     RepayDateRule = "FIX_DAY"
)

type CreditState = string

const (
	CREDIT_WAIT_EFF CreditState = "CREDIT_WAIT_EFF"
	CREDIT_NORMAL   CreditState = "CREDIT_NORMAL"
	CREDIT_FROZEN   CreditState = "CREDIT_FROZEN"
	CREDIT_EXPIRE   CreditState = "CREDIT_EXPIRE"
	CRT_GRANTED     CreditState = "CRT_GRANTED"
)

type QueryCreditDetailRequest struct {
	vered.BaseRequest
	ApplicationNum string `json:"applicationNum"` // 申请编号
}

func (this *QueryCreditDetailRequest) Method() string {
	return "vfc-intf-partner-myy.qryCreditDetail"
}

type QueryCreditDetailResponse struct {
	Info *CreditInfo `json:"basicInfo,omitempty"`            // 额度基本信息
	Meta *CreditMeta `json:"creditDetailMatadata,omitempty"` // 额度信息
}

type CreditInfo struct {
	ApplyCreditAmt        float64       `json:"applyCreditAmt"`        // 申请额度金额
	BaseRate              float64       `json:"baseRate"`              // 基准利率
	BaseRateType          string        `json:"baseRateType"`          // 基准利率类型, 02 固定利率
	CreditTerm            uint          `json:"creditTerm"`            // 额度期限
	CreditTermUnit        uint          `json:"creditTermUnit"`        // 额度期限单位, 01:天(DAY)；02:月(MONTH)；03:年(YEAR)
	IntCalMode            string        `json:"drawRpyMthdType"`       // 还款方式, 必输 01 等额本息 02 等额本金 03 一次还本付息 04 按期付息到期还本 05 等本等息
	ExecutionRate         float64       `json:"executionRate"`         // 借款利率
	GrantDate             string        `json:"grantDate"`             // 发放日期
	GuaranteeType         string        `json:"guaranteeType"`         // 担保方式, 01:保证, 02:抵押, 03:质押, 04:信用
	IntFloatPercent       float64       `json:"intFloatPercent"`       // 利率浮动比例
	IntFloatType          string        `json:"intFloatType"`          // 利率浮动类型, 01浮动比例 02 浮动值 03非浮动利率
	IntFloatValue         float64       `json:"intFloatValue"`         // 利率浮动值
	PnlRate               float64       `json:"pnlRate"`               // 罚息利率
	RateFloatTypeInd      string        `json:"rateFloatTypeInd"`      // 是否浮动利率, Y:是 / N:否
	RepaymentDateRuleCode RepayDateRule `json:"repaymentDateRuleCode"` // 还款日规则, TRANCHE_DAY：放款日对日 FIX_DAY：固定日历日
}

type CreditMeta struct {
	AvailableAmt float64     `json:"availableAmt"` // 可用额度
	CreditNum    string      `json:"creditNum"`    // 额度编号
	State        CreditState `json:"creditState"`  // 额度状态, CREDIT_WAIT_EFF(待发放), CREDIT_NORMAL(正常), CREDIT_FROZEN(冻结), CREDIT_EXPIRE(已失效),@mock=CRT_GRANTED
	EffectDate   string      `json:"effectDate"`   // 额度生效时间
	ExpiredDate  string      `json:"expiredDate"`  // 额度失效时间
	TotalAmt     float64     `json:"totalAmt"`     // 额度总金额
}

func QueryCreditDetail(clt *vered.Client, req *QueryCreditDetailRequest) (*QueryCreditDetailResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &QueryCreditDetailResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

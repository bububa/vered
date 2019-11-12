package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type ApplyDrawingCreditRequest struct {
	vered.BaseRequest
	ApplicationNum string            `json:"applicationNum"`     // 申请编号
	CreditNum      string            `json:"creditNum"`          // 额度编号
	Ext            map[string]string `json:"extInfo,omitempty"`  // 扩展字段
	Files          []File            `json:"fileList,omitempty"` // 文件列表
	Loan           *LoanBasicInfo2   `json:"loanBasicInfo"`      // 借款基本信息
}

type LoanBasicInfo2 struct {
	Amount              float64 `json:"imumAmount,omitempty"`
	IntCalMode          string  `json:"intCalMode"`             // 还款方式, 必输 01 等额本息 02 等额本金 03 一次还本付息 04 按期付息到期还本 05 等本等息
	ProductDeadline     int     `json:"productDeadline"`        // 借款期限
	ProductDeadlineUnit string  `json:"productDeadlineUnit"`    // 借款期限单位, 必输 01 天 02 月 03 年
	PurposeCode         string  `json:"purposeCode"`            // 借款用途, 必输，00消费类
	PurposeDetail       string  `json:"purposeDetail"`          // 借款用途详情
	IntRateFloat        float64 `json:"intRateFloat,omitempty"` // 浮动值, BigDecimal 非必输；单位：% ；精度：（9，4）
	IntRateIdenf        float64 `json:"intRateIdenf,omitempty"` // 借款利率, BigDecimal 必输；单位：% ；精度：（9，4）
	IntRateRatio        float64 `json:"intRateRatio,omitempty"` // 浮动比例, BigDecimal 非必输 ；单位：% ；精度：（9，4）
	RepaymentDay        uint    `json:"repaymentDay,omitempty"` // 还款日
}

func (this *ApplyDrawingCreditRequest) Method() string {
	return "vfc-intf-partner-myy.applyDrawingCredit"
}

type ApplyDrawingCreditResponse struct {
	Status     string `json:"status"`              // 申请结果状态, 00：成功，01：失败；02：处理中
	StatusCode string `json:"statusCode"`          // 01失败：错误码值不一一描述；02处理中：0201-进件审批中 0202-合同签署中 0203-客户签署完成，待发放；
	StatusMsg  string `json:"statusMsg,omitempty"` // 申请结果描述
}

func ApplyDrawingCredit(clt *vered.Client, req *ApplyDrawingCreditRequest) (*ApplyDrawingCreditResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &ApplyDrawingCreditResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

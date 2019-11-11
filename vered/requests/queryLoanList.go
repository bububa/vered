package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type LoanStatus = string

const (
	LOAN_NORMAL     LoanStatus = "LOAN_NORMAL"
	LOAN_OVER_DUE   LoanStatus = "LOAN_OVER_DUE"
	LOAN_SETTLEMENT LoanStatus = "LOAN_SETTLEMENT"
)

type QueryLoanListRequest struct {
	vered.BaseRequest
	ApplicationNumList []string `json:"applicationNumList,omitempty"` // 申请编号列表, 与借据编号列表二选一，优先使用借据编号列表
	loanNumList        []string `json:"loanNumList,omitempty"`        // 借据编号列表,与申请编号列表二选一，优先使用借据编号列表
}

func (this *QueryLoanListRequest) Method() string {
	return "vfc-intf-partner-myy.qryLoanList"
}

type QueryLoanListResponse struct {
	List []LoanInfo `json:"loanList"` // 借据列表
}

type LoanInfo struct {
	ApplicationNum string     `json:"applicationNum"` // 申请编号
	DueDate        string     `json:"dueDate"`        // 借据到期日
	LoanAmount     float64    `json:"loanAmount"`     // 借据金额
	LoanNum        string     `json:"loanNum"`        // 借据编号
	StartDate      string     `json:"startDate"`      // 借据开始日
	Status         LoanStatus `json:"status"`         // 借据状态, LOAN_NORMAL:正常；LOAN_OVER_DUE:逾期；LOAN_SETTLEMENT:结清
}

func QueryLoanList(clt *vered.Client, req *QueryLoanListRequest) (*QueryLoanListResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &QueryLoanListResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

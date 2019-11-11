package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type QueryRepaymentPlanRequest struct {
	vered.BaseRequest
	loanNum string `json:"loanNum"` // 借据编号
}

func (this *QueryRepaymentPlanRequest) Method() string {
	return "vfc-intf-partner-myy.qryRepaymentPlan"
}

type QueryRepaymentPlanResponse struct {
	Plans []RepaymentPlan `json:"receivablePlanList"` // 还款计划列表
}

func QueryRepaymentPlan(clt *vered.Client, req *QueryRepaymentPlanRequest) (*QueryRepaymentPlanResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &QueryRepaymentPlanResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type QueryDrawingResultRequest struct {
	vered.BaseRequest
	ApplicationNum string `json:"applicationNum"` // 申请编号
}

func (this *QueryDrawingResultRequest) Method() string {
	return "vfc-intf-partner-myy.qryDrawingReqResult"
}

type QueryDrawingResultResponse struct {
	LoanNum     string `json:"loanNum"`               // 借据编号
	Status      string `json:"status"`                // 申请结果状态, 00：成功，01：失败；02：处理中
	StatusCode  string `json:"statusCode"`            // 消息码值, 01失败：错误码值不一一描述；02处理中：0201-进件审批中 0202-合同签署中 0203-客户签署完成，待发放；
	StatusMsg   string `json:"statusMsg,omitempty"`   // 申请结果描述
	TrancheDate string `json:"trancheDate,omitempty"` // 额度发放日期
	TransDate   string `json:"transDate,omitempty"`   // 交易日期
}

func QueryDrawingResult(clt *vered.Client, req *QueryDrawingResultRequest) (*QueryDrawingResultResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &QueryDrawingResultResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

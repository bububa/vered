package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type QueryRepaymentResultRequest struct {
	vered.BaseRequest
	OrgChannelJnlNo string `json:"orgChannelJnlNo"` // 原渠道流水号
}

func (this *QueryRepaymentResultRequest) Method() string {
	return "vfc-intf-partner-myy.qryRepaymentResult"
}

type QueryRepaymentResultResponse struct {
	Status            string `json:"status"`                      // 申请结果状态, 00：成功，01：失败；02：处理中
	StatusCode        string `json:"statusCode"`                  // 消息码值, 01失败：错误码值不一一描述；02处理中：0201-进件审批中 0202-合同签署中 0203-客户签署完成，待发放；
	StatusMsg         string `json:"statusMsg,omitempty"`         // 申请结果描述
	TransCompleteTime string `json:"transCompleteTime,omitempty"` // 划扣成功时间
}

func QueryRepaymentResult(clt *vered.Client, req *QueryRepaymentResultRequest) (*QueryRepaymentResultResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &QueryRepaymentResultResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

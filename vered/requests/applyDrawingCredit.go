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
	Loan           *LoanBasicInfo    `json:"loanBasicInfo"`      // 借款基本信息
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

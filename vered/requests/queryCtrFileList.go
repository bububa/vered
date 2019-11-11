package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type QueryCtrFileListRequest struct {
	vered.BaseRequest
	loanNum string `json:"loanNum"` // 借据编号
}

func (this *QueryCtrFileListRequest) Method() string {
	return "vfc-intf-partner-myy.qryCtrFileList"
}

type QueryCtrFileListResponse struct {
	Id    string    `json:"ctrId"`       // 合同编号, 唯一标识一套合同
	Files []CtrFile `json:"ctrFileList"` // 合同文件列表
}

type CtrFile struct {
	Format       string `json:"ctrFileFormat"`          // 文件格式类型
	Id           string `json:"ctrFileId"`              // 合同文件编号, 合同文件编号，可用于电子签章签署、查询、下载
	Name         string `json:"ctrFileName"`            // 合同文件名称
	TplId        string `json:"ctrFileTplId,omitempty"` // 合同文件模板ID, 唯一标识一套合同中某一份合同，字典由蔷薇提供
	ElecSignFlag string `json:"elecSignFlag"`           // 是否电子签章标志, 0否1是
}

func QueryCtrFileList(clt *vered.Client, req *QueryCtrFileListRequest) (*QueryCtrFileListResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &QueryCtrFileListResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

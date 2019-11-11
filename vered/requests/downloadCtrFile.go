package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type DownloadCtrFileRequest struct {
	vered.BaseRequest
	Id   string  `json:"ctrId"`   // 合同编号, 唯一标识一套合同
	File CtrFile `json:"ctrFile"` // 合同文件
}

func (this *DownloadCtrFileRequest) Method() string {
	return "vfc-intf-partner-myy.downloadCtrFile"
}

type DownloadCtrFileResponse struct {
	Format     string `json:"gwDownloadFileFormat"`     // 文件格式
	FssReadKey string `json:"gwDownloadFileFssReadKey"` // 文件定位符
	Name       string `json:"gwDownloadFileName"`       // 文件名
}

func DownloadCtrFile(clt *vered.Client, req *DownloadCtrFileRequest) ([]byte, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &DownloadCtrFileResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	downloadFileReq := &vered.DownloadFileRequest{
		BaseRequest: vered.BaseRequest{
			PartnerNo:   req.PartnerNo,
			Channel:     req.Channel,
			GwTransCode: req.GwTransCode,
		},
		Format:     resp.Format,
		FssReadKey: resp.FssReadKey,
		Name:       resp.Name,
	}
	fileBytes, err := clt.Post(downloadFileReq)
	if err != nil {
		return nil, err
	}
	return fileBytes, nil
}

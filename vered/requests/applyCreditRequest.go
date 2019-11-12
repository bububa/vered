package requests

import (
	"encoding/json"

	"github.com/afterpay/sdk/vered"
)

type Gender = string

const (
	MALE   Gender = "M"
	FEMALE Gender = "F"
)

type ApplyCreditRequest struct {
	vered.BaseRequest
	ApplicationNum string            `json:"applicationNum"`     //申请编号
	Customer       *CustomerInfo     `json:"customerInfo"`       // 基本信息
	Ext            map[string]string `json:"extInfo,omitempty"`  // 扩展字段
	Files          []File            `json:"fileList,omitempty"` // 文件列表
	Loan           *LoanBasicInfo    `json:"loanBasicInfo"`      // 借款基本信息
	PrjNum         string            `json:"prjNum"`             // 项目编号
}

type CustomerInfo struct {
	Np            *NpInfo      `json:"npInfo"`              // 自然人信息
	OtherNps      []ContactNp  `json:"npOtherInfoUnitList"` // 其他联系人
	ReceptAccount *BankAccount `json:"receiptAccountInfo"`  // 收款账户
	RepayAccount  *BankAccount `json:"repayAccountInfo"`    // 还款账户
}

type NpInfo struct {
	BaseInfo *BaseInfo   `json:"baseInfo"`           // 基本信息
	Detail   *DetailInfo `json:"detailInfo"`         // 详细信息
	Work     *WorkInfo   `json:"workInfo,omitempty"` // 工作信息
}

type BaseInfo struct {
	Birth  string `json:"birth,omitempty"` // 生日
	Name   string `json:"custName"`        // 客户名称
	Gender Gender `json:"gender"`          // 性别
	IdNo   string `json:"idNo"`            // 证件号码
	IdType string `json:"idType"`          // 证件类型 00：身份证
	Phone  string `json:"phone"`           // 联系电话
}

type DetailInfo struct {
	ChildNum               uint    `json:"childNum,omitempty"`           // 子女个数
	CommunAddress          string  `json:"communAddress"`                // 通讯详细地址
	CommunArea             string  `json:"communArea"`                   // 通讯地址所在区
	CommunCity             string  `json:"communCity"`                   // 通讯地址所在市
	CommunPro              string  `json:"communPro"`                    // 通讯地址所在省
	CommunZip              string  `json:"communZip,omitempty"`          // 通讯地址邮编
	CredentialsInvalidDate string  `json:"credentialsInvalidDate"`       // 证件失效日期
	CredentialsIssueDate   string  `json:"credentialsIssueDate"`         // 证件签发日期
	IdAddress              string  `json:"idAddress"`                    // 身份证地址
	IdArea                 string  `json:"idArea"`                       // 身份证所在区
	IdCity                 string  `json:"idCity"`                       // 身份证所在市
	IdProvince             string  `json:"idProvince"`                   // 身份证所在省
	IssuingOrgan           string  `json:"issuingOrgan"`                 // 发证机关
	Phone                  string  `json:"phone"`                        // 联系电话
	CustCategory           string  `json:"custCategory,omitempty"`       // 客户类别, 非必输 00：农村客户 01：城市客户
	CustType               string  `json:"custType,omitempty"`           // 客户类型, 非必输，00：领薪人士 01：一般受薪人士 02：优良职业 03：公务员 04:事业单位员工 05；金融机构员工 06：部队中高级干部 07：自雇人士 08：个体工商户 09：私营企业主 10：自由职业 11：小企业实际控制人 12：农户 13：学生 14：其他 15：未知
	Email                  string  `json:"email,omitempty"`              // 邮箱
	FamliyMonthOutcome     float64 `json:"famliyMonthOutcome,omitempty"` // 每月家庭支出
	FamliyStructure        string  `json:"famliyStructure,omitempty"`    // 家庭结构, 非必输00有未成年子女 01 有子女都成年 02 无子女
	GraduationSchool       string  `json:"graduationSchool,omitempty"`   // 毕业学校
	GraduationTime         string  `json:"graduationTime,omitempty"`     // 毕业时间, 非必输yyyy-MM-dd
	HighestDegree          string  `json:"highestDegree,omitempty"`      // 最高学位, 非必输 00 学士 01 硕士 02 博士 03 其他
	HighestEducation       string  `json:"highestEducation,omitempty"`   // 最高学历, 非必输 00 博士及以上 01 研究生 02 大学本科 03 大学专科 04 中等职业教育 05 普通高级中学教育 06 初级中学教育 07 小学教育
	HouseAddress           string  `json:"houseAddress,omitempty"`       // 住宅地址
	HukouType              string  `json:"hukouType,omitempty"`          // 户口类型, 非必输 00：城镇户口 01：农业户口
	IsHasChildren          string  `json:"isHasChildren,omitempty"`      // 是否有子女
	LiveAddress            string  `json:"liveAddress,omitempty"`        // 居住地详细地址
	LiveAddressZip         string  `json:"liveAddressZip,omitempty"`     // 居住地邮编
	LiveArea               string  `json:"liveArea,omitempty"`           // 居住地所在区
	LiveCity               string  `json:"liveCity,omitempty"`           // 居住地所在市
	LiveProvince           string  `json:"liveProvince,omitempty"`       // 居住地所在省
	LiveStatus             string  `json:"liveStatus,omitempty"`         // 居住状况, 非必输 00：租房 01：本地有住无贷款) 02:本地有住无贷款(有贷款) 03:异地自有住房(无贷款) 04:异地自有住房(有贷款) 05:无自有住房
	MarriageStatus         string  `json:"marriageStatus,omitempty"`     // 婚姻状况, 非必输 01 未婚 00 已婚 02 离异 03 丧偶
	MonthIncome            float64 `json:"monthIncome,omitempty"`        // 月收入
	SpouseIdNo             string  `json:"spouseIdNo,omitempty"`         // 配偶证件号码
	SpouseIdType           string  `json:"spouseIdType,omitempty"`       // 配偶证件类型
	SpouseName             string  `json:"spouseName,omitempty"`         // 配偶姓名
	SupportFamilyNum       uint    `json:"supportFamilyNum,omitempty"`   // 供养亲属人数
	QQ                     string  `json:"qq,omitempty"`                 // QQ
	Wx                     string  `json:"wx,omitempty"`                 // 微信
}

type WorkInfo struct {
	WorkFlag string `json:"work_flag,omitempty"` // 供职情况, 00：受薪 01：自雇 02 ：无业/失业
}

type ContactNp struct {
	Relation string `json:"contactRelation"` // 联系人与申请人关系, 必输，{ '父母': '00' },{ '子女': '01' },{ '祖父母': '02' },{ '亲属',: '03' },{ '朋友',: '04' },{:'同事': '05' },{:'同学': '06' },{'其他': '07' }
	Phone    string `json:"phone"`           // 手机号
	Name     string `json:"relName"`         // 姓名
}

type BankAccount struct {
	Name     string `json:"acName"`          // 账户户名
	Num      string `json:"acNo"`            // 账户号
	OpenBank string `json:"openAccountBank"` // 开户行
	Phone    string `json:"phoneno"`         // 银行预留手机号
}

type File struct {
	Format string `json:"fileFormat"` // 文件内容类型, 必输，与http中ContentType保持一致
	Key    string `json:"fileKey"`    // 文件的唯一key值
	Name   string `json:"fileName"`   // 文件名称
	Type   string `json:"fileType"`   // 文件类型
}

type LoanBasicInfo struct {
	Amount              float64 `json:"imumAmount,omitempty"`
	GuaType             string  `json:"guaType,omitempty"`         // 担保方式, 必输 04信用
	IntCalMode          string  `json:"intCalMode"`                // 还款方式, 必输 01 等额本息 02 等额本金 03 一次还本付息 04 按期付息到期还本 05 等本等息
	ProductDeadline     string  `json:"productDeadline,omitempty"` // 借款期限
	ProductDeadlineInt  int     `json:"productDeadline,omitempty"` // 借款期限
	ProductDeadlineUnit string  `json:"productDeadlineUnit"`       // 借款期限单位, 必输 01 天 02 月 03 年
	PurposeCode         string  `json:"purposeCode"`               // 借款用途, 必输，00消费类
	PurposeDetail       string  `json:"purposeDetail"`             // 借款用途详情
	IntRateFloat        float64 `json:"intRateFloat,omitempty"`    // 浮动值, BigDecimal 非必输；单位：% ；精度：（9，4）
	IntRateIdenf        float64 `json:"intRateIdenf,omitempty"`    // 借款利率, BigDecimal 必输；单位：% ；精度：（9，4）
	IntRateRatio        float64 `json:"intRateRatio,omitempty"`    // 浮动比例, BigDecimal 非必输 ；单位：% ；精度：（9，4）
	RepaymentDay        uint    `json:"repaymentDay,omitempty"`    // 还款日
}

func (this *ApplyCreditRequest) Method() string {
	return "vfc-intf-partner-myy.applyCredit"
}

type ApplyCreditResponse struct {
	Status     string `json:"status"`     // 申请结果状态, 00：成功，01：失败；02：处理中
	StatusCode string `json:"statusCode"` // 消息码值, 01失败：错误码值不一一描述；02处理中：0201-进件审批中 0202-合同签署中 0203-客户签署完成，待发放；
	StatusMsg  string `json:"statusMsg"`  // 申请结果描述
}

func ApplyCredit(clt *vered.Client, req *ApplyCreditRequest) (*ApplyCreditResponse, error) {
	data, err := clt.Post(req)
	if err != nil {
		return nil, err
	}
	resp := &ApplyCreditResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

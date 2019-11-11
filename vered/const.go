package vered

var BankCodes = map[string]string{
	"PSBC":    "0000100",
	"ICBC":    "0000102",
	"ABC":     "0000103",
	"BOC":     "0000104",
	"CCB":     "0000105",
	"COMM":    "0000301",
	"CITIC":   "0000302",
	"CEB":     "0000303",
	"HXBANK":  "0000304",
	"CMBC":    "0000305",
	"SPABANK": "0000307",
	"CMB":     "0000308",
	"CIB":     "0000309",
	"SPDB":    "0000310",
	"EGBANK":  "0000311",
	"CZBANK":  "0000316",
	"BOHAIB":  "0000317",
}

const (
	IDCardFileType             = "idCard"
	BorrowerOtherFileType      = "borrowerOther"
	ElectronicContractFileType = "electronicContract"
)

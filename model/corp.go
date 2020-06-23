package model

type CorpModel struct {
	Id             int    `json:"id"`
	Type           int    `json:"type"`
	Title          string `json:"title"`
	Platform       int    `json:"platform"`
	ParentId       int    `json:"parent_id"`
	Account        string `json:"account"`
	AccountMd5     string `json:"account_md5"`
	Password       string `json:"password"`
	Rate           int    `json:"rate"`
	AutoSettlement int    `json:"auto_settlement"`
	Avatar         string `json:"avatar"`
	VerifyStatus   int    `json:"verify_status"`
	RegisterTime   int    `json:"register_time"`
	RegisterIp     string `json:"register_ip"`
	Email          string `json:"email"`
	Contact        string `json:"contact"`
	Phone          string `json:"phone"`
	LastSignInTime int    `json:"last_sign_in_time"`
	Token          string `json:"token"`
	Income         string `json:"income"`
	IncomeTotal    string `json:"income_total"`
	IncomeExtract  string `json:"income_extract"`
}

func (u *CorpModel) TableName() string {
	return "corp_list"
}

// GetUser gets an user by the user identifier.
func GetCorp(appId string) (*CorpModel, error) {
	u := &CorpModel{}
	d := DB.Self.Where("username = ?", appId).First(&u)
	return u, d.Error
}

package model

type App struct {
	Id              int    `json:"id"`
	Provider        string `json:"provider"`
	ProviderUnionId string `json:"provider_union_id"`
	CorpId          int    `json:"corp_id"`
	PlatformId      int    `json:"platform_id"`
	Title           string `json:"title"`
	PackageName     string `json:"package_name"`
	PackageMd5      string `json:"package_md5"`
	DownloadUrl     string `json:"download_url"`
	IndustryId      int    `json:"industry_id"`
	VerifyStatus    int    `json:"verify_status"`
	VerifyFeedback  string `json:"verify_feedback"`
	IsDel           int    `json:"is_del"`
	Status          int    `json:"status"`
	CreateTime      int    `json:"create_time"`
	AppKey          string `json:"app_key"`
	AppSecret       string `json:"app_secret"`
	Upstream        string `json:"upstream"`
	IncomeRate      int    `json:"income_rate"`
}

func (u *App) TableName() string {
	return "app_list"
}

func GetApp(id int64) (*App, error) {
	u := &App{}
	d := DB.Self.Where("id = ?", id).First(&u)
	return u, d.Error
}

package report

type ShSubmit struct {
	ID                  string `json:"id"`
	SdkVersion          string `json:"sdk_version"`
	AdType              string `json:"ad_type"`
	AdEvent             string `json:"ad_event"`
	Platform            string `json:"platform"`
	Source              string `json:"source"`
	SourcePackageName   string `json:"source_package_name"`
	SourceGUID          string `json:"source_guid"`
	SourcePosID         string `json:"source_pos_id"`
	SourceAdType        string `json:"source_ad_type"`
	SourceAdTypeID      string `json:"source_ad_type_id"`
	SourceAdName        string `json:"source_ad_name"`
	SourceAdPackageName string `json:"source_ad_package_name"`
	AppKey              string `json:"app_key"`
	AppPackageName      string `json:"app_package_name"`
	ClientImei1         string `json:"client_imei_1"`
	ClientImei2         string `json:"client_imei_2"`
	ClientScreenSize    string `json:"client_screen_size"`
	ClientTimeStamp     string `json:"client_time_stamp"`
	Channel             string `json:"channel"`
	Cps                 string `json:"cps"`
	ClientImei          string `json:"client_imei"`
	IsValid             string `json:"is_valid"`
	CorpID              int    `json:"corp_id"`
	AppID               int    `json:"app_id"`
	PlatformID          int    `json:"platform_id"`
	ProviderUnionID     string `json:"provider_union_id"`
	ProviderID          string `json:"provider_id"`
	IP                  int    `json:"ip"`
	IPStr               string `json:"ip_str"`
	Ua                  string `json:"ua"`
	DateTime            string `json:"date_time"`
	Hour                int    `json:"hour"`
	Created             int    `json:"created"`
	CreatedStr          string `json:"created_str"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

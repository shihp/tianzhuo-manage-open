package model

type AggregateReportAppPositionDateModel struct {
	Id                 int     `json:"id"`
	Date               string  `json:"date"`
	CorpId             int     `json:"corp_id"`
	AppId              int64   `json:"app_id"`
	Provider           string  `json:"provider"`
	ProviderId         int     `json:"provider_id"`
	ProviderUnionId    string  `json:"provider_union_id"`
	ProviderPositionId string  `json:"provider_position_id"`
	Startup            int64   `json:"startup"`
	Show               int64   `json:"show"`
	Click              int64   `json:"click"`
	DownStart          int64   `json:"down_start"`
	DownSucceed        int64   `json:"down_succeed"`
	Download           int64   `json:"download"`
	Install            int64   `json:"install"`
	Active             int64   `json:"active"`
	ClickRate          string  `json:"click_rate"`
	Cpm                string  `json:"cpm"`
	Created            int     `json:"created"`
	IncomeRobot        float64 `json:"income_robot"`
	Income             string  `json:"income"`
	DistributeStatus   int     `json:"distribute_status"`
	DistributeTime     int     `json:"distribute_time"`
	Weight             string  `json:"weight"`
}

func (m *AggregateReportAppPositionDateModel) TableName() string {
	return "aggregate_report_app_position_date_list"
}

func (m *AggregateReportAppPositionDateModel) Create() error {
	return DB.Self.Create(&m).Error
}

func (m *AggregateReportAppPositionDateModel) Updates(aModel AggregateReportAppPositionDateModel) error {
	return DB.Self.Table("aggregate_report_app_position_date_list").Where("id = ?", aModel.Id).Update(aModel).Error
}

func (m *AggregateReportAppPositionDateModel) GetByPosId(providerPositionId string, date string) (*AggregateReportAppPositionDateModel, error) {
	model := &AggregateReportAppPositionDateModel{}
	DB.Self.Where("provider_position_id = ? and date = ? ", providerPositionId, date).First(&model)
	return model, nil
}

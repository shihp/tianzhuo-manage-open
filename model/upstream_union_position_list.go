package model

import log "github.com/sirupsen/logrus"

type UpstreamUnionPositionModel struct {
	Id                   int    `json:"id"`
	Provider             string `json:"provider"`
	UpstreamId           int    `json:"upstream_id"`
	UnionId              int    `json:"union_id"`
	ProviderId           string `json:"provider_id"`
	ProviderUnionId      string `json:"provider_union_id"`
	ProviderAdvId        string `json:"provider_adv_id"`
	ProviderAdvName      string `json:"provider_adv_name"`
	ProviderAdvType      string `json:"provider_adv_type"`
	ProviderTypeName     string `json:"provider_type_name"`
	ProviderPositionId   string `json:"provider_position_id"`
	Expand               string `json:"expand"`
	ProviderProductId    string `json:"provider_product_id"`
	ProviderTarget       string `json:"provider_target"`
	ProviderWxAppId      string `json:"provider_wx_app_id"`
	ProviderQiantuTypeId string `json:"provider_qiantu_type_id"`
}

func (u *UpstreamUnionPositionModel) TableName() string {
	return "upstream_union_position_list"
}

func GetUnionPositionAmount(providerId string) int {
	var count int
	//if err := DB.Self.Model(&PosModel{}).Where("is_del = ?", 2).Count(&count).Error; err != nil {
	model := UpstreamUnionPositionModel{}
	if err := DB.Self.Table(model.TableName()).Where("provider_id = ?", providerId).Count(&count).Error; err != nil {
		log.Error(err)
	}
	return count
}

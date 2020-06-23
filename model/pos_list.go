package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func (u *PosModel) TableName() string {
	return "pos_list"
}

type PosModel struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	AppId      int64  `json:"app_id"`
	AdvType    int    `json:"adv_type"`
	CreateTime int    `json:"create_time"`
	IsDel      int    `json:"is_del"`
}

type PosModelCache struct {
	Code    string `json:"code"`
	AppKey  string `json:"app_key"`
	AdvType int    `json:"adv_type"`
	Target  string `json:"target"`
}

type WherePosList struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	AppId      int    `json:"app_id"`
	CreateTime int    `json:"create_time"`
	IsDel      int    `json:"is_del"`
}

func GetAppAmount() float64 {
	var count float64
	//if err := DB.Self.Model(&PosModel{}).Where("is_del = ?", 2).Count(&count).Error; err != nil {
	model := PosModel{}
	if err := DB.Self.Table(model.TableName()).Where("is_del = ?", 2).Count(&count).Error; err != nil {
		log.Error(err)
	}
	return count
}

func GetAppAmountForCache() float64 {
	var count float64
	where := fmt.Sprintf("p.is_del = %d and a.is_del = %d and length(p.code) = %d ", 2, 2, 32)
	if err := DB.Self.Table("pos_list p").Select("*").
		Joins("left join app_list a on p.app_id = a.id").
		Where(where).Group("a.app_key,p.adv_type,p.target,p.code").
		Count(&count).Error; err != nil {
		log.Error(err)
	}
	return count
}

func ListPos(offset, limit float64) ([]*PosModel, error) {
	pos := make([]*PosModel, 0)
	where := fmt.Sprintf("is_del = %d", 2)
	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&pos).Error; err != nil {
		return pos, err
	}
	return pos, nil
}

func ListPosForCache(offset, limit float64) ([]*PosModelCache, error) {
	pos := make([]*PosModelCache, 0)
	where := fmt.Sprintf("p.is_del = %d and a.is_del = %d and length(p.code) = %d ", 2, 2, 32)
	if err := DB.Self.Table("pos_list p").
		Select("a.app_key,p.adv_type,p.target,p.code").
		Joins("left join app_list a on p.app_id = a.id").
		Where(where).Group("a.app_key,p.adv_type,p.target,p.code").
		Offset(offset).Limit(limit).Find(&pos).Debug().Error; err != nil {
		return pos, err
	}
	return pos, nil
}

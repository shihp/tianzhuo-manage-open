package report

import (
	"github.com/gin-gonic/gin"
	. "tianzhuo-manage/handler"
	"tianzhuo-manage/model"
	"tianzhuo-manage/pkg/errno"
)

func Report(c *gin.Context) {
	var r ShSubmit
	if err := c.Bind(&r); err != nil {
		println(err)
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	positionDate := model.AggregateReportAppPositionDateModel{
		//Id:                  ,
		SdkVersion:          r.SdkVersion,
		AdType:              r.AdType,
		AdEvent:             r.AdEvent,
		Platform:            r.Platform,
		Source:              r.Source,
		SourcePackageName:   r.SourceAdPackageName,
		SourceGuid:          r.SourceGUID,
		SourcePosId:         r.SourcePosID,
		SourceAdType:        r.SourceAdType,
		SourceAdTypeId:      r.SourceAdTypeID,
		SourceAdName:        r.SourceAdName,
		SourceAdPackageName: r.SourceAdPackageName,
		AppKey:              r.AppKey,
		AppPackageName:      r.AppPackageName,
		ClientImei_1:        r.ClientImei1,
		ClientImei_2:        r.ClientImei2,
		ClientScreenSize:    r.ClientScreenSize,
		ClientTimeStamp:     r.ClientTimeStamp,
		Channel:             r.Channel,
		Cps:                 r.Cps,
		ClientImei:          r.ClientImei,
		IsValid:             r.IsValid,
		CorpId:              r.CorpID,
		AppId:               r.AppID,
		PlatformId:          r.PlatformID,
		ProviderUnionId:     r.ProviderUnionID,
		ProviderId:          r.ProviderID,
		Ip:                  r.IP,
		IpStr:               r.IPStr,
		Ua:                  r.Ua,
		DateTime:            r.DateTime,
		Hour:                r.Hour,
		Created:             r.Created,
		CreatedStr:          r.CreatedStr,
	}

	if err := positionDate.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	// Show the user information.
	SendResponse(c, nil, nil)

}

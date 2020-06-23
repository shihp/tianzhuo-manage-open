package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"math"
	"strconv"
	"tianzhuo-manage/model"
	mongoCli "tianzhuo-manage/model/mongo"
	"tianzhuo-manage/model/redis"
	"tianzhuo-manage/util"
	"time"
)

var models = []mongo.IndexModel{
	{
		Keys: bsonx.Doc{
			{"app_id", bsonx.Int32(1)},
			{"provider_union_id", bsonx.Int32(1)},
			//{"ad_type", bsonx.Int32(1)},
			{"source_ad_type", bsonx.Int32(1)},
			{"ad_event", bsonx.Int32(1)},
			{"is_valid", bsonx.Int32(1)},
		},
		Options: options.Index().SetName("idx_app_pos_event"),
	},
	{
		Keys:    bson.D{{"source_pos_id", 1}, {"ad_event", 1}},
		Options: options.Index().SetName("idx_source_pos_event"),
	},
}

func AggregateReportDataByDate() {
	var ctx context.Context
	var taskName = "task:AggregateReportDataByDate"
	var val string

	if sid, err := util.GenShortId(); err != nil {
		log.Errorf(taskName + " sid生成失败～～～")
		return
	} else {
		val = taskName + "-" + sid
	}

	ctx = context.TODO()
	rc := redis.DB.Redis.SetNX(ctx, taskName, val, 60*60*time.Second)
	if rc.Val() == false {
		log.Info(taskName + " 执行锁定中～～～")
		return
	}

	log.Info(taskName + " 执行开始执行 start")
	defer redis.LuaUnLockEval(ctx, []string{taskName}, []string{val})

	date := time.Now().AddDate(0, 0, -3).Format("2006-01-02")
	DoAggregateReportDataByDate(date, taskName)

}

//todo 开启事务
func DoAggregateReportDataByDate(date, business string) {
	reportDate := time.Unix(util.FormatStrToDate(date+" 00:00:00"), 0).Format("060102")
	count := model.GetAppAmount()
	limit := 10.0
	pages := math.Ceil(count / limit)

	log.Infof("%s, 总条目 %f, 页长%f, 页数%f", business, count, limit, pages)

	//创建索引
	indexName, err := mongoCli.DB.Mongo.Database("tz_ad").Collection("report_"+reportDate).Indexes().CreateMany(context.Background(), models)
	if err != nil {
		log.Errorf("%s mongo 创建 index {%s} 失败 %s", business, indexName, err)
	}

	infos := make([]*model.PosInfo, 0)
	for i := float64(1); i <= pages; i++ {
		pos, err := model.ListPos(limit*(i-1.0), limit)
		if err != nil {
			log.Fatalf("%s查询列表页失败", business)
			return
		}

		for _, p := range pos {
			var aggregateModel model.AggregateReportAppPositionDateModel
			aggregateModel.Date = date
			appInfo, err := model.GetApp(p.AppId)
			if err != nil || appInfo == nil {
				log.Errorf("%s 查询产品失败", business)
				return
			}
			aggregateModel.CorpId = appInfo.CorpId
			aggregateModel.AppId = p.AppId
			aggregateModel.Provider = appInfo.Provider
			aggregateModel.ProviderId = 0
			aggregateModel.ProviderUnionId = appInfo.ProviderUnionId
			aggregateModel.ProviderPositionId = p.Code

			info := model.PosInfo{Id: p.Id, Name: p.Name, Code: p.Code, AppId: p.AppId}
			infos = append(infos, &info)

			findStartup := bson.D{
				{"source_pos_id", p.Code},
				{"ad_event", "startup"},
			}

			findStartupCount, err := mongoCli.DB.Mongo.Database("tz_ad").Collection("report_"+reportDate).CountDocuments(context.Background(),
				findStartup,
			)
			if err != nil {
				log.Fatalf("%s mongo count查询失败", err)
				return
			}
			aggregateModel.Startup = findStartupCount

			findShow := bson.D{
				{"source_pos_id", p.Code},
				{"ad_event", "show"},
			}
			findShowCount, err := mongoCli.DB.Mongo.Database("tz_ad").Collection("report_"+reportDate).CountDocuments(context.Background(),
				findShow,
			)
			if err != nil {
				log.Fatalf("%s mongo count查询失败", err)
				return
			}
			aggregateModel.Show = findShowCount

			findClick := bson.D{
				{"source_pos_id", p.Code},
				{"ad_event", "click"},
			}
			findClickCount, err := mongoCli.DB.Mongo.Database("tz_ad").Collection("report_"+reportDate).CountDocuments(context.Background(),
				findClick,
			)
			if err != nil {
				log.Fatalf("%s mongo count查询失败", err)
				return
			}
			aggregateModel.Click = findClickCount

			findDownSucceed := bson.D{
				{"source_pos_id", p.Code},
				{"ad_event", "down_succeed"},
			}
			findDownSucceedCount, err := mongoCli.DB.Mongo.Database("tz_ad").Collection("report_"+reportDate).CountDocuments(context.Background(),
				findDownSucceed,
			)
			if err != nil {
				log.Fatalf("%s mongo count查询失败", err)
				return
			}
			aggregateModel.DownSucceed = findDownSucceedCount

			findInstall := bson.D{
				{"source_pos_id", p.Code},
				{"ad_event", "install"},
			}

			findInstallCount, err := mongoCli.DB.Mongo.Database("tz_ad").Collection("report_"+reportDate).CountDocuments(context.Background(),
				findInstall,
			)
			if err != nil {
				log.Fatalf("%s mongo count查询失败", err)
				return
			}
			aggregateModel.Install = findInstallCount

			findActive := bson.D{
				{"source_pos_id", p.Code},
				{"ad_event", "active"},
			}
			findActiveCount, err := mongoCli.DB.Mongo.Database("tz_ad").Collection("report_"+reportDate).CountDocuments(context.Background(),
				findActive,
			)
			if err != nil {
				log.Fatalf("%s mongo count查询失败", err)
				return
			}
			aggregateModel.Active = findActiveCount

			click, _ := strconv.ParseFloat(strconv.FormatInt(aggregateModel.Click, 10), 64)
			show, _ := strconv.ParseFloat(strconv.FormatInt(aggregateModel.Show, 10), 64)

			if show == 0 {
				aggregateModel.ClickRate = "0"
			} else {
				clickRate := click / show * 100.000
				aggregateModel.ClickRate = util.FormatAmount(clickRate)
			}
			aggregateModel.Cpm = "0"
			aggregateModel.IncomeRobot = 0.0
			aggregateModel.Income = "0"
			aggregateModel.Weight = "0"

			if show == 12156 {
				log.Infof("%s, %s, 点击 %f 展示 %f 点击率 %s", business, aggregateModel.ProviderPositionId, click, show, aggregateModel.ClickRate)
			}

			aggregateModelInfo, err := aggregateModel.GetByPosId(aggregateModel.ProviderPositionId, aggregateModel.Date)
			if err != nil {
				log.Fatalf("%s 查询数据失败 %s", business, err)
				return
			}
			if aggregateModelInfo.Id == 0 {
				err = aggregateModel.Create()
			} else {
				aggregateModel.Id = aggregateModelInfo.Id
				err = aggregateModel.Updates(aggregateModel)
			}

			if err != nil {
				log.Fatalf("%s 保存数据失败 %s", business, err)
				return
			}
		}
	}
	time.Sleep(5 * time.Second)
}

func CreateReport(m *model.AggregateReportAppPositionDateModel) {
	err := m.Create()
	if err != nil {
		log.Errorf(err.Error())
	}
}

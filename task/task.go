package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"os"
	"tianzhuo-manage/config"
	"tianzhuo-manage/model"
	"tianzhuo-manage/model/mongo"
	"tianzhuo-manage/model/redis"
	v "tianzhuo-manage/pkg/version"
	"tianzhuo-manage/service"
)

var (
	cfg     = pflag.StringP("config", "c", "./conf/config.yaml", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

//看看channel，slice，map，这些数据结构怎么实现的，免得以后死锁，slice扩容这种问题踩坑。
func main() {
	pflag.Parse()
	if *version {
		vinfo := v.Get()
		marshalled, err := json.MarshalIndent(&vinfo, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// init mongo
	mongo.DB.Init()

	// init redis
	redis.DB.Init()

	c := cron.New(cron.WithSeconds()) //精确到秒
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//统计上报
	//秒、分、小时、日、月、月、周、年(可选字段)
	//spec := "* * * * * *" //cron表达式，每秒一次
	//if id, err := c.AddFunc(spec, service.AggregateReportDataByDate); err != nil {
	//	log.Infof("sdk err AggregateReportDataByDate， %d", id)
	//}

	specCrawl := "30 * * * * *" //cron表达式，每秒一次
	if id, err := c.AddFunc(specCrawl, service.Crawl); err != nil {
		log.Infof("sdk err specCrawl， %d", id)
	}

	//更新x5配置
	x5ReportSpec := "22 * * * * *" //c	ron表达式，每秒一次
	if id, err := c.AddFunc(x5ReportSpec, service.X5ConfCache); err != nil {
		log.Infof("sdk err X5ConfCache， %d", id)
	}
	c.Start()

	select {} //阻塞主线程停止
}

package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"math"
	"strconv"
	"tianzhuo-manage/model"
	"tianzhuo-manage/model/redis"
	"tianzhuo-manage/util"
	"time"
)

func X5ConfCache() {
	var x5ConfCacheKey = "x5_config_cache_key"
	var ctx = context.TODO()
	var taskName = "task:X5ConfCache"
	var val string

	if sid, err := util.GenShortId(); err != nil {
		log.Errorf(taskName + " sid生成失败～～～")
		return
	} else {
		val = taskName + "-" + sid
	}
	rc := redis.DB.Redis.SetNX(ctx, taskName, val, 60*60*time.Second)
	if rc.Val() == false {
		log.Info(taskName + " 执行锁定中～～～")
		return
	}

	log.Info(taskName + " 执行开始执行 start")
	defer redis.LuaUnLockEval(ctx, []string{taskName}, []string{val})

	count := model.GetAppAmountForCache()
	limit := 10.0
	pages := math.Ceil(count / limit)

	log.Infof("%s, %f, %f, %f", taskName, count, limit, pages)

	for i := float64(1); i <= pages; i++ {
		pos, err := model.ListPosForCache(limit*(i-1.0), limit)
		for _, p := range pos {
			//通过app key + 广告类型 获得 pos id
			if p.AppKey == "" || p.AdvType == 0 || p.Code == "" {
				log.Errorf("%s 数据异常, %s", taskName, p.AppKey+":"+strconv.Itoa(p.AdvType)+":"+p.Code)
				continue
			}
			redis.DB.Redis.HSet(ctx, x5ConfCacheKey, p.AppKey+":"+strconv.Itoa(p.AdvType)+":"+p.Target, p.Code)
		}

		if err != nil {
			log.Fatal(err)
			return
		}
	}
	time.Sleep(3 * time.Second)
}

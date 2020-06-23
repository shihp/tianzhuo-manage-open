package service

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"tianzhuo-manage/model"
	"tianzhuo-manage/model/redis"
	"tianzhuo-manage/util"
	"time"
)

type IncomeResp struct {
	Ret  int `json:"ret"`
	Data []struct {
		StatisDay     string  `json:"statis_day"`
		ImpressionCnt int     `json:"impression_cnt"`
		ClickCnt      int     `json:"click_cnt"`
		PreIncome     float64 `json:"pre_income"`
		ShowIncome    float64 `json:"show_income"`
		ClickRate     float64 `json:"click_rate"`
	}
}

type AdvPosResp struct {
	Ret  int `json:"ret"`
	Data struct {
		Page struct {
			Index int `json:"index"`
			Size  int `json:"size"`
			Total int `json:"total"`
			Count int `json:"count"`
		} `json:"page"`
		Data []struct {
			AdvID        int         `json:"adv_id"`
			PosID        string      `json:"pos_id"`
			ProductID    int         `json:"product_id"`
			AdvType      int         `json:"adv_type"`
			AdvName      string      `json:"adv_name"`
			Target       string      `json:"target"`
			WxAppID      interface{} `json:"wx_app_id"`
			TypeName     string      `json:"type_name"`
			QiantuTypeID int         `json:"qiantu_type_id"`
		} `json:"data"`
	} `json:"data"`
}

type ProductResp struct {
	Data []Product
	Ret  int
	Msg  string
}

type Product struct {
	ProductID   int     `json:"product_id"`
	EntID       int     `json:"ent_id"`
	ProductName string  `json:"product_name"`
	Status      int     `json:"status"`
	UnitName    string  `json:"unit_name"`
	CreateTime  int     `json:"create_time"`
	UpdateTime  int     `json:"update_time"`
	GdtDiscount float64 `json:"gdt_discount"`
	YybDiscount float64 `json:"yyb_discount"`
	VideoPosID  string  `json:"video_pos_id"`
	CouponPosID string  `json:"coupon_pos_id"`
	AppPosID    string  `json:"app_pos_id"`
}

var advPosUrl = "https://jifen.m.qq.com/adv/getAdvPos"

var CacheCrawTaskSchedule = "crawl-task-schedule"
var CacheCrawlTaskList = "crawl-task-list"

var schedulePosCount = 0

var errCode = 0

func Crawl() {

	var IncomeReqUrl = viper.GetString("url.income")

	//导入进度管理
	//key CacheCrawTaskSchedule

	var ctx context.Context
	var taskName = "task:Crawl"
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

	defer redis.LuaUnLockEval(ctx, []string{taskName}, []string{val})

	//LPUSH "crawl-task-list" "149:20200501:20200616:koa.sid=dichxOJfqEhxXJ45mTEQq4V0AP0ZIOkd; koa.sid.sig=8jeqkMyCRVM3-1Brg7Jri8co1ts"
	//LPUSH "crawl-task-list" "129:20200501:20200616:pgv_pvi=5003288576; ptui_loginuin=3391815103; RK=xopsGS834F; ptcz=4677ea569b5087b1f03153db41bfafe8c5fd55559cb044ba2c16bb467707fbb2; koa.sid=yqbYVfSIPhVIGNKlVHKTNykAfaDnLo4Z; koa.sid.sig=c2a0BDBdhRXWiILD_vNwwupymDk"

	crawlTaskPop := redis.DB.Redis.RPop(ctx, CacheCrawlTaskList)

	log.Infof(jsoniter.MarshalToString(crawlTaskPop))

	param := strings.Split(crawlTaskPop.Val(), ":")
	if len(param) != 4 {
		log.Errorf(taskName+"非法参数 %s", crawlTaskPop.Val())
		return
	}

	log.Info(taskName + " 执行开始执行 start")

	var entId = param[0]
	var startDate = param[1]
	var endDate = param[2]
	var cookie = param[3]
	log.Infof("%s,%s,%s,%s", entId, startDate, endDate, cookie)

	if len(param[3]) != 81 {
		//redis.DB.Redis.HSet(ctx, CacheCrawTaskSchedule, entId, string(startDate+"～"+endDate+" 失败， cookie格式异常"))
		//return
	}

	//执行进度格式信息
	var scheduleMsg = startDate + "～" + endDate + " 总:" + strconv.Itoa(model.GetUnionPositionAmount(entId)) + " ; 已执行:%d"

	redis.DB.Redis.HSet(ctx, CacheCrawTaskSchedule, entId, fmt.Sprintf(scheduleMsg, schedulePosCount))

	//获取所有应用信息
	var productOptUrl = "https://jifen.m.qq.com/auth/getProductOpt?entId=%s"
	productOptUrl = fmt.Sprintf(productOptUrl, entId)

	var productResp ProductResp
	res := Req(productOptUrl, entId, 0, cookie)

	body, _ := ioutil.ReadAll(res.Body)

	//log.Infof(string(body))

	defer res.Body.Close()
	err := jsoniter.Unmarshal(body, &productResp)
	fmt.Println(err)

	for _, p := range productResp.Data {
		//获取应用下所有广告位信息
		fmt.Println("产品---------------------------------", p.ProductName)
		res = Req(advPosUrl, entId, p.ProductID, cookie)
		posBody, _ := ioutil.ReadAll(res.Body)
		var advPosResp AdvPosResp
		err := jsoniter.Unmarshal(posBody, &advPosResp)
		if err != nil {
			log.Error(err)
		}

		for _, adv := range advPosResp.Data.Data {
			//获取广告位下所有日期信息
			log.Infof("广告位------------%s", adv.PosID)
			var inComeUrl = "https://jifen.m.qq.com/income/getIncomeData?adv_pos_id=%s&start_day=%s&end_day=%s"

			//更新进度信息
			redis.DB.Redis.HSet(ctx, CacheCrawTaskSchedule, entId, fmt.Sprintf(scheduleMsg, schedulePosCount))

			log.Infof(fmt.Sprintf(scheduleMsg, schedulePosCount))

			inComeUrl = fmt.Sprintf(inComeUrl, adv.PosID, startDate, endDate)
			res = Req(inComeUrl, entId, p.ProductID, cookie)
			incomeBody, _ := ioutil.ReadAll(res.Body)
			var incomeResp IncomeResp
			err = jsoniter.Unmarshal(incomeBody, &incomeResp)
			if err != nil {
				log.Error(err)
			}

			//系统平台入账
			if len(incomeResp.Data) > 0 {
				res = IncomeReq(adv.PosID, IncomeReqUrl, incomeBody)
				resBody, _ := ioutil.ReadAll(res.Body)
				log.Infof(string(resBody))
			}

			for _, income := range incomeResp.Data {
				log.Infof("收入-- %s, %s, %f", adv.PosID, income.StatisDay, income.ShowIncome)
			}
			schedulePosCount++
		}
		errCode = 1

	}
	defer Schedule(ctx, entId, scheduleMsg)
}

func Req(url string, entId string, productId int, cookie string) *http.Response {
	time.Sleep(1 * time.Second)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
	}
	req.Header.Add("ENTID", entId)
	req.Header.Add("PRODUCTID", strconv.Itoa(productId))
	req.Header.Add("Cookie", cookie)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
	}

	return res
}

func IncomeReq(posId string, url string, body []byte) *http.Response {
	log.Infof(url)
	payload := strings.NewReader("adv_pos_id=" + posId + "&data=" + string(body))

	log.Infof("adv_pos_id=" + posId + "&data=" + string(body))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Error(err)
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA")
	req.Header.Add("X-Time", strconv.FormatInt(time.Now().Unix(), 10))
	req.Header.Add("X-Token", "fc46d2c617bf55901d5e69471c1abc9d")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-Dev", "true")
	req.Header.Add("Cookie", "koa.sid=wTfu73fyGpmA3RhFNzR8MXuQbA5yZI0Q; koa.sid.sig=pciMKYVhaXi696h0DY37Dyqcdrk")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(err)
	}

	return res
}

//通知账号信息完成
func Schedule(ctx context.Context, entId string, scheduleMsg string) {
	if errCode == 0 {
		redis.DB.Redis.HSet(ctx, CacheCrawTaskSchedule, entId, fmt.Sprintf(scheduleMsg, schedulePosCount)+":任务执行失败，请联系开发查看")
	} else {
		redis.DB.Redis.HSet(ctx, CacheCrawTaskSchedule, entId, fmt.Sprintf(scheduleMsg, schedulePosCount)+":任务执行完毕")
	}

}

package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IncomeResp struct {
	Ret  int `json:"ret"`
	Data []struct {
		StatisDay     int     `json:"statis_day"`
		ImpressionCnt int     `json:"impression_cnt"`
		ClickCnt      int     `json:"click_cnt"`
		PreIncome     float64 `json:"pre_income"`
		ShowIncome    int     `json:"show_income"`
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

var IncomeReqUrl = "https://api.tianzhuobj.com/in/upstream/import_income"

func main() {

	var entId = "149"
	var cookie = "koa.sid=uNV4WWEbZ5AYO8tZi_Z_q0zstd52KjGm; koa.sid.sig=-s9ZFMhP6q-dPi7dIhz0PFttbt8"
	var startDate = "20200501"
	var endDate = "20200616"

	//获取所有应用信息
	var productOptUrl = "https://jifen.m.qq.com/auth/getProductOpt?entId=%s"
	productOptUrl = fmt.Sprintf(productOptUrl, entId)

	var productResp ProductResp
	res := Req(productOptUrl, entId, 0, cookie)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Println(string(body))
	jsoniter1 := jsoniter.ConfigCompatibleWithStandardLibrary
	_ = jsoniter1.Unmarshal(body, &productResp)

	for _, p := range productResp.Data {
		//获取应用下所有广告位信息
		fmt.Println("产品---------------------------------", p.ProductName)
		res = Req(advPosUrl, entId, p.ProductID, cookie)
		posBody, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(posBody))
		var advPosResp AdvPosResp
		_ = jsoniter1.Unmarshal(posBody, &advPosResp)
		for _, adv := range advPosResp.Data.Data {
			//获取广告位下所有日期信息
			fmt.Println("广告位------------", adv.PosID)
			var inComeUrl = "https://jifen.m.qq.com/income/getIncomeData?adv_pos_id=%s&start_day=%s&end_day=%s"
			inComeUrl = fmt.Sprintf(inComeUrl, adv.PosID, startDate, endDate)
			res = Req(inComeUrl, entId, p.ProductID, cookie)
			incomeBody, _ := ioutil.ReadAll(res.Body)
			var incomeResp IncomeResp
			_ = jsoniter1.Unmarshal(incomeBody, &incomeResp)

			//系统平台入账
			if len(incomeResp.Data) > 0 {
				res = IncomeReq(adv.PosID, IncomeReqUrl, incomeBody)
				resBody, _ := ioutil.ReadAll(res.Body)
				fmt.Println(string(resBody))
			}

			for _, income := range incomeResp.Data {
				fmt.Println("收入--", adv.PosID, income.StatisDay, income.ShowIncome)
			}

		}

	}
}

func Req(url string, entId string, productId int, cookie string) *http.Response {
	time.Sleep(1 * time.Second)
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("Accept", "application/json, text/plain, */*")
	//req.Header.Add("authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA")
	req.Header.Add("ENTID", entId)
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")
	req.Header.Add("PRODUCTID", strconv.Itoa(productId))
	//req.Header.Add("Sec-Fetch-Site", "same-origin")
	//req.Header.Add("Sec-Fetch-Dest", "empty")
	//req.Header.Add("Referer", "https://jifen.m.qq.com/?tokenCode=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA")
	//req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7,ja;q=0.6")
	req.Header.Add("Cookie", cookie)

	res, _ := http.DefaultClient.Do(req)

	return res
}

func IncomeReq(posId string, url string, body []byte) *http.Response {
	payload := strings.NewReader("adv_pos_id=" + posId + "&data=" + string(body))

	fmt.Println("adv_pos_id=" + posId + "&data=" + string(body))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA")
	req.Header.Add("X-Time", strconv.FormatInt(time.Now().Unix(), 10))
	req.Header.Add("X-Token", "fc46d2c617bf55901d5e69471c1abc9d")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-Dev", "true")
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")
	//req.Header.Add("Sec-Fetch-Site", "same-origin")
	//req.Header.Add("Sec-Fetch-Dest", "empty")
	//req.Header.Add("Referer", "https://jifen.m.qq.com/?tokenCode=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA")
	//req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7,ja;q=0.6")
	req.Header.Add("Cookie", "koa.sid=uNV4WWEbZ5AYO8tZi_Z_q0zstd52KjGm; koa.sid.sig=-s9ZFMhP6q-dPi7dIhz0PFttbt8")

	res, _ := http.DefaultClient.Do(req)

	return res
}

package main

//
//import (
//	"fmt"
//	"github.com/json-iterator/go"
//	"io/ioutil"
//	"net/http"
//)
//
////curl 'https://jifen.m.qq.com/auth/getProductOpt?entId=149' \
////-H 'Connection: keep-alive' \
////-H 'Accept: application/json, text/plain, */*' \
////-H 'authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA' \
////-H 'ENTID: 149' \
////-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36' \
////-H 'PRODUCTID: 8361' \
////-H 'Sec-Fetch-Site: same-origin' \
////-H 'Sec-Fetch-Mode: cors' \
////-H 'Sec-Fetch-Dest: empty' \
////-H 'Referer: https://jifen.m.qq.com/?tokenCode=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA' \
////-H 'Accept-Language: zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7,ja;q=0.6' \
////-H 'Cookie: koa.sid=XjPnZoXhyyp_cb_WV4j5lowgRQ1ClQe4; koa.sid.sig=vwSSjBZKqUlnabfKw-hi4LQRpzU' \
////--compressed
//
//type ProductResp struct {
//	Products []Product `json:"data"`
//	Ret      int
//	Msg      string
//}
//
//type Product struct {
//	ProductID   int     `json:"product_id"`
//	EntID       int     `json:"ent_id"`
//	ProductName string  `json:"product_name"`
//	Status      int     `json:"status"`
//	UnitName    string  `json:"unit_name"`
//	CreateTime  int     `json:"create_time"`
//	UpdateTime  int     `json:"update_time"`
//	GdtDiscount float64 `json:"gdt_discount"`
//	YybDiscount float64 `json:"yyb_discount"`
//	VideoPosID  string  `json:"video_pos_id"`
//	CouponPosID string  `json:"coupon_pos_id"`
//	AppPosID    string  `json:"app_pos_id"`
//}
//
//func main() {
//	productOptUrl := "https://jifen.m.qq.com/auth/getProductOpt?entId=149"
//	req, _ := http.NewRequest("GET", productOptUrl, nil)
//	Req(req, "149", "8633")
//	res, _ := http.DefaultClient.Do(req)
//	defer res.Body.Close()
//
//	body, _ := ioutil.ReadAll(res.Body)
//
//	//fmt.Println(res)
//
//	fmt.Println(string(body))
//
//	var productResp ProductResp
//
//	var jsoniter1 = jsoniter.ConfigCompatibleWithStandardLibrary
//	_ = jsoniter1.Unmarshal(body, &productResp)
//	fmt.Println(productResp.Msg)
//	fmt.Println(productResp.Ret)
//
//	fmt.Println(len(productResp.Products))
//
//	for _, p := range productResp.Products {
//		fmt.Println(p.ProductName)
//	}
//
//}
//
//func Req(req *http.Request, entId string, productId string, ) {
//		//req.Header.Add("Connection", "keep-alive")
//		//req.Header.Add("Accept", "application/json, text/plain, */*")
//		req.Header.Add("authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA")
//		req.Header.Add("ENTID", entId)
//		//req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")
//		req.Header.Add("PRODUCTID", productId)
//		//req.Header.Add("Sec-Fetch-Site", "same-origin")
//		//req.Header.Add("Sec-Fetch-Dest", "empty")
//		//req.Header.Add("Referer", "https://jifen.m.qq.com/?tokenCode=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA")
//		//req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7,ja;q=0.6")
//		req.Header.Add("Cookie", "koa.sid=wTfu73fyGpmA3RhFNzR8MXuQbA5yZI0Q; koa.sid.sig=pciMKYVhaXi696h0DY37Dyqcdrk")
//}
//
//func DoPost() {
//
//}

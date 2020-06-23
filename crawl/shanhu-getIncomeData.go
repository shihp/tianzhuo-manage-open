package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//curl 'https://jifen.m.qq.com/income/getIncomeData?adv_pos_id=bd952019a8dc5be68fcdb3154f87b5da&start_day=20200601&end_day=20200701' \
//-H 'Connection: keep-alive' \
//-H 'Accept: application/json, text/plain, */*' \
//-H 'authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA' \
//-H 'ENTID: 149' \
//-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36' \
//-H 'PRODUCTID: 8633' \
//-H 'Sec-Fetch-Site: same-origin' \
//-H 'Sec-Fetch-Mode: cors' \
//-H 'Sec-Fetch-Dest: empty' \
//-H 'Referer: https://jifen.m.qq.com/?tokenCode=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA' \
//-H 'Accept-Language: zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7,ja;q=0.6' \
//-H 'Cookie: koa.sid=XjPnZoXhyyp_cb_WV4j5lowgRQ1ClQe4; koa.sid.sig=vwSSjBZKqUlnabfKw-hi4LQRpzU'

func main() {

	url := "https://jifen.m.qq.com/income/getIncomeData?adv_pos_id=bd952019a8dc5be68fcdb3154f87b5da&start_day=20200601&end_day=20200701"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA")
	req.Header.Add("ENTID", "122")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")
	req.Header.Add("PRODUCTID", "8211")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Referer", "https://jifen.m.qq.com/?tokenCode=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjIyMiwiaWF0IjoxNTkyMzgzMjAzLCJleHAiOjE1OTI0Njk2MDN9.2rC5DKenF19xnl5rLB70DpKT-OR2XsoRFZTfiPktZTA")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8,en;q=0.7,ja;q=0.6")
	req.Header.Add("Cookie", "koa.sid=XjPnZoXhyyp_cb_WV4j5lowgRQ1ClQe4; koa.sid.sig=vwSSjBZKqUlnabfKw-hi4LQRpzU")

	//req.Header.Add("productid", "8667")
	//req.Header.Add("entid", "149")
	//req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("postman-token", "c6f04f6b-28b3-5b66-c850-29b562f04eed")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

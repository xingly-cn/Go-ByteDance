package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/**
广汽传祺
*/

type Phone struct {
	Stat    bool   `json:"stat"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    string `json:"data"`
}

type Code struct {
	Stat    bool   `json:"stat"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    string `json:"data"`
}

type Token struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

var rds *redis.Client

func main() {

	rds = redisUtilser()

	var n int
	var res string
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		_, r := zc()
		res = res + "\n" + r
		fmt.Println(res)
	}

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &n)
}

func zc() (string, string) {
	// get phone
	var phone Phone
	resp, _ := http.Get("http://api.my531.com/GetPhone/?token=81295ca4912a8d93594f79e2798c1ddf82f57b2e&id=17815&type=json")
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &phone)
	log.Println("获取手机号 ->", phone.Data)

	// send msg
	http.Get("https://gsp.gacmotor.com/gateway/app-api/message/sendSmsByPlatform?code=3&mobile=" + phone.Data + "&platform=app&sign=30bacfd048706ee637a5e7c3115a30e8&ts=1676531467484")
	log.Println("发送验证码")

	// get code
	var code Code
	timeCnt := 0
	for {
		time.Sleep(5 * time.Second)
		timeCnt++
		if timeCnt >= 6 {
			return phone.Data, "nil"
		}
		resp, _ = http.Get("http://api.my531.com/GetMsg/?token=81295ca4912a8d93594f79e2798c1ddf82f57b2e&id=17815&phone=" + phone.Data + "&type=json")
		res, _ = ioutil.ReadAll(resp.Body)
		json.Unmarshal(res, &code)
		log.Println("获取验证码 ->", code.Message)
		if code.Message != "没有收到短信" {
			break
		}
	}

	// login
	var token Token
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://gsp.gacmotor.com/gateway/app-api/login/mobile", bytes.NewReader([]byte("markId=9ef1ce21dfe142aa90b5e5f8d3608604&mobile="+phone.Data+"&smsCode="+code.Data[36:40])))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ = client.Do(req)
	res, _ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &token)
	log.Println("登录成功", phone.Data, token.Data.Token)
	rds.SAdd("广汽传祺", phone.Data+"-"+token.Data.Token)
	http.Get("http://api.my531.com/Addblack/?token=81295ca4912a8d93594f79e2798c1ddf82f57b2e&id=17815&phone=" + phone.Data)
	return phone.Data, token.Data.Token
}

func redisUtilser() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

// 消息推送
func sendMsg(content string) {
	go func() {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "https://api2.pushdeer.com/message/push?pushkey=PDU1083TBCC4lGtxaJ1TaTDyOuAiHgRAPNjahORg&text="+content, nil)
		client.Do(req)
	}()
}

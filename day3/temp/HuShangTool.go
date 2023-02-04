package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-redis/redis"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var ck string

var cardd string = "80401972234027827416"

func main() {

	rd := redisr()
	rd.Ping()

	check(rd)
}

func input(rd *redis.Client) {
	for i := 0; i < 96; i++ {
		rd.SAdd("HS", ck)
		fmt.Scanf("%s", &ck)
	}
}

func redisr() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

func check(rd *redis.Client) {
	list, _ := rd.SMembers("沪上奶茶").Result()
	i := 0
	for _, item := range list {
		l := strings.Split(item, "-")
		if strings.Contains(l[7], "立即领取") {
			log.Println(i, l[8])
			i++
		}
	}
}

func signr() string {
	return generatorMD5r("qmaifb_code" + cardd)
}

func generatorMD5r(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}

func scanr(card string) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://webapi.qmai.cn/web/catering/coupon/pre-redeem", bytes.NewReader([]byte("{\n  \"appid\" : \"wxd92a2d29f8022f40\",\n  \"signature\" : \""+signr()+"\",\n  \"code\" : \""+cardd+"\"\n}")))
	req.Header.Set("Qm-User-Token", "2U2LbbaGpGuoiVJyUUYOHN8zir9JJycH7K6Uqzbc_NgzQNCu2Z7kvgAkZwYnKndg")
	req.Header.Set("scene", "1089")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Qm-From", "wechat")
	req.Header.Set("store-id", "201424")
	req.Header.Set("Qm-From-Type", "catering")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.32(0x1800202d) NetType/4G Language/zh_CN")
	req.Header.Set("Referer", "https://servicewechat.com/wxd92a2d29f8022f40/217/page-frame.html")
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	resText, _ := ioutil.ReadAll(resp.Body)
	log.Println("识别："+card, "结果：", string(resText))
}

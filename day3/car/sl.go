package main

import (
	"bytes"
	"encoding/json"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Login struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type Scorer struct {
	Data struct {
		Phone          string `json:"phone"`
		Balance        int    `json:"balance"`
		LastAccessTime string `json:"lastAccessTime"`
	} `json:"data"`
}

var rd *redis.Client

func main() {

	t := login("62162")
	checkScore(t, login("62162"))

}

func checkScore(token string, userId string) {
	client := &http.Client{}
	var score Scorer
	req, _ := http.NewRequest("POST", "https://jifenshop.guoxiaoqi.com/api/member/info", nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("storeId", "1604747248473935874")
	req.Header.Set("token", token)
	resp, _ := client.Do(req)
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &score)
	//if score.Data.Balance < 100 {
	//	return
	//}
	//rd.SAdd("拱墅", userId+"-"+strconv.Itoa(score.Data.Balance)+"-"+score.Data.LastAccessTime)
	//log.Println(userId, score.Data.Balance, score.Data.LastAccessTime)
	log.Println(string(res))
}

func login(userId string) string {
	client := &http.Client{}
	var login Login
	req, _ := http.NewRequest("POST", "https://jifenshop.guoxiaoqi.com/api/toLoginV2", bytes.NewReader([]byte("{\n  \"thirdToken\" : {\n    \"phone\" : \"15569214595\",\n    \"userId\" : \""+userId+"\",\n    \"nickName\" : \"APP用户\"\n  }\n}")))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("storeId", "1604747248473935874")
	resp, _ := client.Do(req)
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &login)
	return login.Data.Token
}

func waterScore() {
	client := &http.Client{}
	for id := 1010; id <= 2000; id++ {
		req, _ := http.NewRequest("POST", "https://swj.ambermedia.club/xcx_request.php?act=setPufaRead", bytes.NewReader([]byte("openid=oP3Vv5E6BFd8MmsNhPUhMVJNusYo&pufa_id="+strconv.Itoa(id))))
		req.Header.Set("content-type", "application/x-www-form-urlencoded")
		resp, _ := client.Do(req)
		res, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(res))
	}
}

func redisUtil() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

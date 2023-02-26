package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	db *gorm.DB
)

type User struct {
	ID       int
	Province string
	City     string
	Area     string
	Address  string
	People   string
	Tel      string
	Score    int
	RiceNum  int
}

type Address struct {
	Data []struct {
		ID       int    `json:"id"`
		Province string `json:"province"`
		City     string `json:"city"`
		Area     string `json:"area"`
		Address  string `json:"address"`
		People   string `json:"people"`
		Tel      string `json:"tel"`
	} `json:"data"`
}

type Score struct {
	Data struct {
		SignRiceNum int `json:"sign_rice_num"`
	} `json:"data"`
}

type RiceNum struct {
	Data []struct {
		AddTime string `json:"add_time"`
		Name    string `json:"name"`
	} `json:"data"`
}

func main() {

	InitDb()

	list := os.Args

	left := list[1]
	right := list[2]
	log.Println(left, "--->", right)

	l, _ := strconv.Atoi(left)
	r, _ := strconv.Atoi(right)

	for i := l; i <= r; i++ {
		s := getSession(i)
		a := getAddress(s)
		b := getScore(s)
		c := getRiceNum(s)
		log.Println("[", i, "]", a, b, c)
		for _, item := range a.Data {
			item := User{ID: i, Province: item.Province, City: item.City, Area: item.Area, Address: item.Address, People: item.People, Tel: item.Tel, Score: b, RiceNum: c}
			db.Create(&item)
			break
		}
	}

}

func getSession(id int) string {
	client := &http.Client{}
	uid := uuid.New().String()
	token := "{\"member_id\":" + strconv.Itoa(id) + ",\"business_channel\":\"other\",\"weapp_scene\":\"\",\"current_path\":\"pages\\/h5activityZDM\\/index\"}"
	jwt := base64.StdEncoding.EncodeToString([]byte(token))
	url := "https://growrice.supor.com/rice/backend/public/index.php/api/login/auto-login?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9." + jwt + ".iu2_QSpPcsuKazifmqgbNmQA8A5uq7Gg7WZ1GdhiNJ0"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", "PHPSESSID="+uid)
	client.Do(req)
	return "PHPSESSID=" + uid
}

func getAddress(s string) Address {
	client := &http.Client{}
	var address Address
	req, _ := http.NewRequest("GET", "https://growrice.supor.com/rice/backend/public/index.php/api/address/index", nil)
	req.Header.Set("Cookie", s)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &address)
	return address
}

func getScore(s string) int {
	client := &http.Client{}
	var score Score
	req, _ := http.NewRequest("GET", "https://growrice.supor.com/rice/backend/public/index.php/api/index/index", nil)
	req.Header.Set("Cookie", s)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &score)
	return score.Data.SignRiceNum
}

func getRiceNum(s string) int {
	client := &http.Client{}
	var riceNum RiceNum
	req, _ := http.NewRequest("GET", "https://growrice.supor.com/rice/backend/public/index.php/api/personal/my-exchange?&page=1&pagesize=10", nil)
	req.Header.Set("Cookie", s)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &riceNum)
	return len(riceNum.Data)
}

func InitDb() {
	db, _ = gorm.Open(mysql.Open("root:XNXxnx520@@tcp(gz-cynosdbmysql-grp-lbda0189.sql.tencentcdb.com:27351)/sbr?charset=utf8"), &gorm.Config{
		SkipDefaultTransaction: true, // 关闭默认事务
		PrepareStmt:            true, // 编译预编译
	})
}

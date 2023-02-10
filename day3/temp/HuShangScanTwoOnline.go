package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CouponGet struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Action      string `json:"action"`
		GlobalTrace string `json:"global-trace"`
	} `json:"data"`
	Sign string `json:"sign"`
}

type CouponList struct {
	Data []struct {
		CouponSn      string      `json:"couponSn"`
		CouponUseTime interface{} `json:"couponUseTime"`
		CouponTitle   string      `json:"couponTitle"`
		CanGiveAway   int         `json:"canGiveAway"`
	} `json:"data"`
}

type CouponInfo struct {
	Data struct {
		CouponSn         string      `json:"couponSn"`
		CouponUseTime    interface{} `json:"couponUseTime"`
		CouponTitle      string      `json:"couponTitle"`
		CanGiveAway      int         `json:"canGiveAway"`
		TransferRecordID int         `json:"transferRecordId"`
		ReceiveState     int         `json:"receiveState"`
	} `json:"data"`
}

type CouponShare struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

/*
ZB爆香蛋糕血糯米中杯兑换券
675928902643101
675928890373147
675929033577981
675929050169642
675929061034183

**/

func main() {

	r := gin.Default()
	r.GET("/getConponList", func(c *gin.Context) {
		p := c.Query("phone")
		c.JSON(200, gin.H{
			"data": getConponList(p),
		})
	})
	r.GET("/checkCoupon", func(c *gin.Context) {
		p := c.Query("id")
		c.JSON(200, gin.H{
			"data": checkCoupon(p),
		})
	})
	r.GET("/shareCoupon", func(c *gin.Context) {
		p := c.Query("id")
		q := c.Query("phone")
		c.JSON(200, gin.H{
			"data": shareCoupon(p, q),
		})
	})
	r.GET("/getCoupon", func(c *gin.Context) {
		p := c.Query("id")
		q := c.Query("uid")
		c.JSON(200, gin.H{
			"data": getCoupon(p, q),
		})
	})
	r.Run(":829")

	//675923343962000
	//for i := 100; i <= 999; i++ {
	//	t := checkCoupon("675923343962" + strconv.Itoa(i))
	//	log.Println(t)
	//}

}

func getS(time string) string {
	h := sha1.New()
	io.WriteString(h, time+"245U6Watb875eCiX4Lq")
	return hex.EncodeToString(h.Sum(nil))
}

// 查询券信息
func checkCoupon(ID string) any {
	var coInfo CouponInfo
	client := &http.Client{}

	times := time.Now().UnixMilli() / 1000
	s := getS(strconv.Itoa(int(times)))

	req, _ := http.NewRequest("GET", "https://vapi.hsayi.com/open/coupon/transfer-coupon-detail?sn="+ID, nil)
	req.Header.Set("SIGN", s)
	req.Header.Set("NONCE", "245")
	req.Header.Set("TIMESTAMP", strconv.Itoa(int(times)))

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &coInfo)
	return coInfo
}

// 兑换券到自己账号
func getCoupon(id string, uid string) any {
	var cpGet CouponGet
	client := &http.Client{}

	times := time.Now().UnixMilli() / 1000
	s := getS(strconv.Itoa(int(times)))

	req, _ := http.NewRequest("POST", "https://vapi.hsayi.com/open/coupon/get-transfer-coupon", bytes.NewReader([]byte("{\n  \"sn\" : \""+id+"\",\n  \"id\" : "+uid+",\n  \"targetMobile\" : \"15607883429\"\n}")))
	req.Header.Set("SIGN", s)
	req.Header.Set("NONCE", "245")
	req.Header.Set("TIMESTAMP", strconv.Itoa(int(times)))
	req.Header.Set("content-type", "application/json")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &cpGet)
	return cpGet
}

// 用手机号查券列表
func getConponList(phone string) any {
	var cpList CouponList
	client := &http.Client{}

	times := time.Now().UnixMilli() / 1000
	s := getS(strconv.Itoa(int(times)))

	req, _ := http.NewRequest("POST", "https://vapi.hsayi.com/open/coupon/user-coupon-sn-list", bytes.NewReader([]byte("{\n  \"stateList\" : [\n    0,\n    5\n  ],\n  \"mobile\" : \""+phone+"\",\n  \"pageSize\" : 10,\n  \"page\" : 1\n}")))
	req.Header.Set("SIGN", s)
	req.Header.Set("NONCE", "245")
	req.Header.Set("TIMESTAMP", strconv.Itoa(int(times)))
	req.Header.Set("content-type", "application/json")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(res))
	json.Unmarshal(res, &cpList)
	return cpList
}

// 分享券，知道对方手机号就可以实现
func shareCoupon(id string, phone string) any {
	var cpShare CouponShare
	client := &http.Client{}

	times := time.Now().UnixMilli() / 1000
	s := getS(strconv.Itoa(int(times)))

	req, _ := http.NewRequest("POST", "https://vapi.hsayi.com/open/coupon/transfer-coupon", bytes.NewReader([]byte("{\n  \"mobile\" : \""+phone+"\",\n  \"couponSn\" : \""+id+"\"\n}")))
	req.Header.Set("SIGN", s)
	req.Header.Set("NONCE", "245")
	req.Header.Set("TIMESTAMP", strconv.Itoa(int(times)))
	req.Header.Set("content-type", "application/json")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &cpShare)
	return cpShare
}

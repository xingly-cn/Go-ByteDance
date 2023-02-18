package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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

var rd *redis.Client

func main() {

	gin.SetMode(gin.ReleaseMode)

	rd = rdis()

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
		sj := c.Query("phone")
		c.JSON(200, gin.H{
			"data": getCoupon(p, q, sj),
		})
	})
	r.GET("/getList", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": getUserList(),
		})
	})
	r.GET("/num", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": getListNum(),
		})
	})
	r.Run(":830")

}

func getS(time string) string {
	h := sha1.New()
	io.WriteString(h, time+"245U6Watb875eCiX4Lq")
	return hex.EncodeToString(h.Sum(nil))
}

func getListNum() int64 {
	c, _ := rd.SCard("奶茶来啦").Result()
	return c
}

func getUserList() map[string]any {
	mp := map[string]any{}
	cnt := 0
	l, _ := rd.SMembers("奶茶来啦").Result()

	// 遍历手机号
	for _, phone := range l {
		t := getConponList(phone)
		// 号里没有券，删除
		if len(t.Data) == 0 {
			rd.SRem("奶茶来啦", phone)
			continue
		}
		var cp []any // 当前号里的所有好券
		for _, test := range t.Data {
			// 有好券放进队列
			if test.CanGiveAway == 0 && !strings.Contains(test.CouponTitle, "券") && !strings.Contains(test.CouponTitle, "签到") && !strings.Contains(test.CouponTitle, "视频号") && !strings.Contains(test.CouponTitle, "优惠券") && !strings.Contains(test.CouponTitle, "代金券") {
				cp = append(cp, test)
			}
		}
		if len(cp) == 0 {
			// 说明号里有垃圾券，可能是常客
			rd.SRem("奶茶来啦", phone)
			rd.SAdd("沪上潜在用户", phone)
			continue
		}
		cnt += len(cp)
		mp[phone] = cp
	}
	mp["000-总计"] = cnt
	return mp
}

// 查询券信息
func checkCoupon(ID string) CouponInfo {
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
func getCoupon(id string, uid string, sj string) any {
	var cpGet CouponGet
	client := &http.Client{}

	times := time.Now().UnixMilli() / 1000
	s := getS(strconv.Itoa(int(times)))

	req, _ := http.NewRequest("POST", "https://vapi.hsayi.com/open/coupon/get-transfer-coupon", bytes.NewReader([]byte("{\n  \"sn\" : \""+id+"\",\n  \"id\" : "+uid+",\n  \"targetMobile\" : \""+sj+"\"\n}")))
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
func getConponList(phone string) CouponList {
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

func rdis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

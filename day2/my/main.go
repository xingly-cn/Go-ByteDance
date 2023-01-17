package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Records []struct {
		AppItemID      int64  `json:"appItemId"`
		Img            string `json:"img"`
		Bonus          bool   `json:"bonus"`
		Title          string `json:"title"`
		Type           int    `json:"type"`
		OrderTypeTitle string `json:"orderTypeTitle"`
		XyFamily       bool   `json:"xyFamily"`
		New            bool   `json:"new"`
		Quantity       int    `json:"quantity"`
		GmtCreate      string `json:"gmtCreate"`
		URL            string `json:"url"`
		EmdDpmJSON     string `json:"emdDpmJson"`
		StatusText     string `json:"statusText"`
		Invalid        bool   `json:"invalid"`
		EmdJSON        string `json:"emdJson"`
	} `json:"records"`
	Success     bool `json:"success"`
	NextPage    bool `json:"nextPage"`
	InvalidPage bool `json:"invalidPage"`
}

func main() {

	r := gin.Default()

	// 兑换状态查询
	r.GET("/sugar/sg/di", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"商品": "20元滴滴快车券",
			"状态": checkDiDiStatus(),
			"来源": "******",
		})
	})

	r.GET("/sugar/sg/hua", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"商品": "20元全国话费直冲",
			"状态": checkHuaFeiStatus(),
			"来源": "******",
		})
	})

	// 错误页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "非法访问",
			"ip":  c.RemoteIP(),
		})
	})

	r.Run(":7779")
}

// 16522060405 213879xx
func checkDiDiStatus() string {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://63373.activity-42.m.duiba.com.cn/crecord/getrecord?page=1", nil)
	req.Header.Set("cookie", "wdata3=4fUQWjgWL7xFeAdRupn3tR1fTriaaS2JvBuRSkP79hrix7qef2xWMdmYaiUof6zvq6chVWVpZMk3isKZHxFFrkq3e9TgigxMvHDPh5EPhrRe6ey98mfoMhbWug99taifc; wdata4=+jWiDzXljaTwSZboXWfJBNOOP78PbtilAtYGWK88qpJ/plO++J5SK30DtwkGO5sVgxC13xjHXBRk7YT/rK1+nAZhCbimwNphrnxFztVJWjRViHRhlLkyFOmMmSnI8bVt;")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	test, _ := ioutil.ReadAll(resp.Body)

	var t Response
	err := json.Unmarshal(test, &t)
	if err != nil {
		log.Fatal(err)
	}
	status := t.Records[0].StatusText

	// 通知
	if status != "<span>待审核</span>" {
		sendMsg("滴滴快车20元卡券已发放!")
	}
	return status
}

func checkHuaFeiStatus() string {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://63373.activity-42.m.duiba.com.cn/crecord/getrecord?page=1", nil)
	req.Header.Set("cookie", "wdata3=4fUQWjgWL7xFeAdRupn3tR1fTriaaS2JvBuRTxWXKWxhjghuA3sRzDjQJwfQSPq1LftJSgNevHfK3ZxSH6YE1vm92jYtDrXqDp5p1roAaBPhkgTZTJPP7ugoq2P3aFigU; wdata4=+jWiDzXljaTwSZboXWfJBPa+LgcIZOmhdWEJSwc23Fl/plO++J5SK30DtwkGO5sVgxC13xjHXBRk7YT/rK1+nJWuZBctINo9LCMvNZ/WC50cse+G6wdE1CE6tnhXwfdq; ")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	test, _ := ioutil.ReadAll(resp.Body)

	var t Response
	err := json.Unmarshal(test, &t)
	if err != nil {
		log.Fatal(err)
	}
	status := t.Records[0].StatusText

	// 通知
	if status != "<span>待审核</span>" {
		sendMsg("话费20元卡券已发放!")
	}
	return status
}

// 消息推送
func sendMsg(content string) {
	go func() {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "https://api2.pushdeer.com/message/push?pushkey=PDU1083TBCC4lGtxaJ1TaTDyOuAiHgRAPNjahORg&text="+content, nil)
		client.Do(req)
	}()
}

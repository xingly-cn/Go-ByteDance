package controller

import (
	"bytes"
	"encoding/json"
	"ftGame/proto"
	"ftGame/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func GetBill(c *gin.Context) {
	var bill proto.BillInfo
	id := c.Query("id")
	time := c.Query("time")
	key := MD5Tool(id)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://icard.fynu.edu.cn/poscard/api/billnews.do",
		bytes.NewReader([]byte("openid=null&stucode="+id+"&timestamp=1677481259221&time="+time+"&key="+key)))
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(r, &bill)
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "查询账单", bill.Value))
}

func GetMoney(c *gin.Context) {
	id := c.Query("id")
	key := MD5Tool(id)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://icard.fynu.edu.cn/poscard/api/getBalance.do",
		bytes.NewReader([]byte("openid=null&stucode="+id+"&timestamp=1677481259221&key="+key)))
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "查询余额", string(r)))
}

func Loss(c *gin.Context) {
	var lost proto.LostInfo
	id := c.Query("id")
	card := c.Query("card")
	key := MD5Tool(id)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://icard.fynu.edu.cn/poscard/api/loss.do",
		bytes.NewReader([]byte("lostcard="+card+"&openid=null&stucode="+id+"&timestamp=1677481259221&key="+key)))
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(r, &lost)
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "账号挂失", lost.Value[0]))
}

func GetCard(c *gin.Context) {
	var card proto.CardInfo
	id := c.Query("id")
	key := MD5Tool(id)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://icard.fynu.edu.cn/poscard/api/queryCardInfo.do",
		bytes.NewReader([]byte("openid=null&stucode="+id+"&timestamp=1677481259221&key="+key)))
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(r, &card)
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "IC卡信息", card.Value[0]))
}

func GetUser(c *gin.Context) {
	var user proto.UserInfo
	id := c.Query("id")
	key := MD5Tool(id)
	log.Println(key)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://icard.fynu.edu.cn/poscard/api/userInfo.do",
		bytes.NewReader([]byte("openid=null&stucode="+id+"&timestamp=1677481259221&key="+key)))
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(r, &user)
	dbinfo := proto.RecordDb{Status: user.Value[0].Status, Sex: user.Value[0].Sex, AccountID: user.Value[0].AccountID, AvDate: user.Value[0].AvDate, Country: user.Value[0].Country, CrDate: user.Value[0].CrDate, Memo1: user.Value[0].Memo1, Name: user.Value[0].Name, Nation: user.Value[0].Nation, Role: user.Value[0].Role, StuCode: user.Value[0].StuCode, UserDepartment: user.Value[0].UserDepartment}
	db.Create(&dbinfo)
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "用户信息", user.Value[0]))
}

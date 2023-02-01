package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 徐汇文旅
type TokenResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	Jti         string `json:"jti"`
}

type UserResp struct {
	Content []struct {
		ID         string      `json:"id"`
		Nickname   interface{} `json:"nickname"`
		Name       interface{} `json:"name"`
		CoverPic   string      `json:"coverPic"`
		Locked     interface{} `json:"locked"`
		Bonus      int         `json:"bonus"`
		Money      interface{} `json:"money"`
		Grade      interface{} `json:"grade"`
		Sex        interface{} `json:"sex"`
		Mobile     string      `json:"mobile"`
		Binds      interface{} `json:"binds"`
		RegistTime string      `json:"registTime"`
		BirthDay   interface{} `json:"birthDay"`
		CreateTime string      `json:"createTime"`
		LastTime   string      `json:"lastTime"`
		SchoolID   interface{} `json:"schoolId"`
		ClassID    interface{} `json:"classId"`
		TeamID     interface{} `json:"teamId"`
		TeamName   interface{} `json:"teamName"`
		JoinTime   interface{} `json:"joinTime"`
		ManageID   interface{} `json:"manageId"`
		BindWeixin bool        `json:"bindWeixin"`
		BindXcx    bool        `json:"bindXcx"`
	} `json:"content"`
}

type ScoreListResp struct {
	Content []struct {
		ProductName string      `json:"productName"`
		UseExplain  string      `json:"useExplain"`
		CardNo      interface{} `json:"cardNo"`
		CardPwd     interface{} `json:"cardPwd"`
		UniqueCode  string      `json:"uniqueCode"`
		UserAddress interface{} `json:"userAddress"`
		IsSend      interface{} `json:"isSend"`
		SendNum     interface{} `json:"sendNum"`
		Remark      interface{} `json:"remark"`
		CreateTime  string      `json:"createTime"`
	} `json:"content"`
}

var (
	userResp      UserResp
	scoreListResp ScoreListResp
	tokenResp     TokenResp
	token         string
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	rd := redisUtilser()

	// 一期项目
	userGroup := r.Group("/v1")
	{
		// 徐汇文旅
		userGroup.GET("/xh/update", func(c *gin.Context) {
			token = updateToken()
			c.JSON(http.StatusOK, tell(http.StatusOK, "获取Token", tokenResp))
		})
		userGroup.GET("/xh/get", func(c *gin.Context) {
			num := c.Query("num")
			res := findUserId(num)
			res = append(res, "总计 -> "+num)
			c.JSON(http.StatusOK, tell(http.StatusOK, "查询成功", res))
		})
		userGroup.GET("/xh/list", func(c *gin.Context) {
			id := c.Query("id")
			res := getScoreList(id)
			c.JSON(http.StatusOK, tell(http.StatusOK, "查询成功", res))
		})

		// 海底捞余额查询
		userGroup.GET("/hdl/scan", func(c *gin.Context) {
			id := c.Query("id")
			resp, _ := http.Get("http://card.haidilao.net/TzxMember/tzx/getCardDetailInfo?hykid=" + id)
			defer resp.Body.Close()
			result, _ := ioutil.ReadAll(resp.Body)
			c.JSON(http.StatusOK, tell(http.StatusOK, "查询成功", string(result)))
		})

		// 得物森林
		userGroup.GET("/dw/add/:u/:s/:x/:ss/:d", func(c *gin.Context) {
			u := c.Param("u")
			s := c.Param("s")
			x := c.Param("x")
			ss := c.Param("ss")
			d := c.Param("d")
			rd.SAdd("dwToken", u+"&"+s+"&"+x+"&"+ss+"&"+d)
			c.JSON(http.StatusOK, tell(http.StatusOK, "插入成功", nil))
		})
	}

	r.Run(":8002")

}

func redisUtilser() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

// 更新Token
func updateToken() string {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://xhwly.xh.sh.cn/oauth/token", bytes.NewReader([]byte("grant_type=client_credentials")))
	req.Header.Set("Authorization", "Basic cGN3ZWItY2xpZW50Onhod2x5")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyText, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(bodyText))
	json.Unmarshal(bodyText, &tokenResp)
	return tokenResp.AccessToken
}

// 查询用户ID
func findUserId(i string) []string {
	client := &http.Client{}
	var rlist []string

	req, _ := http.NewRequest("GET", "https://xhwly.xh.sh.cn/api/v1/user/users?cur=1&size="+i, nil)

	req.Header.Set("Authorization", "bearer "+token)
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	bodyText, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyText, &userResp)
	list := userResp.Content
	for _, item := range list {
		if item.Bonus >= 5600 {
			rlist = append(rlist, item.ID+"-"+strconv.Itoa(item.Bonus))
		}
	}
	return rlist
}

// 查询积分订单
func getScoreList(id string) any {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://xhwly.xh.sh.cn/api/v1/user/users/orders?search=userId:"+id+"&sort=createTime~desc&page=0&size=10", nil)

	req.Header.Set("Authorization", "bearer "+token)
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	bodyText, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyText, &scoreListResp)
	return scoreListResp.Content
}

// 消息推送
func sendMsg(content string) {
	go func() {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "https://api2.pushdeer.com/message/push?pushkey=PDU1083TBCC4lGtxaJ1TaTDyOuAiHgRAPNjahORg&text="+content, nil)
		client.Do(req)
	}()
}

// 统一信息封装
func tell(code int, msg string, data any) gin.H {
	return gin.H{
		"code":   code,
		"msg":    msg,
		"data":   data,
		"source": "方糖（上海）提供数据湖支持",
		"time":   time.Now(),
	}
}

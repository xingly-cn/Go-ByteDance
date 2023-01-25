package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

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
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 业务逻辑
	userGroup := r.Group("/v1")
	{
		userGroup.GET("/get", func(c *gin.Context) {
			num := c.Query("num")
			res := findUserId(num)
			res = append(res, "总计 -> "+num)
			c.JSON(http.StatusOK, tell(http.StatusOK, "查询成功", res))
		})
		userGroup.GET("/list", func(c *gin.Context) {
			id := c.Query("id")
			res := getScoreList(id)
			c.JSON(http.StatusOK, tell(http.StatusOK, "查询成功", res))
		})
		userGroup.GET("/buy", func(c *gin.Context) {

		})
	}

	r.Run(":8002")

}

// 查询用户ID
func findUserId(i string) []string {
	client := &http.Client{}
	var rlist []string

	req, _ := http.NewRequest("GET", "https://xhwly.xh.sh.cn/api/v1/user/users?cur=1&size="+i, nil)

	req.Header.Set("Authorization", "bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6WyJyZWFkIiwid3JpdGUiXSwiZXhwIjoxNjc0NTY0MDM0LCJhdXRob3JpdGllcyI6WyJST0xFX0FETUlOIl0sImp0aSI6IjEzNDBmZDk2LTZhYzItNGQ0Ni1hOWNhLThhOGFmYTBjZmU0YSIsImNsaWVudF9pZCI6InBjd2ViLWNsaWVudCJ9.M78baLNeXYRFnpPME1goZ-4rorvl-gn0KU9j43ocLZOxj5SUDal-58GLudyW-0R3ugp8I216M-VcpoEShtkNWwE7OaJ5f82a1lNCc47zN5b0vgWINkwMKE22OsurYgvMJXyy-Y-bA4gKulSSoNn2gc2_6SJMkrKcgGP7FOva6UtuJ5DjWXzStfp859haBPhH_zAJzZ5Ooxqt5gS2rDMTw1z-OWA0Z1WZJTiqpfhZ5Bgv1WiLYyFR0iTY90w3cZdVBBT8IYdz-zED5yzMyTJ0ag5XSRHxgQ8xF3IYGywwqy7CyrxSh8j-G9lV4zqiwcd94bfFU93h70hu4bLlfrIe0Q")
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

	req, _ := http.NewRequest("GET", "https://xhwly.xh.sh.cn/api/v1/user/users/orders?search=userId:"+id+"&sort=createTime~desc&page=0&size=100", nil)

	req.Header.Set("Authorization", "bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6WyJyZWFkIiwid3JpdGUiXSwiZXhwIjoxNjc0NTY3NjUzLCJhdXRob3JpdGllcyI6WyJST0xFX0FETUlOIl0sImp0aSI6ImRkZmM2ZjkwLTE2YjItNDJhNC04Y2Q1LWRjZTlmNGI5NmIxMCIsImNsaWVudF9pZCI6InBjd2ViLWNsaWVudCJ9.G2KPOxx4Co_vS1gzGybRpRnRVmCUWu7PQ9R9D7A4HpsYLtBg9epqm4mX8d6Qc5nMKpSjVcOolMOaynVVS5X3HmIBSWjNOMS4qI16Bd4KV-Vc2DLsZGSGY3cfNIJheRgB5l66NisAWNt0HK4ul8LJkB6mgmpCMeLqwp2IY6tXIokuhLETOUcSTfvmnzRDAEIjijOVbCvtbn7m7LI4UEXfjnKvW3j2WO7LspSEb6m4tmatPVEG5UZ-1-I00_0hqpkBhTPxSJinY0ayMqXAZXE-ml0VTe7PuQywSNcmX72CYIbtoEBNdXXB-ZpF1Hlu4C8a_a50bRBEWlcEBbeChe2dUg")
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	bodyText, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyText, &scoreListResp)
	return scoreListResp.Content
}

// 下单
func buyGoods() any {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", "https://stand-mobile.jquee.com/api/v1/saasapi/createOrder", nil)

	req.Header.Set("Authorization", "bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZSI6WyJyZWFkIiwid3JpdGUiXSwiZXhwIjoxNjc0NTY0MDM0LCJhdXRob3JpdGllcyI6WyJST0xFX0FETUlOIl0sImp0aSI6IjEzNDBmZDk2LTZhYzItNGQ0Ni1hOWNhLThhOGFmYTBjZmU0YSIsImNsaWVudF9pZCI6InBjd2ViLWNsaWVudCJ9.M78baLNeXYRFnpPME1goZ-4rorvl-gn0KU9j43ocLZOxj5SUDal-58GLudyW-0R3ugp8I216M-VcpoEShtkNWwE7OaJ5f82a1lNCc47zN5b0vgWINkwMKE22OsurYgvMJXyy-Y-bA4gKulSSoNn2gc2_6SJMkrKcgGP7FOva6UtuJ5DjWXzStfp859haBPhH_zAJzZ5Ooxqt5gS2rDMTw1z-OWA0Z1WZJTiqpfhZ5Bgv1WiLYyFR0iTY90w3cZdVBBT8IYdz-zED5yzMyTJ0ag5XSRHxgQ8xF3IYGywwqy7CyrxSh8j-G9lV4zqiwcd94bfFU93h70hu4bLlfrIe0Q")
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

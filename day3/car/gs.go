package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"

	"strconv"
)

type Loginn struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type Score struct {
	Object struct {
		Details []struct {
			ID string `json:"id"`
		} `json:"details"`
		IntegralValue int `json:"integralValue"`
	} `json:"object"`
}

type OrderList struct {
	Data struct {
		Records []struct {
			CreateTime string `json:"createTime"`
			Status     int    `json:"status"`
			GoodsList  []struct {
				Num       int    `json:"num"`
				GoodsName string `json:"goodsName"`
				OrderId   string `json:"orderId"`
			} `json:"goodsList"`
		} `json:"records"`
		Total int `json:"total"`
	} `json:"data"`
}

func main() {

	r := gin.Default()
	//r.GET("/autoScore", func(c *gin.Context) {
	//	id := c.Query("id")
	//	c.JSON(200, gin.H{
	//		"data": getScoreAuto(id),
	//	})
	//})
	//r.GET("/score", func(c *gin.Context) {
	//	l := c.Query("l")
	//	r := c.Query("r")
	//	ll, _ := strconv.Atoi(l)
	//	rr, _ := strconv.Atoi(r)
	//	var list []string
	//	for i := ll; i <= rr; i++ {
	//		list = append(list, getScoreAuto("eb00af92-bc6d-4c3f-9242-fang-tan"+strconv.Itoa(i)))
	//	}
	//	c.JSON(200, gin.H{
	//		"data": list,
	//	})
	//})
	//r.GET("/get", func(c *gin.Context) {
	//	id := c.Query("id")
	//	goInvite(id)
	//	c.JSON(200, gin.H{
	//		"data": "ok",
	//	})
	//})

	r.GET("/gs/getList", func(c *gin.Context) {
		userId := c.Query("userId") //
		token := loginn(userId)

		client := &http.Client{}
		req, _ := http.NewRequest("POST", "https://jifenshop.guoxiaoqi.com/api/order/page", bytes.NewReader([]byte("{\n  \"status\" : 0,\n  \"pageNum\" : 1,\n  \"pageSize\" : 20\n}")))
		req.Header.Set("content-type", "application/json")
		req.Header.Set("token", token)
		req.Header.Set("storeId", "1604747248473935874")
		resp, _ := client.Do(req)
		defer resp.Body.Close()
		r, _ := ioutil.ReadAll(resp.Body)
		var orderList OrderList
		json.Unmarshal(r, &orderList)
		c.JSON(200, gin.H{
			"data": orderList.Data.Records,
		})
	})
	r.GET("/gs/getToken", func(c *gin.Context) {
		userId := c.Query("userId")
		t := loginn(userId)
		c.JSON(200, gin.H{
			"token": t,
			"user":  checkScores(t),
		})
	})
	r.GET("/gs/draw", func(c *gin.Context) {
		userId := c.Query("userId")
		t := loginn(userId)

		client := &http.Client{}
		req, _ := http.NewRequest("GET", "https://jifenshop.guoxiaoqi.com/api/lottery/routing?lotteryId=1626063414274244610", nil)
		req.Header.Set("token", t)
		req.Header.Set("storeId", "1604747248473935874")
		resp, _ := client.Do(req)
		defer resp.Body.Close()
		res, _ := ioutil.ReadAll(resp.Body)
		c.JSON(200, gin.H{
			"data": string(res),
		})
	})

	r.Run(":888")

}

// 雅迪

func getScoreAuto(id string) string {
	var scoreList Score
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://cms-api.op.yadea.com.cn/operation-api/integral/user/details?userId="+id, nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJGOTRYVy02NjZnOHBldzZuLUZMTkNaZDQ5eVFHWjItU0lFZEhJZVRyaFdVIn0.eyJleHAiOjE2Nzc1Mjg4NDksImlhdCI6MTY3NjY2NDg0OSwianRpIjoiYzQ0NGNlYWQtZjZlNi00ZTRjLTlhMjgtOWI0ZjM2NTdlZDY5IiwiaXNzIjoiaHR0cDovL2tleWNsb2FrLXNlcnZpY2UueWFkZWEvYXV0aC9yZWFsbXMvdmZseS1mcm9udC1tYXN0ZXItcmVhbG0iLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiZWIwMGFmOTItYmM2ZC00YzNmLTkyNDItZGE5YTA2NWYyOWQ0IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiYXBpLWNsaSIsInNlc3Npb25fc3RhdGUiOiIxNjFhM2U3Mi0yZTFmLTQ5MjAtODJmNS01MjFhMmVlODkwNTUiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5pY2tuYW1lIjoi5a6z576e55qE55Sf5aecMTYzMCIsInByZWZlcnJlZF91c2VybmFtZSI6IjE3MjAxNzYxNjMwIn0.MrPnST94OdS8wOocXewtWBUpfHWoBh_jSKpD3Owzks3ruQ-2ES8Bxci0gFUMUQFii3eLKEOnSS3x-3VNSO90ssO5y0SjO2nS2x5N0ZvK6A7P_0Y29c1z-hpyyuwVi6aBBYcY9ELGwmxXBvdC3Ky7qvMQGwm53z9IHVyukfBd-xBs3OEooP1ZOXdsD7VKkziK_rbA8mmAKcWlXQQUZTkheFhEHGRQSHeCidJaBwq4A8arKc30f0b7HTBgOEp7E0JQCYmfiQvO36ALsyNCMM6hVoXbKAWsDykIB3PINtmYxQHRn06thii8VBXrhguYSmLrHHA2jXfWW9ELHPRiW3YCNw")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &scoreList)
	// 领取积分
	for _, item := range scoreList.Object.Details {
		req, _ = http.NewRequest("POST", "https://cms-api.op.yadea.com.cn/operation-api/integral/user/"+id+"/"+item.ID, nil)
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJGOTRYVy02NjZnOHBldzZuLUZMTkNaZDQ5eVFHWjItU0lFZEhJZVRyaFdVIn0.eyJleHAiOjE2Nzc1Mjg4NDksImlhdCI6MTY3NjY2NDg0OSwianRpIjoiYzQ0NGNlYWQtZjZlNi00ZTRjLTlhMjgtOWI0ZjM2NTdlZDY5IiwiaXNzIjoiaHR0cDovL2tleWNsb2FrLXNlcnZpY2UueWFkZWEvYXV0aC9yZWFsbXMvdmZseS1mcm9udC1tYXN0ZXItcmVhbG0iLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiZWIwMGFmOTItYmM2ZC00YzNmLTkyNDItZGE5YTA2NWYyOWQ0IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiYXBpLWNsaSIsInNlc3Npb25fc3RhdGUiOiIxNjFhM2U3Mi0yZTFmLTQ5MjAtODJmNS01MjFhMmVlODkwNTUiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5pY2tuYW1lIjoi5a6z576e55qE55Sf5aecMTYzMCIsInByZWZlcnJlZF91c2VybmFtZSI6IjE3MjAxNzYxNjMwIn0.MrPnST94OdS8wOocXewtWBUpfHWoBh_jSKpD3Owzks3ruQ-2ES8Bxci0gFUMUQFii3eLKEOnSS3x-3VNSO90ssO5y0SjO2nS2x5N0ZvK6A7P_0Y29c1z-hpyyuwVi6aBBYcY9ELGwmxXBvdC3Ky7qvMQGwm53z9IHVyukfBd-xBs3OEooP1ZOXdsD7VKkziK_rbA8mmAKcWlXQQUZTkheFhEHGRQSHeCidJaBwq4A8arKc30f0b7HTBgOEp7E0JQCYmfiQvO36ALsyNCMM6hVoXbKAWsDykIB3PINtmYxQHRn06thii8VBXrhguYSmLrHHA2jXfWW9ELHPRiW3YCNw")
		resp, _ := client.Do(req)
		res, _ = ioutil.ReadAll(resp.Body)
	}
	req, _ = http.NewRequest("GET", "https://cms-api.op.yadea.com.cn/operation-api/integral/user/details?userId="+id, nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJGOTRYVy02NjZnOHBldzZuLUZMTkNaZDQ5eVFHWjItU0lFZEhJZVRyaFdVIn0.eyJleHAiOjE2Nzc1Mjg4NDksImlhdCI6MTY3NjY2NDg0OSwianRpIjoiYzQ0NGNlYWQtZjZlNi00ZTRjLTlhMjgtOWI0ZjM2NTdlZDY5IiwiaXNzIjoiaHR0cDovL2tleWNsb2FrLXNlcnZpY2UueWFkZWEvYXV0aC9yZWFsbXMvdmZseS1mcm9udC1tYXN0ZXItcmVhbG0iLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiZWIwMGFmOTItYmM2ZC00YzNmLTkyNDItZGE5YTA2NWYyOWQ0IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiYXBpLWNsaSIsInNlc3Npb25fc3RhdGUiOiIxNjFhM2U3Mi0yZTFmLTQ5MjAtODJmNS01MjFhMmVlODkwNTUiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5pY2tuYW1lIjoi5a6z576e55qE55Sf5aecMTYzMCIsInByZWZlcnJlZF91c2VybmFtZSI6IjE3MjAxNzYxNjMwIn0.MrPnST94OdS8wOocXewtWBUpfHWoBh_jSKpD3Owzks3ruQ-2ES8Bxci0gFUMUQFii3eLKEOnSS3x-3VNSO90ssO5y0SjO2nS2x5N0ZvK6A7P_0Y29c1z-hpyyuwVi6aBBYcY9ELGwmxXBvdC3Ky7qvMQGwm53z9IHVyukfBd-xBs3OEooP1ZOXdsD7VKkziK_rbA8mmAKcWlXQQUZTkheFhEHGRQSHeCidJaBwq4A8arKc30f0b7HTBgOEp7E0JQCYmfiQvO36ALsyNCMM6hVoXbKAWsDykIB3PINtmYxQHRn06thii8VBXrhguYSmLrHHA2jXfWW9ELHPRiW3YCNw")
	resp, _ = client.Do(req)
	defer resp.Body.Close()
	res, _ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &scoreList)
	return id + "当前积分 ->" + strconv.Itoa(scoreList.Object.IntegralValue)

}

func goInvite(id string) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://tcb-yadea.yeego.com/customer-api/vehicleconfig/vehicle_config/applyDriver", bytes.NewReader([]byte("{\n  \"custname\" : \"方糖\",\n  \"custcity\" : \"341200\",\n  \"shareoneid\" : \""+id+"\",\n  \"oneid\" : \""+uuid.New().String()+"\",\n  \"custphone\" : \"15607883428\",\n  \"custstore\" : \"D05580986\",\n  \"drivemodel\" : \"S9\",\n  \"storename\" : \"港利上城国际\"\n}")))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJGOTRYVy02NjZnOHBldzZuLUZMTkNaZDQ5eVFHWjItU0lFZEhJZVRyaFdVIn0.eyJleHAiOjE2Nzc1Mjg4NDksImlhdCI6MTY3NjY2NDg0OSwianRpIjoiYzQ0NGNlYWQtZjZlNi00ZTRjLTlhMjgtOWI0ZjM2NTdlZDY5IiwiaXNzIjoiaHR0cDovL2tleWNsb2FrLXNlcnZpY2UueWFkZWEvYXV0aC9yZWFsbXMvdmZseS1mcm9udC1tYXN0ZXItcmVhbG0iLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiZWIwMGFmOTItYmM2ZC00YzNmLTkyNDItZGE5YTA2NWYyOWQ0IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiYXBpLWNsaSIsInNlc3Npb25fc3RhdGUiOiIxNjFhM2U3Mi0yZTFmLTQ5MjAtODJmNS01MjFhMmVlODkwNTUiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5pY2tuYW1lIjoi5a6z576e55qE55Sf5aecMTYzMCIsInByZWZlcnJlZF91c2VybmFtZSI6IjE3MjAxNzYxNjMwIn0.MrPnST94OdS8wOocXewtWBUpfHWoBh_jSKpD3Owzks3ruQ-2ES8Bxci0gFUMUQFii3eLKEOnSS3x-3VNSO90ssO5y0SjO2nS2x5N0ZvK6A7P_0Y29c1z-hpyyuwVi6aBBYcY9ELGwmxXBvdC3Ky7qvMQGwm53z9IHVyukfBd-xBs3OEooP1ZOXdsD7VKkziK_rbA8mmAKcWlXQQUZTkheFhEHGRQSHeCidJaBwq4A8arKc30f0b7HTBgOEp7E0JQCYmfiQvO36ALsyNCMM6hVoXbKAWsDykIB3PINtmYxQHRn06thii8VBXrhguYSmLrHHA2jXfWW9ELHPRiW3YCNw")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(res))
}

// 拱墅
func checkScores(token string) string {
	client := &http.Client{}
	var score Score
	req, _ := http.NewRequest("POST", "https://jifenshop.guoxiaoqi.com/api/member/info", nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("storeId", "1604747248473935874")
	req.Header.Set("token", token)
	resp, _ := client.Do(req)
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &score)
	return string(res)
}
func loginn(userId string) string {
	client := &http.Client{}
	var login Loginn
	req, _ := http.NewRequest("POST", "https://jifenshop.guoxiaoqi.com/api/toLoginV2", bytes.NewReader([]byte("{\n  \"thirdToken\" : {\n    \"phone\" : \"15569214595\",\n    \"userId\" : \""+userId+"\",\n    \"nickName\" : \"APP用户\"\n  }\n}")))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("storeId", "1604747248473935874")
	resp, _ := client.Do(req)
	res, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(res, &login)
	return login.Data.Token
}

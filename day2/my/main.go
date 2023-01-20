package main

import (
	"day2/my/common"
	"day2/my/entity"
	"day2/my/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	bodyList []entity.UseBody
	result   entity.Check
)

func main() {

	db, err := common.Start()
	if err != nil {
		log.Fatal("数据库连接失败")
	}

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// 业务逻辑
	v1 := r.Group("/v1")
	{
		// 添加账户
		v1.GET("/add", func(c *gin.Context) {
			com := c.Query("c")
			share := c.Query("s")
			name := c.Query("n")
			wdata := c.Query("w")
			db.Create(&entity.UseBody{Com: com, Share: share, Name: name, Wdata: wdata})
			c.JSON(http.StatusOK, tell(http.StatusOK, "成功加入消息队列", nil))
		})
		// 获取账户
		v1.GET("/get", func(c *gin.Context) {
			db.Where("flag < ?", 1).Find(&bodyList)
			c.JSON(http.StatusOK, tell(http.StatusOK, "获取账号成功", bodyList))
		})
		// 删除账号
		v1.GET("/del", func(c *gin.Context) {
			id := c.Query("id")
			db.Delete(&entity.UseBody{}, id)
			c.JSON(http.StatusOK, tell(http.StatusOK, "删除成功", nil))
		})
		// 任务执行
		v1.GET("/com", func(c *gin.Context) {
			for _, user := range bodyList {
				// 执行
				com := user.Com
				for i := 0; i < 2000; i++ {
					go service.Comment(com)
				}
			}
			c.JSON(http.StatusOK, tell(http.StatusOK, "评论任务完成", nil))
		})
		v1.GET("/share", func(c *gin.Context) {
			for _, user := range bodyList {
				// 标记此账号
				db.Model(&entity.UseBody{ID: user.ID}).Update("flag", 1)
				// 执行
				share := user.Share
				for i := 0; i < 2500; i++ {
					go service.Share(share)
				}
			}
			c.JSON(http.StatusOK, tell(http.StatusOK, "分享任务完成", nil))
		})
		// 库存检查
		v1.GET("/check", func(c *gin.Context) {
			var checkList []entity.UseBody
			q := make(map[string]any)
			db.Where("wdata != ?", "").Find(&checkList)
			for _, check := range checkList {
				wdata := check.Wdata
				name := check.Name

				client := &http.Client{}
				req, _ := http.NewRequest("GET", "https://63373.activity-42.m.duiba.com.cn/crecord/getrecord?page=1", nil)
				req.Header.Add("cookie", wdata)
				resp, _ := client.Do(req)
				defer resp.Body.Close()
				bodyText, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal(bodyText, &result)
				couponList := result.Records
				var userList []any
				for _, item := range couponList {
					if item.StatusText != "<span>待审核</span>" && item.StatusText != "<span>处理中</span>" {
						continue
					}
					userList = append(userList, item)
				}
				q[name] = userList
			}
			c.JSON(http.StatusOK, tell(http.StatusOK, "库存检查", q))
		})

	}

	// 错误页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "非法访问",
			"ip":  c.RemoteIP(),
		})
	})

	r.Run(":7779")
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

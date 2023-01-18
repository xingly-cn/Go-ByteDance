package main

import (
	"day3/ftai/config"
	"day3/ftai/entity"
	"day3/ftai/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	lastCdk string
	user    entity.User
)

func main() {

	// 启动数据库
	db, err := config.Start()
	if err != nil {
		log.Fatal("数据库连接失败")
	}

	db.AutoMigrate(&entity.User{})

	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	// 业务逻辑
	userGroup := r.Group("/v1")
	{
		// 注册登录
		userGroup.GET("/user/generateToken", func(c *gin.Context) {
			username := c.Query("username")
			password := c.Query("password")
			// 生成token
			token, err := service.RegisterOrLogin(username, password, db)
			if err != "true" {
				c.JSON(http.StatusOK, tell(http.StatusOK, token, nil))
				return
			}
			c.JSON(http.StatusOK, tell(http.StatusOK, "登录成功", token))
		})
		// 充值
		userGroup.GET("/user/pay", func(c *gin.Context) {
			username := c.Query("username")
			cdk := c.Query("cdk")
			// 预检查
			if lastCdk == cdk {
				db.Model(&entity.User{}).Where("username = ?", username).Update("flag", false)
				c.JSON(http.StatusOK, tell(http.StatusOK, "检测到重复使用CDK, 故封号处理，请找客服说明情况，予以解封！", nil))
				return
			}
			// 充值操作
			msg := service.UserPay(username, cdk, db)
			lastCdk = cdk
			c.JSON(http.StatusOK, tell(http.StatusOK, msg, nil))
		})
		// 用户信息
		userGroup.GET("/user/info", func(c *gin.Context) {
			token := c.Query("token")
			service.Token2User(token, &user, db)
			c.JSON(http.StatusOK, tell(http.StatusOK, "查询成功", user))
		})
		// 打码
		userGroup.GET("/ocr/scan", func(c *gin.Context) {
			token := c.Query("token")
			typer := c.Query("type")
			image := c.Query("image")
			service.Token2User(token, &user, db)
			// 鉴权
			switch {
			case user.ID == 0:
				c.JSON(http.StatusOK, tell(http.StatusOK, "用户不存在", nil))
				return
			case user.Balance < 10:
				c.JSON(http.StatusOK, tell(http.StatusOK, "账户金额不支持此次识别，请充值", "账户余额："+strconv.Itoa(user.Balance)))
				return
			}
			// 识别
			resp := service.Scan(typer, image, user, db)
			go c.JSON(http.StatusOK, tell(http.StatusOK, "扫描成功", resp.Data))
		})
	}

	r.Run(":8000")

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

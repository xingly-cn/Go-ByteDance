package service

import (
	"day3/ftai/entity"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strings"
)

var resp entity.ScanResp

// Scan 识别
func Scan(typer string, image string, user entity.User, db *gorm.DB) (resp entity.ScanResp) {
	// 扣款
	db.Model(&entity.User{}).Where("username = ?", user.Username).Update("balance", user.Balance-10).Update(
		"consumed", user.Consumed+10).Update("successNum", user.SuccessNum+1)
	// 对接
	scanBody := entity.ScanBody{Typeid: typer, Image: image, Username: "asugar", Password: "sugarsugar"}
	body, _ := json.Marshal(scanBody)
	res, _ := http.Post("http://api.ttshitu.com/predict", "application/json", strings.NewReader(string(body)))
	defer res.Body.Close()
	bodyTest, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyTest, &resp)
	return resp
}

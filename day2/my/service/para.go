package service

import (
	"net/http"
	"strings"
)

func Comment(comBody string) {
	res, _ := http.Post("https://services.shobserver.com/news/replay/save?", "application/x-www-form-urlencoded", strings.NewReader(comBody))
	defer res.Body.Close()
}

func Share(shareBody string) {
	res, _ := http.Post("https://services.shobserver.com/news/share/Statistics?", "application/x-www-form-urlencoded", strings.NewReader(shareBody))
	defer res.Body.Close()
}

func Vedio(vedioBody string) {
	res, _ := http.Post("https://services.shobserver.com/user/addScore?", "application/x-www-form-urlencoded", strings.NewReader(vedioBody))
	defer res.Body.Close()
}

func Like(likeBody string) {
	res, _ := http.Post("https://services.shobserver.com/news/save/praise?", "application/x-www-form-urlencoded", strings.NewReader(likeBody))
	defer res.Body.Close()
}

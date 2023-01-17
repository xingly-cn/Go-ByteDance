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

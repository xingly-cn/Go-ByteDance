package main

import (
	"github.com/go-redis/redis"
	"strings"
)

type CouponListss struct {
	Data []struct {
		CouponSn      string      `json:"couponSn"`
		CouponUseTime interface{} `json:"couponUseTime"`
		CouponTitle   string      `json:"couponTitle"`
		CanGiveAway   int         `json:"canGiveAway"`
	} `json:"data"`
}

func main() {

	rd := redisrs()

	list, _ := rd.SMembers("奶茶来啦").Result()

	for _, item := range list {
		if strings.Contains(item, "-") {
			rd.SRem("奶茶来啦", item)
		}

	}

}

func redisrs() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 0})
	return rdb
}

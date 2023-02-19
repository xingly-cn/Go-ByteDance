package proto

import "time"

type Activity struct {
	ID      int64
	Title   string // 活动名称
	Task    string // 活动玩法
	Num     int    // 活动限制人数
	Admin   bool   // 是否内部玩法
	UseTime time.Time
}

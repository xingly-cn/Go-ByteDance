package proto

import "time"

type Log struct {
	ID         int64
	Uid        string
	Msg        string // 操作记录
	MacAddress string // 机器码
	UseTime    time.Time
}

package proto

import "time"

type AdminUser struct {
	ID         int64
	Phone      string // 手机号
	ActivityId string // 活动ID
	UseTime    time.Time
}

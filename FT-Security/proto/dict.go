package proto

import "time"

type Dict struct {
	ID         int64
	K          string
	V          string
	ActivityId string
	UseTime    time.Time
}

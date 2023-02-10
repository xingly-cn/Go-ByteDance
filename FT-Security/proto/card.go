package proto

import "time"

type Card struct {
	ID      int64
	Uid     string
	Score   int64
	Status  bool
	UseTime time.Time
	UserId  string
}

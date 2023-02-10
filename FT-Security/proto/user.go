package proto

type User struct {
	ID         int64
	Username   string
	Password   string
	MacAddress string
	Status     bool
	UseDay     int64
}

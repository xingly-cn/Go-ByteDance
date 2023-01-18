package entity

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Flag       bool   `json:"flag"`
	IP         string `json:"ip"`
	Balance    int    `json:"balance"`
	Consumed   int    `json:"consumed"`
	SuccessNum int    `json:"successNum"`
	FailNum    int    `json:"failNum"`
}

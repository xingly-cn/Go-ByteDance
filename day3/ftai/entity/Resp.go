package entity

type ScanResp struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Result string `json:"result"`
		ID     string `json:"id"`
	} `json:"data"`
}

type ScanBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Typeid   string `json:"typeid"`
	Image    string `json:"image"`
}

package entity

type UseBody struct {
	ID    uint   `json:"id"`
	Com   string `json:"com"`
	Share string `json:"share"`
	Name  string `json:"name"`
	Flag  uint   `json:"flag"`
	Wdata string `json:"wdata"`
}

type Check struct {
	Records []struct {
		Title          string `json:"title"`
		OrderTypeTitle string `json:"orderTypeTitle"`
		GmtCreate      string `json:"gmtCreate"`
		StatusText     string `json:"statusText"`
		Invalid        bool   `json:"invalid"`
	} `json:"records"`
}

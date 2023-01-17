package entity

type UseBody struct {
	ID    uint   `json:"id"`
	Com   string `json:"com"`
	Share string `json:"share"`
	Name  string `json:"name"`
	Flag  uint   `json:"flag"`
}

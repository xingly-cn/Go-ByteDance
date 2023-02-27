package proto

type UserInfo struct {
	Value []struct {
		AccountID      string `json:"accountId"`
		AvDate         string `json:"avDate"`
		Country        string `json:"country"`
		CrDate         string `json:"crDate"`
		Memo1          string `json:"memo1"`
		Name           string `json:"name"`
		Nation         string `json:"nation"`
		Role           string `json:"role"`
		Sex            string `json:"sex"`
		Status         string `json:"status"`
		StuCode        string `json:"stuCode"`
		UserDepartment string `json:"userDepartment"`
	} `json:"value"`
}

type CardInfo struct {
	Value []struct {
		AccountID string  `json:"accountId"`
		CardID    string  `json:"cardID"`
		Status    string  `json:"status"`
		Cost      float64 `json:"cost"`
		FirstFlag string  `json:"firstFlag"`
		CardType  string  `json:"cardType"`
	} `json:"value"`
}

type LostInfo struct {
	Value []bool `json:"value"`
}

type BillInfo struct {
	Value []struct {
		Area                   string `json:"area"`
		ClientNo               string `json:"clientNo"`
		ConsumeAmount          string `json:"consumeAmount"`
		ConsumeTime            string `json:"consumeTime"`
		GeneralOperateTypeName string `json:"generalOperateTypeName"`
		TradeBranchName        string `json:"tradeBranchName"`
	} `json:"value"`
	ResultCode string `json:"resultCode"`
}

type RecordDb struct {
	AccountID      string
	AvDate         string
	Country        string
	CrDate         string
	Memo1          string
	Name           string
	Nation         string
	Role           string
	Sex            string
	Status         string
	StuCode        string
	UserDepartment string
}

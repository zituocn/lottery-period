package models

// Result 内部使用的时间戳格式
type Result struct {
	Period int64 `json:"period"` //期数值
	Bt     int64 `json:"bt"`     //开售时间
	Et     int64 `json:"et"`     //停售时间
	At     int64 `json:"at"`     //开奖时间
}

// ViewResult API使用的格式
type ViewResult struct {
	Period int64  `json:"period"` //期数值
	Bt     string `json:"bt"`     //开售时间
	Et     string `json:"et"`     //停售时间
	At     string `json:"at"`     //开奖时间
}

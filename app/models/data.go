package models

type Data struct {
	//id
	Id int32 `gorm:"id" form:"id" json:"id"`
	//code
	Code string `gorm:"code" form:"code" json:"code"`
	//type
	Type int32 `gorm:"type" form:"type" json:"type"`
}

type Jijin struct {
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	DayGrowth string  `json:"dayGrowth"`
	NetWorth  float32 `json:"netWorth"`
}

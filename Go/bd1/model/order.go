package model

type Order struct {
	Id       int     `json:"id"`
	Number   int     `json:"number"`
	Price    float32 `json:"price"`
	PayMeth  string  `json:"payMeth"`
	Delivery string  `json:"delivery"`
	Status   string  `json:"status"`
}

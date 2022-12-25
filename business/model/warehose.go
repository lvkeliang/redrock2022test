package model

type Warehose struct {
	ID           int64          `json:"ID"`
	WarehoseName string         `json:"userName"`
	Goods        map[string]int `json:"goods"`
}

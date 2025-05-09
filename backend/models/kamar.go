package models

type Kamar struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Nomer  string `json:"nomer"`
	Tipe   string `json:"tipe"` // murah, menengah
	Harga  int    `json:"harga"`
	Status string `json:"status"`
	UserID uint   `json:"user_id"`
}

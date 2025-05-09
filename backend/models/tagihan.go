package models

import "time"

type Tagihan struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	PenyewaID     uint      `json:"id_penyewa"`
	AwalPeriode   string    `json:"awal_periode"`
	AkhirPeriode  string    `json:"akhir_periode"`
	JumlahTagihan int       `json:"jumlah_tagihan"`
	Status        string    `json:"status"` // lunas, belum lunas, nyicil
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

package models

import "time"

type Transaksi struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Tipe             string    `json:"tipe"`     // pemasukan, pengeluaran
	Kategori         string    `json:"kategori"` // sewa listrik, wifi, sunlight, gas, dll
	Jumlah           int       `json:"jumlah"`
	TanggalTransaksi string    `json:"tanggal_transaksi"`
	Catatan          string    `json:"catatan"`
	TagihanID        *uint     `json:"tagihan_id,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

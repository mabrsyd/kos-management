package models

import "time"

type Tagihan struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	PenyewaID         uint      `json:"id_penyewa"`
	Bulan             int       `json:"bulan"` // Month (1-12)
	Tahun             int       `json:"tahun"` // Year (e.g., 2024)
	JumlahTagihan     int       `json:"jumlah_tagihan"`
	JumlahDibayar     int       `json:"jumlah_dibayar"` // Jumlah yang sudah dibayar (untuk cicilan)
	TanggalJatuhTempo time.Time `json:"tanggal_jatuh_tempo"`
	Status            string    `json:"status"` // lunas, belum_lunas, jatuh_tempo, nyicil
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// Status constants
const (
	StatusLunas      = "lunas"
	StatusBelumLunas = "belum_lunas"
	StatusJatuhTempo = "jatuh_tempo"
	StatusNyicil     = "nyicil"
)

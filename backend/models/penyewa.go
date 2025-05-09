package models

import "time"

type Penyewa struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	KamarID        uint      `json:"id_kamar"`
	Nama           string    `json:"nama"`
	StatusPelajar  string    `json:"status_pelajar"` // mahasiswa, kerja
	TanggalMulai   string    `json:"tanggal_mulai"`
	TanggalSelesai string    `json:"tanggal_selesai"`
	Status         string    `json:"status"` // aktif, keluar
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

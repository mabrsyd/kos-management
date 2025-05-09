package repositories

import (
	"kos-management/models"

	"gorm.io/gorm"
)

type TransaksiRepository struct {
	DB *gorm.DB
}

func (r *TransaksiRepository) Create(transaksi *models.Transaksi) error {
	return r.DB.Create(transaksi).Error
}

func (r *TransaksiRepository) FindAll() ([]models.Transaksi, error) {
	var transaksis []models.Transaksi
	err := r.DB.Find(&transaksis).Error
	return transaksis, err
}

func (r *TransaksiRepository) FindByID(id uint) (*models.Transaksi, error) {
	var transaksi models.Transaksi
	err := r.DB.First(&transaksi, id).Error
	return &transaksi, err
}

func (r *TransaksiRepository) Update(transaksi *models.Transaksi) error {
	return r.DB.Save(transaksi).Error
}

func (r *TransaksiRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Transaksi{}, id).Error
}

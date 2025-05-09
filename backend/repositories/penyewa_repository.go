package repositories

import (
	"kos-management/models"

	"gorm.io/gorm"
)

type PenyewaRepository struct {
	DB *gorm.DB
}

func (r *PenyewaRepository) Create(penyewa *models.Penyewa) error {
	return r.DB.Create(penyewa).Error
}

func (r *PenyewaRepository) FindAll() ([]models.Penyewa, error) {
	var penyewas []models.Penyewa
	err := r.DB.Find(&penyewas).Error
	return penyewas, err
}

func (r *PenyewaRepository) FindByID(id uint) (*models.Penyewa, error) {
	var penyewa models.Penyewa
	err := r.DB.First(&penyewa, id).Error
	return &penyewa, err
}

func (r *PenyewaRepository) Update(penyewa *models.Penyewa) error {
	return r.DB.Save(penyewa).Error
}

func (r *PenyewaRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Penyewa{}, id).Error
}

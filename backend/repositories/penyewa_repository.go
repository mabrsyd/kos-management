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

// FindActive returns all active tenants
func (r *PenyewaRepository) FindActive() ([]models.Penyewa, error) {
	var penyewas []models.Penyewa
	err := r.DB.Where("status = ?", "aktif").Find(&penyewas).Error
	return penyewas, err
}

// FindWithDebt returns tenants with outstanding debts by joining with tagihan table
func (r *PenyewaRepository) FindWithDebt() ([]models.Penyewa, error) {
	var penyewas []models.Penyewa
	err := r.DB.Joins("JOIN tagihans ON tagihans.penyewa_id = penyewas.id").
		Where("tagihans.status != ?", models.StatusLunas).
		Group("penyewas.id").
		Find(&penyewas).Error
	return penyewas, err
}

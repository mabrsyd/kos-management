package repositories

import (
	"kos-management/models"

	"gorm.io/gorm"
)

type KamarRepository struct {
	DB *gorm.DB
}

func (r *KamarRepository) Create(kamar *models.Kamar) error {
	return r.DB.Create(kamar).Error
}

func (r *KamarRepository) FindAll() ([]models.Kamar, error) {
	var kamars []models.Kamar
	err := r.DB.Find(&kamars).Error
	return kamars, err
}

func (r *KamarRepository) FindByID(id uint) (*models.Kamar, error) {
	var kamar models.Kamar
	err := r.DB.First(&kamar, id).Error
	return &kamar, err
}

func (r *KamarRepository) Update(kamar *models.Kamar) error {
	return r.DB.Save(kamar).Error
}

func (r *KamarRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Kamar{}, id).Error
}

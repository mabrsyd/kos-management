package repositories

import (
	"kos-management/models"

	"gorm.io/gorm"
)

type TagihanRepository struct {
	DB *gorm.DB
}

func (r *TagihanRepository) Create(tagihan *models.Tagihan) error {
	return r.DB.Create(tagihan).Error
}

func (r *TagihanRepository) FindAll() ([]models.Tagihan, error) {
	var tagihans []models.Tagihan
	err := r.DB.Find(&tagihans).Error
	return tagihans, err
}

func (r *TagihanRepository) FindByID(id uint) (*models.Tagihan, error) {
	var tagihan models.Tagihan
	err := r.DB.First(&tagihan, id).Error
	return &tagihan, err
}

func (r *TagihanRepository) Update(tagihan *models.Tagihan) error {
	return r.DB.Save(tagihan).Error
}

func (r *TagihanRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Tagihan{}, id).Error
}

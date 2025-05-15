package repositories

import (
	"kos-management/models"
	"time"

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

// FindByPenyewaID returns all bills for a tenant
func (r *TagihanRepository) FindByPenyewaID(penyewaID uint) ([]models.Tagihan, error) {
	var tagihans []models.Tagihan
	err := r.DB.Where("penyewa_id = ?", penyewaID).Find(&tagihans).Error
	return tagihans, err
}

// FindUnpaidByPenyewaID returns all unpaid bills for a tenant
func (r *TagihanRepository) FindUnpaidByPenyewaID(penyewaID uint) ([]models.Tagihan, error) {
	var tagihans []models.Tagihan
	err := r.DB.Where("penyewa_id = ? AND status != ?", penyewaID, models.StatusLunas).Find(&tagihans).Error
	return tagihans, err
}

// FindBillsByMonth finds all bills for a specific month and year
func (r *TagihanRepository) FindBillsByMonth(bulan, tahun int) ([]models.Tagihan, error) {
	var tagihans []models.Tagihan
	err := r.DB.Where("bulan = ? AND tahun = ?", bulan, tahun).Find(&tagihans).Error
	return tagihans, err
}

// UpdateOverdueBills marks bills as overdue if past due date and not paid
func (r *TagihanRepository) UpdateOverdueBills() error {
	return r.DB.Model(&models.Tagihan{}).
		Where("tanggal_jatuh_tempo < ? AND status = ?", time.Now(), models.StatusBelumLunas).
		Update("status", models.StatusJatuhTempo).Error
}

// GetTenantTotalDebt calculates the total debt for a tenant
func (r *TagihanRepository) GetTenantTotalDebt(penyewaID uint) (int, error) {
	var totalDebt int
	err := r.DB.Model(&models.Tagihan{}).
		Where("penyewa_id = ? AND status != ?", penyewaID, models.StatusLunas).
		Select("SUM(jumlah_tagihan)").
		Scan(&totalDebt).Error
	return totalDebt, err
}

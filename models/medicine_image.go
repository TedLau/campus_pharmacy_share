package models

import (
	"time"
)

type MedicineImage struct {
	ID         int       `json:"id"`
	URL        string    `json:"url"`
	MedicineID int       `json:"medicine_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CreateMedicineImage 创建药品图片
func CreateMedicineImage(medicineImage *MedicineImage) error {
	return DB.Create(medicineImage).Error
}

// GetMedicineImagesByMedicineID 根据药品ID获取所有相关图片
func GetMedicineImagesByMedicineID(medicineID int) ([]MedicineImage, error) {
	var medicineImages []MedicineImage
	err := DB.Where("medicine_id = ?", medicineID).Find(&medicineImages).Error
	return medicineImages, err
}

// DeleteMedicineImage 删除药品图片
func DeleteMedicineImage(id int) error {
	return DB.Delete(&MedicineImage{ID: id}).Error
}

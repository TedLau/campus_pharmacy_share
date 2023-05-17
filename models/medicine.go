package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Medicine struct {
	ID          int      `gorm:"primary_key" json:"id"`
	Name        string   `gorm:"unique_index;not null" json:"name"`
	Description string   `gorm:"not null" json:"description"`
	OTC         bool     `gorm:"not null" json:"otc"`
	RegionID    int      `gorm:"not null" json:"region_id"`
	Region      Region   `gorm:"foreign_key:RegionID" json:"-"`
	Category    Category `gorm:"foreignkey:CategoryID" json:"category"`
	CategoryID  int      `json:"-"`
	Tags        []Tag    `gorm:"many2many:medicine_tags" json:"tags"`
}

// CreateMedicine 创建药品记录
func CreateMedicine(medicine *Medicine) error {
	return DB.Create(medicine).Error
}

// GetMedicinesByRegion 获取特定区域内的药品
func GetMedicinesByRegion(regionID int) ([]Medicine, error) {
	var medicines []Medicine
	err := DB.Where("region_id = ?", regionID).Find(&medicines).Error
	if err != nil {
		return nil, errors.New("获取药品列表失败")
	}

	return medicines, nil
}

// GetMedicineByID 根据ID获取药品
func GetMedicineByID(id int) (*Medicine, error) {
	var medicine Medicine
	err := DB.Where("id = ?", id).First(&medicine).Error
	if err != nil {
		return nil, errors.New("药品不存在")
	}

	return &medicine, nil
}

func GetMedicineDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的药品ID"})
		return
	}

	medicine, err := models.GetMedicineByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "药品不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medicine": medicine})
}

// GetMedicineImages 根据药品 ID 获取药品图片
func GetMedicineImages(medicineID int) ([]MedicineImage, error) {
	var medicineImages []MedicineImage
	result := DB.Where("medicine_id = ?", medicineID).Find(&medicineImages)
	return medicineImages, result.Error
}

// SearchMedicines 根据关键词搜索药品
func SearchMedicines(query string, livingAreaID int) ([]Medicine, error) {
	var medicines []Medicine
	err := DB.Where("name LIKE ? AND living_area_id = ?", "%"+query+"%", livingAreaID).Find(&medicines).Error
	if err != nil {
		return nil, err
	}
	return medicines, nil
}

// UpdateMedicine 更新药品信息
func UpdateMedicine(medicine *Medicine) error {
	return DB.Save(medicine).Error
}

// DeleteMedicine 删除药品记录
func DeleteMedicine(id int) error {
	return DB.Delete(&Medicine{ID: id}).Error
}

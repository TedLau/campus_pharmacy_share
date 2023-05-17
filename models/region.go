package models

import "github.com/pkg/errors"

type Region struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Name        string `gorm:"unique_index;not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	InviteCode  string `gorm:"unique_index;not null" json:"invite_code"`
}

// CreateRegion 创建生活园区
func CreateRegion(region *Region) error {
	return DB.Create(region).Error
}

// GetRegions 获取所有生活园区
func GetRegions() ([]Region, error) {
	var regions []Region
	err := DB.Find(&regions).Error
	if err != nil {
		return nil, errors.New("获取生活园区列表失败")
	}

	return regions, nil
}

// GetRegionByID 根据ID获取生活园区
func GetRegionByID(id int) (*Region, error) {
	var region Region
	err := DB.Where("id = ?", id).First(&region).Error
	if err != nil {
		return nil, errors.New("生活园区不存在")
	}

	return &region, nil
}

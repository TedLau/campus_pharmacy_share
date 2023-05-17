package models

import (
	"time"
)

type LivingArea struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	InviteCode   string    `json:"invite_code"`
	UniversityID int       `json:"university_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CreateLivingArea 创建生活园区
func CreateLivingArea(livingArea *LivingArea) error {
	return DB.Create(livingArea).Error
}

// GetLivingAreaByID 根据ID获取生活园区
func GetLivingAreaByID(id int) (*LivingArea, error) {
	var livingArea LivingArea
	err := DB.First(&livingArea, id).Error
	return &livingArea, err
}

// UpdateLivingArea 更新生活园区
func UpdateLivingArea(livingArea *LivingArea) error {
	return DB.Save(livingArea).Error
}

// DeleteLivingArea 删除生活园区
func DeleteLivingArea(id int) error {
	return DB.Delete(&LivingArea{ID: id}).Error
}

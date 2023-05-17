package models

import (
	"time"
)

type Announcement struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	LivingAreaID int       `json:"living_area_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CreateAnnouncement 创建公告
func CreateAnnouncement(announcement *Announcement) error {
	return DB.Create(announcement).Error
}

// GetAnnouncementsByLivingAreaID 获取某个生活园区的所有公告
func GetAnnouncementsByLivingAreaID(livingAreaID int) ([]Announcement, error) {
	var announcements []Announcement
	err := DB.Where("living_area_id = ?", livingAreaID).Find(&announcements).Error
	return announcements, err
}

// UpdateAnnouncement 更新公告
func UpdateAnnouncement(announcement *Announcement) error {
	return DB.Save(announcement).Error
}

// DeleteAnnouncement 删除公告
func DeleteAnnouncement(id int) error {
	return DB.Delete(&Announcement{ID: id}).Error
}

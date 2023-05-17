package models

type Tag struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"unique_index;not null" json:"name"`
}

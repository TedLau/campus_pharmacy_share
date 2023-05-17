// 文件：models/post.go

package models

import (
	"errors"
)

type Post struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Contact   string `json:"contact"`
	UserID    int    `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// TableName 设置表名
func (Post) TableName() string {
	return "posts"
}

// CreatePost 创建帖子
func CreatePost(post *Post) error {
	if err := DB.Create(post).Error; err != nil {
		return errors.New("创建帖子失败")
	}
	return nil
}

// GetPostsByUserID 根据用户ID获取帖子
func GetPostsByUserID(userID int) ([]Post, error) {
	var posts []Post
	if err := DB.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, errors.New("获取帖子失败")
	}
	return posts, nil
}

// UpdatePost 更新帖子
func UpdatePost(post *Post) error {
	if err := DB.Save(post).Error; err != nil {
		return errors.New("更新帖子失败")
	}
	return nil
}

// DeletePostByID 根据ID删除帖子
func DeletePostByID(id int) error {
	if err := DB.Where("id = ?", id).Delete(Post{}).Error; err != nil {
		return errors.New("删除帖子失败")
	}
	return nil
}

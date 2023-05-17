package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRole string

const (
	RoleUser            UserRole = "user"
	RoleLivingAreaAdmin UserRole = "living_area_admin"
	RoleSuperAdmin      UserRole = "super_admin"
)

type User struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Username  string `gorm:"unique_index;not null" json:"username"`
	Password  string `gorm:"not null" json:"-"`
	Role      string `gorm:"type:enum('admin','manager','user');default:'user'" json:"role"`
	AvatarURL string `json:"avatar"`
	Nickname  string `json:"nickname"`
	Contact   string `json:"contact"`
	Email     string `json:"email"`
}

func CreateUser(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return DB.Create(user).Error
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	return &user, nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// UpdateUser 更新用户信息
func UpdateUser(user *User) error {
	if err := DB.Save(user).Error; err != nil {
		return errors.New("更新用户信息失败")
	}
	return nil
}

// GetUserByID 根据用户ID获取用户信息
func GetUserByID(id int) (*User, error) {
	var user User
	if err := DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user, nil
}

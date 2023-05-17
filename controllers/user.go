package controllers

import (
	"campus_pharmacy_share/models"
	"campus_pharmacy_share/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateUserInfo struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AvatarURL string `json:"avatar_url"`
}

type RegisterForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var form RegisterForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不正确"})
		return
	}

	user := &models.User{
		Username: form.Username,
		Password: form.Password,
		Role:     models.RoleUser,
	}

	if err := models.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败，请重试"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})

}

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不正确"})
		return
	}
	user, err := models.GetUserByUsername(form.Username)
	if err != nil || !user.CheckPassword(form.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	token, err := utils.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成令牌失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetUser 获取用户信息
func GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	user, err := models.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	var userInfo UpdateUserInfo
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt("id")
	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	if userInfo.Username != "" {
		user.Username = userInfo.Username
	}
	if userInfo.Email != "" {
		user.Email = userInfo.Email
	}
	if userInfo.Password != "" {
		hashedPassword, _ := models.HashPassword(userInfo.Password)
		user.Password = hashedPassword
	}
	if userInfo.AvatarURL != "" {
		user.AvatarURL = userInfo.AvatarURL
	}

	if err := models.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户信息失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功"})
}

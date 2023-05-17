// controllers/admin.go
package controllers

import (
	"campus_pharmacy_share/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ... Other handler functions

// PromoteUserToManager 提升用户为区域管理员
func PromoteUserToManager(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	err = models.PromoteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "提升用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户已成功提升为区域管理员"})
}

// GetAllMedicines 获取所有药品
func GetAllMedicines(c *gin.Context) {
	medicines, err := models.GetAllMedicines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取药品失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medicines": medicines})
}

// SearchMedicines 搜索药品
func SearchMedicines(c *gin.Context) {
	query := c.Query("q")
	medicines, err := models.SearchMedicines(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "搜索药品失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medicines": medicines})
}

// CreateAnnouncement 创建公告
func CreateAnnouncement(c *gin.Context) {
	var announcement models.Announcement
	if err := c.ShouldBindJSON(&announcement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	if err := models.CreateAnnouncement(&announcement); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建公告失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "公告创建成功", "announcement": announcement})
}

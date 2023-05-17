package controllers

import (
	"campus_pharmacy_share/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRegion 创建生活园区
func CreateRegion(c *gin.Context) {
	var region models.Region
	if err := c.ShouldBindJSON(&region); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不正确"})
		return
	}

	if err := models.CreateRegion(&region); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建生活园区失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "生活园区创建成功"})
}

// GetRegions 获取所有生活园区
func GetRegions(c *gin.Context) {
	regions, err := models.GetRegions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取生活园区列表失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, regions)
}

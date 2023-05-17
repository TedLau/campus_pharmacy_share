package controllers

import (
	"net/http"
	"strconv"

	"campus_pharmacy_share/models"
	"github.com/gin-gonic/gin"
)

// CreateLivingArea 创建生活园区
func CreateLivingArea(c *gin.Context) {
	var livingArea models.LivingArea
	if err := c.ShouldBindJSON(&livingArea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateLivingArea(&livingArea); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, livingArea)
}

// GetLivingAreaByID 根据ID获取生活园区信息
func GetLivingAreaByID(c *gin.Context) {
	id := c.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的园区ID"})
		return
	}
	livingArea, err := models.GetLivingAreaByID(ids)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "生活园区未找到"})
		return
	}

	c.JSON(http.StatusOK, livingArea)
}

// UpdateLivingArea 更新生活园区信息
func UpdateLivingArea(c *gin.Context) {
	var livingArea models.LivingArea
	if err := c.ShouldBindJSON(&livingArea); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateLivingArea(&livingArea); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, livingArea)
}

// DeleteLivingArea 删除生活园区
func DeleteLivingArea(c *gin.Context) {
	id := c.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的园区ID"})
		return
	}
	if err := models.DeleteLivingArea(ids); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除生活园区失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "生活园区删除成功"})
}

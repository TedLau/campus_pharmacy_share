package controllers

import (
	"campus_pharmacy_share/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateMedicine 创建药品
func CreateMedicine(c *gin.Context) {
	var medicine models.Medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不正确"})
		return
	}

	if err := models.CreateMedicine(&medicine); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建药品失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "药品创建成功"})
}

// GetMedicinesByRegion 获取特定区域内的药品
func GetMedicinesByRegion(c *gin.Context) {
	regionID := c.GetInt("region_id")
	medicines, err := models.GetMedicinesByRegion(regionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取药品列表失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, medicines)
}

// UpdateMedicine 更新药品信息
func UpdateMedicine(c *gin.Context) {
	var medicine models.Medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不正确"})
		return
	}

	// 更新药品信息
	if err := models.UpdateMedicine(&medicine); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新药品信息失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "药品信息更新成功"})
}

// DeleteMedicine 删除药品
func DeleteMedicine(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // 将字符串类型转换为整数类型
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的药品ID"})
		return
	}

	if err := models.DeleteMedicine(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除药品失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "药品删除成功"})
}

// SearchMedicines 搜索药品
func SearchMedicines(c *gin.Context) {
	query := c.Query("q")
	livingAreaID, _ := strconv.Atoi(c.Query("living_area_id"))
	medicines, err := models.SearchMedicines(query, livingAreaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "搜索药品失败，请重试"})
		return
	}
	c.JSON(http.StatusOK, medicines)
}

// GetMedicineImages 获取药品图片
func GetMedicineImages(c *gin.Context) {
	medicineID, _ := strconv.Atoi(c.Param("medicine_id"))
	images, err := models.GetMedicineImages(int(uint(medicineID)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取药品图片失败，请重试"})
		return
	}
	c.JSON(http.StatusOK, images)
}

// GetMedicineDetail 根据ID获取药品信息
func GetMedicineDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的药品ID"})
		return
	}

	medicine, err := models.GetMedicineByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "药品不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medicine": medicine})
}

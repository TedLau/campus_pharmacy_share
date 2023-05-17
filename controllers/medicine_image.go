package controllers

import (
	"net/http"
	"strconv"

	"campus_pharmacy_share/models"
	"github.com/gin-gonic/gin"
)

// CreateMedicineImage 创建药品图片
func CreateMedicineImage(c *gin.Context) {
	var medicineImage models.MedicineImage
	if err := c.ShouldBindJSON(&medicineImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateMedicineImage(&medicineImage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, medicineImage)
}

// GetMedicineImagesByMedicineID 根据药品ID获取药品图片列表
func GetMedicineImagesByMedicineID(c *gin.Context) {
	medicineIDs := c.Param("medicine_id")
	medicineID, err := strconv.Atoi(medicineIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的药品ID"})
		return
	}
	medicineImages, err := models.GetMedicineImagesByMedicineID(medicineID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "药品图片未找到"})
		return
	}

	c.JSON(http.StatusOK, medicineImages)
}

// DeleteMedicineImage 删除药品图片
func DeleteMedicineImage(c *gin.Context) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的图片ID"})
		return
	}
	if err := models.DeleteMedicineImage(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除药品图片失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "药品图片删除成功"})
}

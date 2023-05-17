// 文件：controllers/post.go

package controllers

import (
	"campus_pharmacy_share/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	_ "strconv"
)

// CreatePost 创建帖子
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.UserID = c.GetInt("id")
	if err := models.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建帖子失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "帖子创建成功"})
}

// GetPostsByUserID
// 文件：controllers/post.go

// ...

// GetPostsByUserID 根据用户ID获取帖子
func GetPostsByUserID(c *gin.Context) {
	userID := c.GetInt("id")

	posts, err := models.GetPostsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取帖子失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// UpdatePost 更新帖子
func UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的帖子ID"})
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.ID = id
	if err := models.UpdatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新帖子失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "帖子更新成功"})
}

// DeletePost 根据ID删除帖子
func DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的帖子ID"})
		return
	}

	if err := models.DeletePostByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除帖子失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "帖子删除成功"})
}

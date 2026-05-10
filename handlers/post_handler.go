package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oladejihenry/golang-crud-postgres/database"
	"github.com/oladejihenry/golang-crud-postgres/models"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post

	database.DB.Find(&posts)

	c.JSON(http.StatusOK, posts)
}

func GetPostByID(c *gin.Context) {
	var post models.Post

	if err := database.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	var newPost models.Post

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	if err := database.DB.Create(&newPost).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not create post"})
		return
	}

	// database.DB.Create(&newPost)

	c.JSON(http.StatusCreated, newPost)
}

func UpdatePost(c *gin.Context) {
	var post models.Post

	if err := database.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}

	var input models.Post

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	post.Title = input.Title
	post.Content = input.Content

	if err := database.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not update post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	var post models.Post

	if err := database.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}

	database.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

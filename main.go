package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oladejihenry/golang-crud-postgres/database"
	"github.com/oladejihenry/golang-crud-postgres/handlers"
)

func main() {
	database.ConnectDatabase()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	router.GET("/posts", handlers.GetPosts)
	router.GET("/posts/:id", handlers.GetPostByID)
	router.POST("/posts", handlers.CreatePost)
	router.PUT("/posts/:id", handlers.UpdatePost)
	router.DELETE("/posts/:id", handlers.DeletePost)

	router.Run(":8080")
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klaborda/blog-api/controllers"
	"github.com/klaborda/blog-api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello there from this goapi"})
	})

	r.GET("/posts", controllers.FindPosts)
	r.GET("/posts/:id", controllers.FindPost)
	r.POST("/posts", controllers.CreatePost)

	r.Run()
}

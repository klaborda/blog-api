package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klaborda/blog-api/models"
)

// GET /posts
// Get all posts
func FindPosts(c *gin.Context) {
	posts := []models.Post{}
	models.GetDB().Select(&posts, "SELECT * FROM posts ORDER BY title ASC")

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// GET /posts/:id
// Find a post
func FindPost(c *gin.Context) { // Get model if exist
	var post models.Post

	err := models.GetDB().Get(&post, "SELECT * FROM posts WHERE id=$1", c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "unable to find a post with that ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// POST /posts
// Create new post
func CreatePost(c *gin.Context) {
	// Validate input
	var input models.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create post
	// var err error
	res, err := models.GetDB().NamedExec(`INSERT INTO posts (title, author, content) VALUES(:title, :author, :content)`, &input)
	if err != nil {
		log.Fatalln(err)
	}
	post_id, err := res.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to determine post id"})
	}
	post := models.Post{ID: uint(post_id), Title: input.Title, Author: input.Author, Content: input.Content}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

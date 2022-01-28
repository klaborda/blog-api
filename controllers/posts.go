package controllers

import (
	"log"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/klaborda/blog-api/models"
	"github.com/mitchellh/mapstructure"
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
func FindPost(c *gin.Context) {
	var post models.Post

	err := models.GetDB().Get(&post, "SELECT * FROM posts WHERE id=$1", c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not found"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to determine post id"})
	}
	post := models.Post{ID: uint(post_id), Title: input.Title, Author: input.Author, Content: input.Content}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// PATCH /posts/:id
// Update a post
func UpdatePost(c *gin.Context) {
	var post models.Post
	var deltaStruct models.Post

	id := c.Param("id")
	err := models.GetDB().Get(&post, "SELECT * FROM posts WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not found"})
		return
	}

	var input models.UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deltaMap := findChangesInStruct(post, input)
	err = mapstructure.Decode(deltaMap, &deltaStruct)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = models.GetDB().NamedExec(`UPDATE posts SET title=:title, author=:author, content=:content WHERE id=:id`, &deltaStruct)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": deltaStruct})
}

// DELETE /posts/:id
// Delete a post
func DeletePost(c *gin.Context) {
	var post models.Post

	id := c.Param("id")
	err := models.GetDB().Get(&post, "SELECT * FROM posts WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not found"})
		return
	}

	_, err = models.GetDB().NamedQuery(`DELETE FROM posts WHERE id=:id`, map[string]interface{}{"id": id})
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}

func findChangesInStruct(before interface{}, after interface{}) map[string]interface{} {
	beforeMap := structs.Map(before)
	for _, f := range structs.Fields(after) {
		if f.Value() == "" {
			continue
		}
		beforeMap[f.Name()] = f.Value()
	}
	return beforeMap
}

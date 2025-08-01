package controllers

import (
	"crudApi/internals"
	"crudApi/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type postStruct struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// CreatePost
func CreatePost(c *gin.Context) {
	var post postStruct

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error: invalid request-": err.Error()})
		return
	}
	fmt.Println("post", post)
	postEntry := models.Post{Title: post.Title, Body: post.Body}
	result := internals.DB.Create(&postEntry)
	if result.Error != nil {
		fmt.Println("error", result.Error)
		c.JSON(500, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(200, gin.H{
		"post": postEntry,
	})
}

// GetPost
func GetPost(c *gin.Context) {
	var posts []models.Post
	searchResult := internals.DB.Find(&posts)
	if searchResult.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

// GetPostWithId
func GetPostWithId(c *gin.Context) {
	id := c.Param("id")
	log.Printf("id %v", id)

	var post models.Post
	result := internals.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch post for id"})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// UpdatePost
func UpdatePost(c *gin.Context) {
	// Get post id
	id := c.Param("id")

	// Get data of req body
	var newPost postStruct
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(400, gin.H{"error: invalid request-": err.Error()})
		return
	}

	// Get post
	var post models.Post
	result := internals.DB.Find(&post, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch post for id"})
		return
	}

	// Update post
	internals.DB.Model(&post).Updates(models.Post{Title: newPost.Title, Body: newPost.Body})

	c.JSON(200, gin.H{
		"post": post,
	})
}

// DeletePost
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	internals.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}

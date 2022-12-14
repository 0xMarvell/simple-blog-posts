package controllers

import (
	"fmt"
	"net/http"

	"github.com/0xMarvell/simple-blog-posts/pkg/config"
	"github.com/0xMarvell/simple-blog-posts/pkg/models"
	"github.com/0xMarvell/simple-blog-posts/pkg/utils"
	"github.com/gin-gonic/gin"
)

// CreatePost creates a new blog post.
func CreatePost(c *gin.Context) {
	var newPostPayload struct {
		Title  string `IndentedJSON:"title"`
		Body   string `IndentedJSON:"body"`
		Author string `IndentedJSON:"author"`
	}
	err := c.Bind(&newPostPayload)
	utils.CheckErr("c.Bind error: ", err)

	post := models.Post{
		Title:  newPostPayload.Title,
		Body:   newPostPayload.Body,
		Author: newPostPayload.Author,
	}
	result := config.DB.Create(&post)

	if (result.Error != nil) || (post.Author == "" || post.Body == "" || post.Title == "") {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Bad request: could not create new blog post",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"success": true,
		"post":    post,
	})
}

// GetPosts retrieves all existing blog posts.
func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)

	c.IndentedJSON(http.StatusOK, gin.H{
		"success": true,
		"posts":   posts,
	})
}

// GetPost retrieves single post based on its id.
func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	config.DB.First(&post, id)
	if !blogPostExists(id) {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   fmt.Sprintf("blog post with id {%v} does not exist", id),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"success": true,
		"post":    post,
	})
}

// UpdatePost updates an existing blog post.
func UpdatePost(c *gin.Context) {
	var updatedPostPayload struct {
		Title  string `IndentedJSON:"title"`
		Body   string `IndentedJSON:"body"`
		Author string `IndentedJSON:"author"`
	}
	var post models.Post
	id := c.Param("id")

	config.DB.First(&post, id)
	if !blogPostExists(id) {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   fmt.Sprintf("blog post with id {%v} does not exist", id),
		})
		return
	}

	err := c.Bind(&updatedPostPayload)
	utils.CheckErr("c.Bind error: ", err)

	config.DB.Model(&post).Updates(models.Post{
		Title:  updatedPostPayload.Title,
		Body:   updatedPostPayload.Body,
		Author: updatedPostPayload.Author,
	})

	c.IndentedJSON(http.StatusOK, gin.H{
		"success": true,
		"post":    post,
	})
}

// DeletePost deletes an existing post from the database
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	config.DB.First(&post, id)
	if !blogPostExists(id) {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   fmt.Sprintf("blog post with id {%v} does not exist", id),
		})
		return
	}
	config.DB.Delete(&post, id)

	c.IndentedJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "post deleted successfully",
	})
}

// blogPostExists checks if requested blog post exists in the database.
func blogPostExists(id string) bool {
	var post models.Post

	config.DB.First(&post, id)
	if post.ID == 0 {
		return false
	} else {
		return true
	}
}

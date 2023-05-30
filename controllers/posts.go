// controllers/post.go
package controllers

import (
    "net/http"
    "golang-blog/models"
    "github.com/gin-gonic/gin"
)

type CreatePostInput struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
    var input CreatePostInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    post := models.Post{Title: input.Title, Content: input.Content}
    models.DB.Create(&post)

    c.JSON(http.StatusOK, gin.H{"data": post})
}

// an endpoint to view every post created:
func FindPosts(c *gin.Context) {
    var posts []models.Post
    models.DB.Find(&posts)

    c.JSON(http.StatusOK, gin.H{"data": posts})
}

// a route that fetches only one specified post by URL param:
func FindPost(c *gin.Context) {
    var post models.Post

    if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    // DB.Where() lets you write SQL query commands, replacing the dynamic data with ? and passing the actual data as the second argument. DB.First(&post), like its name, selects the first name of the given collection of data and stores the result inside post.
// context.Param("<param-name>") is a Gin method to fetch the URL parameter by param name.
    c.JSON(http.StatusOK, gin.H{"data": post})
}

type UpdatePostInput struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
    var post models.Post
    if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
        return
    }

    var input UpdatePostInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedPost := models.Post{Title: input.Title, Content: input.Content}

    models.DB.Model(&post).Updates(&updatedPost)
    c.JSON(http.StatusOK, gin.H{"data": post})
}


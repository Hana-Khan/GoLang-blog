// main.go
package main

import (
	"golang-blog/controllers"
	"golang-blog/models"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    models.ConnectDatabase()  // new!

	router.POST("/posts", controllers.CreatePost)  // here!
    router.GET("/posts", controllers.FindPosts)
	router.GET("/posts/:id", controllers.FindPost)  // here!
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
    router.Run("localhost:8080")
}

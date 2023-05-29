// main.go
package main

import (
    

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    models.ConnectDatabase()  // new!

    // ...

    router.Run("localhost:8080")
}
// main.go - Simplified version
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    // Create Gin router
    r := gin.Default()
    
    // Simple hello world endpoint
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World!",
            "status":  "success",
        })
    })
    
    fmt.Println("Server starting on http://localhost:8080")
    r.Run(":8080")
}
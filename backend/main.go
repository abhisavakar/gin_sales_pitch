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
    
    // Simple hello world endpointdhrtjry
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World!",
            "status":  "success",
        })
    })
    
    fmt.Println("Server starting on http://localhost:8080")
    r.Run(":8080")
}

// This code sets up a simple web server using Gin that responds with a JSON message when accessed at the root URL.
// It listens on port 8080 and returns a JSON object with a message and status when the root endpoint is hit.
// To run this code, ensure you have the Gin framework installed and run the program.       
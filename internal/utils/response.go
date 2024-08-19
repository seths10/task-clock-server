package utils

import "github.com/gin-gonic/gin"

func SuccessResponse(c *gin.Context, data interface{}) {
    c.JSON(200, gin.H{"data": data})
}

func ErrorResponse(c *gin.Context, message string) {
    c.JSON(400, gin.H{"error": message})
}

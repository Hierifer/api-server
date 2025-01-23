package service

import (
	"github.com/gin-gonic/gin"
)

func ServiceAccessTV(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Access TV",
	})
}	
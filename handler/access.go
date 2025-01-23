package handler

import gin "github.com/gin-gonic/gin"

func GetAccessTV(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Access TV",
	})
}

func PostAccessTV(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Access TV",
	})
}
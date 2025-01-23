package v1

import (
	"api-server/handler"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	v1 := r.Group("v1")

	v1.GET("/access", handler.GetAccessTV)
	v1.POST("/access", handler.PostAccessTV)
}
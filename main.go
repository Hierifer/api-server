package main

import (
	v1 "api-server/api/v1" // Ensure this path is correct and the module exists

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/swaggo/gin-swagger/example/multiple/docs"
)

func main() {
	r := gin.Default()

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	// Register api/v1 endpoints
	v1.Register(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("v1")))

	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

package main

import (
	docs "PlayerWon/docs"
	"PlayerWon/endpoints/ads"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default() // Creates gin default engine

	// initialize API endpoints
	initializeEndpoints(r)

	r.Run()
}

func initializeEndpoints(r *gin.Engine) {

	// Project Endpoints
	ads.Routes(r)

	docs.SwaggerInfo.Description = "Challenge Documentation"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

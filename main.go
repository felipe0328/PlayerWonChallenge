package main

import (
	dalbase "PlayerWon/dal/dalBase"
	docs "PlayerWon/docs"
	"PlayerWon/endpoints/ads"
	"database/sql"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default() // Creates gin default engine

	// initialize API endpoints
	db := initializeEndpoints(r)
	defer db.Close()

	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func initializeEndpoints(r *gin.Engine) *sql.DB {

	// database
	db := dalbase.GetDB()
	// Project Endpoints
	ads.Routes(r, db)

	docs.SwaggerInfo.Description = "Challenge Documentation"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return db
}

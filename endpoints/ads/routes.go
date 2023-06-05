package ads

import (
	controller "PlayerWon/controllers/ads"
	"PlayerWon/dal/ads"
	adsservice "PlayerWon/services/adsService"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db *sql.DB) {

	// DAL
	adsDal := &ads.Ads{DB: db}

	// Service
	adsService := &adsservice.AdsService{}

	//Controller
	adsController := &controller.AdsController{
		AdsService: adsService,
		Dal:        adsDal,
	}

	// EndpointImplementation
	adsPost := PostAds{controller: adsController}

	///// Endpoints
	r.POST("/ads", adsPost.GetAd)
}

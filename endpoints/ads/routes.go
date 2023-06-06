package ads

import (
	controller "PlayerWon/controllers/ads"
	"PlayerWon/dal/ads"
	adsservice "PlayerWon/services/adsService"
	"PlayerWon/services/clock"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db *sql.DB) {

	// Services
	adsService := &adsservice.AdsService{}
	clock := &clock.Clock{}

	// DAL
	adsDal := &ads.Ads{DB: db, Clock: clock}

	//Controller
	adsController := &controller.AdsController{
		AdsService: adsService,
		Dal:        adsDal,
		Clock:      clock,
	}

	// EndpointImplementation
	adsPost := PostAds{Controller: adsController}

	///// Endpoints
	r.POST("/ads", adsPost.GetAd)
}

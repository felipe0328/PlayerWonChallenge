package ads

import (
	"PlayerWon/controllers/ads/models"
	"PlayerWon/endpoints/ads/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostAds struct{}

//	@Summary		Receive Ad
//	@Tags			Ads
//	@Description	Receive an ad based on input paramenters
//	@Produce		json
//
//	@Param			body	body		models.RequestAd	true	"Request Ad"
//	@Success		200		{object}	models.AdResponse
//	@Failure		400		{object}	error
//
//	@Router			/ads [post]
func (endpoint *PostAds) GetAd(c *gin.Context) {
	var requestAd models.RequestAd
	if err := c.ShouldBindJSON(&requestAd); err != nil {
		c.AbortWithError(http.StatusBadRequest, utils.AdRequestInvalid)
	}

}

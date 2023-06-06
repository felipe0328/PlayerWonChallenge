package ads

import (
	"PlayerWon/controllers/ads"
	"PlayerWon/endpoints/ads/utils"
	"PlayerWon/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostAds struct {
	Controller ads.IAdsController
}

// @Summary		Receive Ad
// @Tags			Ads
// @Description	Receive an ad based on input paramenters
// @Produce		json
//
// @Param			body	body		models.RequestAd	true	"Request Ad"
// @Success		200		{object}	models.AdResponse
// @Failure		400		{object}	error
//
// @Router			/ads [post]
func (endpoint *PostAds) GetAd(c *gin.Context) {
	var requestAd models.RequestAd
	if err := c.ShouldBindJSON(&requestAd); err != nil {
		c.String(http.StatusBadRequest, utils.AdRequestInvalid.Error())
		return
	}

	adData, err := endpoint.Controller.ObtainNewAd(requestAd)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, adData)

}

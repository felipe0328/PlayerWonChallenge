package ads

import (
	"PlayerWon/controllers/ads/models"
	"PlayerWon/controllers/ads/utils"
	adsservice "PlayerWon/services/adsService"
	adsServiceModel "PlayerWon/services/adsService/models"
)

type IAdsController interface {
	ObtainNewAd(requestData models.RequestAd) (models.AdResponse, error)
}

type AdsController struct {
	adsService adsservice.IAdsService
}

func (ac *AdsController) ObtainNewAd(requestData models.RequestAd) (models.AdResponse, error) {
	adResponse := models.AdResponse{}

	availableAds, err := ac.adsService.RequestAdFromServer()

	if err != nil {
		return adResponse, utils.UnableToGetAds
	}

	availableAd, err := GetAvaiableAdForUser(availableAds)

	if err != nil {
		return adResponse, utils.NoAvaiableAds
	}

	adResponse.ID = availableAd.ID
	adResponse.VideoURL = availableAd.VideoURL

	return adResponse, nil
}

func GetAvaiableAdForUser(availableAds []adsServiceModel.Ad) (adsServiceModel.Ad, error)

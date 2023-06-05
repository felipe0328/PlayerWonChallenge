package ads

import (
	"PlayerWon/controllers/ads/utils"
	"PlayerWon/dal/ads"
	"PlayerWon/models"
	adsservice "PlayerWon/services/adsService"
)

type IAdsController interface {
	ObtainNewAd(requestData models.RequestAd) (models.AdResponse, error)
}

type AdsController struct {
	AdsService adsservice.IAdsService
	Dal        ads.IAdsDal
}

func (ac *AdsController) ObtainNewAd(requestData models.RequestAd) (models.AdResponse, error) {
	adResponse := models.AdResponse{}

	availableAds, err := ac.AdsService.RequestAdFromServer()

	if err != nil {
		return adResponse, utils.UnableToGetAds
	}

	availableAd, err := ac.getAvaiableAdForUser(requestData.UserID, availableAds)

	if err != nil {
		return adResponse, utils.NoAvaiableAds
	}

	adResponse.ID = availableAd.ID
	adResponse.VideoURL = availableAd.VideoURL

	userVideo := &models.UserViewedAd{
		UserID:  requestData.UserID,
		VideoID: availableAd.ID,
	}

	go ac.Dal.RegisterNewAdForUser(*userVideo)

	return adResponse, nil
}

func (ac *AdsController) getAvaiableAdForUser(userID string, availableAds []models.Ad) (*models.Ad, error) {
	viewedVideos, err := ac.Dal.GetAdsViewedTodayByUser(userID)

	if err != nil {
		return nil, err
	}

	for _, adVideo := range availableAds {
		isVideoAvailable := true
		for _, viewedVideo := range viewedVideos {
			if viewedVideo == adVideo.ID {
				isVideoAvailable = false
			}
		}

		if isVideoAvailable {
			return &adVideo, nil
		}
	}

	return nil, utils.NoAvaiableAds
}

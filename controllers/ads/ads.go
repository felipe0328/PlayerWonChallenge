package ads

import (
	"PlayerWon/controllers/ads/utils"
	"PlayerWon/dal/ads"
	"PlayerWon/models"
	adsservice "PlayerWon/services/adsService"
	"PlayerWon/services/clock"

	"golang.org/x/exp/slices"
)

type IAdsController interface {
	ObtainNewAd(requestData models.RequestAd) (models.AdResponse, error)
}

type AdsController struct {
	AdsService adsservice.IAdsService
	Dal        ads.IAdsDal
	Clock      clock.IClock
}

func (ac *AdsController) ObtainNewAd(requestData models.RequestAd) (models.AdResponse, error) {
	adResponse := models.AdResponse{}

	availableAds, err := ac.AdsService.RequestAdFromServer()

	if err != nil {
		return adResponse, utils.UnableToGetAds
	}

	availableAd, err := ac.getAvaiableAdForUser(requestData, availableAds)

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

func (ac *AdsController) getAvaiableAdForUser(requestData models.RequestAd, availableAds []models.Ad) (*models.Ad, error) {
	viewedVideos, err := ac.Dal.GetAdsViewedTodayByUser(requestData.UserID)

	if err != nil {
		return nil, err
	}

	for _, adVideo := range availableAds {
		if !slices.Contains(viewedVideos, adVideo.ID) {
			if ac.isVideoInUserRegionAndLanguage(adVideo, requestData) {
				if ac.isVideoInWatchTime(adVideo) {
					return &adVideo, nil
				}
			}
		}
	}

	return nil, utils.NoAvaiableAds
}

func (ac *AdsController) isVideoInWatchTime(adVideo models.Ad) bool {
	currentTime := ac.Clock.Now().Hour()

	// this means the video frame ends the next day
	if adVideo.StartHour > adVideo.EndHour {
		if currentTime >= adVideo.StartHour || currentTime < adVideo.EndHour {
			return true
		}
	}

	if currentTime >= adVideo.StartHour &&
		currentTime < adVideo.EndHour {
		return true
	}

	return false
}

func (ac *AdsController) isVideoInUserRegionAndLanguage(adVideo models.Ad, requestData models.RequestAd) bool {
	if adVideo.Country == requestData.CountryCode &&
		adVideo.Lang == requestData.Language {
		return true
	}
	return false
}

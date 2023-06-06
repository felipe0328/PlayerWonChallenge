package test

import (
	"PlayerWon/controllers/ads"
	dalMock "PlayerWon/dal/ads/mocks"
	"PlayerWon/models"
	serviceMock "PlayerWon/services/adsService/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AdsControllerTestSuite struct {
	suite.Suite
	adsService *serviceMock.IAdsService
	dal        *dalMock.IAdsDal
	controller ads.AdsController
}

func TestAdsController(t *testing.T) {
	suite.Run(t, new(AdsControllerTestSuite))
}

func (t *AdsControllerTestSuite) SetupTest() {
	t.adsService = new(serviceMock.IAdsService)
	t.dal = new(dalMock.IAdsDal)
	t.controller = ads.AdsController{
		AdsService: t.adsService,
		Dal:        t.dal,
	}
}
func (t *AdsControllerTestSuite) TestObtainNewAd_Succesfull() {

	userID := "userID1234"

	requestData := models.RequestAd{
		UserID:      userID,
		Language:    "eng",
		CountryCode: "us",
	}

	adElementUS := models.Ad{
		ID:        "1234",
		VideoURL:  "www.video.com",
		Country:   "us",
		Lang:      "eng",
		StartHour: 1,
		EndHour:   10,
	}

	adElementFR := models.Ad{
		ID:        "5678",
		VideoURL:  "www.video.com",
		Country:   "ca",
		Lang:      "fre",
		StartHour: 1,
		EndHour:   10,
	}

	t.adsService.On("RequestAdFromServer").Return([]models.Ad{adElementUS, adElementFR})
	t.dal.On("GetAdsViewedTodayByUser", userID).Return([]string{}, nil)

	t.controller.ObtainNewAd()

}

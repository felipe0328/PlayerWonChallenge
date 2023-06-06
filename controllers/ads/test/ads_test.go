package test

import (
	"PlayerWon/controllers/ads"
	"PlayerWon/controllers/ads/utils"
	dalMock "PlayerWon/dal/ads/mocks"
	"PlayerWon/models"
	serviceMock "PlayerWon/services/adsService/mocks"
	clockMock "PlayerWon/services/clock/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AdsControllerTestSuite struct {
	suite.Suite
	adsService *serviceMock.IAdsService
	dal        *dalMock.IAdsDal
	clock      *clockMock.IClock
	controller ads.AdsController
}

func TestAdsController(t *testing.T) {
	suite.Run(t, new(AdsControllerTestSuite))
}

func (t *AdsControllerTestSuite) SetupTest() {
	t.adsService = new(serviceMock.IAdsService)
	t.dal = new(dalMock.IAdsDal)
	t.clock = new(clockMock.IClock)
	t.controller = ads.AdsController{
		AdsService: t.adsService,
		Dal:        t.dal,
		Clock:      t.clock,
	}

	t.dal.On("RegisterNewAdForUser", mock.Anything)
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
		VideoURL:  "www.video1.com",
		Country:   "us",
		Lang:      "eng",
		StartHour: 1,
		EndHour:   10,
	}

	adElementFR := models.Ad{
		ID:        "5678",
		VideoURL:  "www.video2.com",
		Country:   "ca",
		Lang:      "fre",
		StartHour: 1,
		EndHour:   10,
	}

	currentTime, _ := time.Parse(time.DateTime, "2023-01-01 05:00:00") // returning 5 am

	t.adsService.On("RequestAdFromServer").Return([]models.Ad{adElementUS, adElementFR}, nil)
	t.dal.On("GetAdsViewedTodayByUser", userID).Return([]string{}, nil)
	t.clock.On("Now").Return(currentTime)

	expectedResponse := models.AdResponse{
		ID:       "1234",
		VideoURL: "www.video1.com",
	}

	response, err := t.controller.ObtainNewAd(requestData)

	t.Nil(err)
	t.Equal(expectedResponse, response)
	t.adsService.AssertNumberOfCalls(t.T(), "RequestAdFromServer", 1)
	t.dal.AssertNumberOfCalls(t.T(), "GetAdsViewedTodayByUser", 1)
	t.clock.AssertNumberOfCalls(t.T(), "Now", 1)
}

func (t *AdsControllerTestSuite) TestObtainNewAd_NoAvailableCountry() {

	userID := "userID1234"

	requestData := models.RequestAd{
		UserID:      userID,
		Language:    "eng",
		CountryCode: "co",
	}

	adElementUS := models.Ad{
		ID:        "1234",
		VideoURL:  "www.video1.com",
		Country:   "us",
		Lang:      "eng",
		StartHour: 1,
		EndHour:   10,
	}

	adElementFR := models.Ad{
		ID:        "5678",
		VideoURL:  "www.video2.com",
		Country:   "ca",
		Lang:      "fre",
		StartHour: 1,
		EndHour:   10,
	}

	currentTime, _ := time.Parse(time.DateTime, "2023-01-01 05:00:00") // returning 5 am

	t.adsService.On("RequestAdFromServer").Return([]models.Ad{adElementUS, adElementFR}, nil)
	t.dal.On("GetAdsViewedTodayByUser", userID).Return([]string{}, nil)
	t.clock.On("Now").Return(currentTime)

	_, err := t.controller.ObtainNewAd(requestData)

	t.NotNil(err)
	t.Equal(utils.NoAvaiableAds, err)

	t.adsService.AssertNumberOfCalls(t.T(), "RequestAdFromServer", 1)
	t.dal.AssertNumberOfCalls(t.T(), "GetAdsViewedTodayByUser", 1)
	t.clock.AssertNotCalled(t.T(), "Now")
}

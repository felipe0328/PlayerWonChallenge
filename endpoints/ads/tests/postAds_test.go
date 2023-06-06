package tests

import (
	"PlayerWon/controllers/ads/mocks"
	"PlayerWon/controllers/ads/utils"
	"PlayerWon/endpoints/ads"
	endpointUtils "PlayerWon/endpoints/ads/utils"
	"PlayerWon/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type PostAdsTestSuite struct {
	suite.Suite
	controller *mocks.IAdsController
	endpoint   ads.PostAds
}

func TestPostAds(t *testing.T) {
	suite.Run(t, new(PostAdsTestSuite))
}

func (t *PostAdsTestSuite) SetupTest() {
	t.controller = new(mocks.IAdsController)
	t.endpoint = ads.PostAds{
		Controller: t.controller,
	}
}

func (t *PostAdsTestSuite) Test_GetAd_Success() {
	requestData := models.RequestAd{
		UserID:      "1234",
		Language:    "eng",
		CountryCode: "us",
	}

	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	bodyData, _ := json.Marshal(requestData)
	bytesBuffer := bytes.NewBuffer(bodyData)
	request := httptest.NewRequest(http.MethodPost, "/ads", bytesBuffer)
	c.Request = request

	obtainedAd := models.AdResponse{
		ID:       "12345",
		VideoURL: "www.amazingVideo.com",
	}

	t.controller.On("ObtainNewAd", requestData).Return(obtainedAd, nil)
	t.endpoint.GetAd(c)

	var response models.AdResponse
	json.Unmarshal(writer.Body.Bytes(), &response)

	t.Equal(http.StatusOK, writer.Code)
	t.Equal(obtainedAd, response)
	t.controller.AssertNumberOfCalls(t.T(), "ObtainNewAd", 1)
}

func (t *PostAdsTestSuite) Test_GetAd_NoAvailableVideo() {
	requestData := models.RequestAd{
		UserID:      "1234",
		Language:    "eng",
		CountryCode: "us",
	}

	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	bodyData, _ := json.Marshal(requestData)
	bytesBuffer := bytes.NewBuffer(bodyData)
	request := httptest.NewRequest(http.MethodPost, "/ads", bytesBuffer)
	c.Request = request

	obtainedAd := models.AdResponse{
		ID:       "12345",
		VideoURL: "www.amazingVideo.com",
	}

	t.controller.On("ObtainNewAd", requestData).Return(obtainedAd, utils.NoAvaiableAds)
	t.endpoint.GetAd(c)

	t.Equal(http.StatusBadRequest, writer.Code)
	t.Equal(utils.NoAvaiableAds.Error(), writer.Body.String())
	t.controller.AssertNumberOfCalls(t.T(), "ObtainNewAd", 1)
}

func (t *PostAdsTestSuite) Test_GetAd_InvalidInputData() {
	requestData := models.Ad{
		ID: "1234",
	}

	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	bodyData, _ := json.Marshal(requestData)
	bytesBuffer := bytes.NewBuffer(bodyData)
	request := httptest.NewRequest(http.MethodPost, "/ads", bytesBuffer)
	c.Request = request

	obtainedAd := models.AdResponse{
		ID:       "12345",
		VideoURL: "www.amazingVideo.com",
	}

	t.controller.On("ObtainNewAd", requestData).Return(obtainedAd, endpointUtils.AdRequestInvalid)
	t.endpoint.GetAd(c)

	t.Equal(http.StatusBadRequest, writer.Code)
	t.Equal(endpointUtils.AdRequestInvalid.Error(), writer.Body.String())
	t.controller.AssertNumberOfCalls(t.T(), "ObtainNewAd", 0)
}

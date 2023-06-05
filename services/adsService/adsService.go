package adsservice

import (
	"PlayerWon/models"
	"encoding/json"
	"net/http"
)

type IAdsService interface {
	RequestAdFromServer() ([]models.Ad, error)
}

type AdsService struct{}

var adServerURL string = "https://gist.githubusercontent.com/victorhurdugaci/22a682eb508e65d97bd5b9152f564ab3/raw/dbf27ef217dba9bbd753de26cdabf8a91bdf1550/sm_ads.json"

func (as *AdsService) RequestAdFromServer() ([]models.Ad, error) {
	resp, err := http.Get(adServerURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var adsData models.AdList
	if err := json.NewDecoder(resp.Body).Decode(&adsData); err != nil {
		return nil, err
	}

	return adsData.Ads, nil
}

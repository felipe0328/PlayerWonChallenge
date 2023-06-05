package ads

import (
	"PlayerWon/models"
	"database/sql"
	"fmt"
	"time"
)

type IAdsDal interface {
	RegisterNewAdForUser(userAdData models.UserViewedAd)
	GetAdsViewedTodayByUser(userID string) ([]string, error)
}

type Ads struct {
	*sql.DB
}

func (ad *Ads) RegisterNewAdForUser(userData models.UserViewedAd) {
	ad.QueryRow("INSERT into userAds (userID, videoID, timestamp) VALUES (?,?,?);", userData.UserID, userData.VideoID, time.Now().Format(time.RFC3339))
}

func (ad *Ads) GetAdsViewedTodayByUser(userID string) ([]string, error) {

	// This query was created to do:
	// SELECT videoID from userAds where userID = $1 and DATE(timestamp) = CURDATE();
	// but the version of ramsql don't work correctly with times so storting the time as string to parse after scan
	query := `SELECT videoID, timestamp from userAds where userID = $1;`
	rows, err := ad.Query(query, userID)

	if err != nil {
		return nil, err
	}

	videoIDs := make([]string, 0)

	for rows.Next() {
		var videoID string
		var timestamp string
		if err := rows.Scan(&videoID, &timestamp); err != nil {
			fmt.Println("OFP::1 ", err)
			return nil, err
		}

		entryTimestamp, err := time.Parse(time.RFC3339, timestamp)

		if err != nil {
			return nil, err
		}

		timeNowMinusDay := time.Now().Add(-time.Hour * 24)

		if entryTimestamp.After(timeNowMinusDay) {
			videoIDs = append(videoIDs, videoID)
		}

	}

	return videoIDs, nil
}

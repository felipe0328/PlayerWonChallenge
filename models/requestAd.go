package models

type RequestAd struct {
	UserID      string `binding:"required"`
	Language    string `binding:"required"`
	CountryCode string `binding:"required"`
}

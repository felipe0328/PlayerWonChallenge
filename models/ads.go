package models

type Ad struct {
	ID        string `json:"id"`
	VideoURL  string `json:"video_url"`
	Country   string `json:"country"`
	Lang      string `json:"lang"`
	StartHour int    `json:"start_hour"`
	EndHour   int    `json:"end_hour"`
}

type AdList struct {
	Ads []Ad `json:"ads"`
}

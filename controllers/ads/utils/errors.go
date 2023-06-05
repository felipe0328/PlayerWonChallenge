package utils

import "errors"

var (
	UnableToGetAds = errors.New("unable to get ads from server.")
	NoAvaiableAds  = errors.New("there is not avaiable ads for user.")
)

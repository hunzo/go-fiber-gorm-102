package models

import (
	"gorm.io/gorm"
)

type TOKEN struct {
	*gorm.Model
	ACCESSTOKEN  string `json:"access_token`
	REFRESHTOKEN string `json:"access_token`
}

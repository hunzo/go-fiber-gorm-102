package services

import (
	"log"

	"github.com/hunzo/go-fiber-gorm-102/models"
	"gorm.io/gorm"
)

var DBConnect *gorm.DB

func CheckData() (models.TOKEN, error) {

	db := DBConnect
	var ret models.TOKEN
	err := db.First(&ret)
	if err != nil {
		log.Println(err.Error)

		var t models.TOKEN
		return t, err.Error
	}
	return ret, nil
}

func CreateData() {
	db := DBConnect
	db.Model(&models.TOKEN{}).Create(map[string]interface{}{
		"ACCESSTOKEN":  "test access token",
		"REFRESHTOKEN": "test refresh token",
	})
}

func DeleteData() {
	db := DBConnect
	var m models.TOKEN
	// db.Delete(&m)
	db.Where("1 = 1").Delete(&m)
}

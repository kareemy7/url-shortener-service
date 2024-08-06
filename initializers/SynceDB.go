package initializers

import "url-shortener3/models"

func SynceDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.URL{})
}

package initializers

import "userAuth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}

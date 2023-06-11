package initializer

import "starterkit-go-auth/models"

func Migration()  {
	DB.AutoMigrate(&models.User{})
}
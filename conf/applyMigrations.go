package conf

import "github.com/awbw/2040/models"

func ApplyMigrations() {
	DB.AutoMigrate(&models.Unit{})
	DB.AutoMigrate(&models.UnitType{})
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Map struct {
	gorm.Model
	Name               string
	Players            int
	Type               string
	CreatorID          int
	Creator            User `gorm:"foreignKey: CreatorID"`
	Published          bool
	PublishedDate      time.Time
	FirstPublishedDate time.Time
	LastViewed         time.Time
}

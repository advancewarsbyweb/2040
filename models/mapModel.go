package models

import "time"

type Map struct {
	ID                 int        `db:"maps_id"`
	Name               string     `db:"maps_name"`
	Players            int        `db:"maps_players"`
	Type               string     `db:"maps_type"`
	CreatorID          int        `db:"maps_users_id"`
	Published          string     `db:"maps_published"`
	PublishedDate      *time.Time `db:"maps_published_date"`
	FirstPublishedDate *time.Time `db:"maps_first_published_date"`
	AwnNumber          int        `db:"maps_awn_number"`
	LastViewed         *time.Time `db:"maps_last_viewed"`
}

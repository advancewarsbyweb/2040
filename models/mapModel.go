package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Map struct {
	ID                 int       `db:"maps_id"`
	Name               string    `db:"maps_name"`
	Players            int       `db:"maps_players"`
	Type               string    `db:"maps_type"`
	CreatorID          int       `db:"maps_users_id"`
	Published          string    `db:"maps_published"`
	publishedDate      null.Time `db:"maps_published_date"`
	firstPublishedDate null.Time `db:"maps_first_published_date"`
	AwnNumber          int       `db:"maps_awn_number"`
	pastViewed         null.Time `db:"maps_last_viewed"`
}

func (m *Map) PublishedDate() time.Time {
	return m.publishedDate.Time
}

func (m *Map) FirstPublishedDate() time.Time {
	return m.firstPublishedDate.Time
}

func (m *Map) PastViewed() time.Time {
	return m.pastViewed.Time
}

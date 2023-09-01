package types

import (
	"time"

	"github.com/awbw/2040/models"
)

type User struct {
	models.User
	DiscordID      string
	SessionID      string
	ActivityDate   time.Time
	FriendDate     time.Time
	ChangelogDate  time.Time
	LastPage       string
	YourGamesDate  time.Time
	TournamentDate time.Time
	LastVacation   time.Time
}

func NewUser(u models.User) User {
	return User{
		User:           u,
		DiscordID:      u.DiscordID.String,
		SessionID:      u.SessionID.String,
		ActivityDate:   u.ActivityDate.Time,
		FriendDate:     u.FriendDate.Time,
		ChangelogDate:  u.ChangelogDate.Time,
		LastPage:       u.LastPage.String,
		YourGamesDate:  u.YourGamesDate.Time,
		TournamentDate: u.TournamentDate.Time,
		LastVacation:   u.LastVacation.Time,
	}
}

package db

import (
	"time"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	"github.com/bxcodec/faker/v4"
	"gopkg.in/guregu/null.v4"
)

type UserFactory struct {
	User types.User
}

var User UserFactory

func NewUserFactory() UserFactory {
	return UserFactory{}
}

func init() {
	User = NewUserFactory()
}

func (f *UserFactory) Create() *UserFactory {
	f.User = types.NewUser(models.User{
		Username:       faker.Word(),
		FirstName:      faker.Word(),
		LastName:       faker.Word(),
		Password:       "password",
		Email:          faker.Email(),
		EmailDisplay:   "N",
		Inactive:       0,
		Games:          0,
		UniqID:         faker.UUIDDigit(),
		DiscordID:      null.StringFrom(faker.UUIDDigit()),
		SessionID:      null.StringFrom(faker.UUIDDigit()),
		ActivityDate:   null.TimeFrom(time.Now()),
		BootRecord:     0,
		ShoalsDisplay:  "new",
		CountryID:      3,
		ImageDir:       "terrain/aw2",
		FriendDate:     null.TimeFrom(time.Now()),
		AutoScroll:     0,
		ChangelogDate:  null.TimeFrom(time.Now()),
		IP:             "0.0.0.32",
		LastPage:       null.StringFrom("game.php"),
		Moderator:      "N",
		Days:           0,
		YourGamesDate:  null.TimeFrom(time.Now()),
		TurnEmail:      "Y",
		Timezone:       "-4",
		EmailMessages:  "Y",
		ShowWarnings:   0,
		CoPortraits:    "ds",
		TournamentDate: null.TimeFrom(time.Now()),
		Gridlines:      "N",
		Vacation:       0,
		LastVacation:   null.Time{},
		MapCommittee:   0,
		DorFC:          null.String{},
		GameAnimations: "Y",
	})
	return f
}

func (f *UserFactory) Build() types.User {
	return f.User
}

func (f *UserFactory) BuildInsert() types.User {
	uID, _ := db.UserRepo.CreateUser(f.User)
	f.User.ID = uID
	return f.User
}

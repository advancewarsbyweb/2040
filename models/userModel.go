package models

import "gopkg.in/guregu/null.v4"

type User struct {
	ID             int         `db:"users_id"`
	Username       string      `db:"users_username"`
	FirstName      string      `db:"users_firstname"`
	LastName       string      `db:"users_lastname"`
	Password       string      `db:"users_password"`
	Email          string      `db:"users_email"`
	EmailDisplay   string      `db:"users_email_display"`
	Inactive       int         `db:"users_inactive"`
	Games          int         `db:"users_games"`
	UniqID         string      `db:"users_uniq_id"`
	DiscordID      null.String `db:"users_discord_id"`
	SessionID      null.String `db:"users_session_id"`
	ActivityDate   null.Time   `db:"users_activity_date"`
	BootRecord     int         `db:"users_boot_record"`
	ShoalsDisplay  string      `db:"users_shoals_display"`
	CountryID      int         `db:"users_countries_id"`
	ImageDir       string      `db:"users_awbw_image_dir"`
	FriendDate     null.Time   `db:"users_friend_date"`
	AutoScroll     int         `db:"users_auto_scroll"`
	ChangelogDate  null.Time   `db:"users_changelog_date"`
	IP             string      `db:"ip"`
	LastPage       null.String `db:"users_awbw_last_page"`
	Moderator      string      `db:"users_awbw_moderator"`
	Days           int         `db:"users_awbw_days"`
	YourGamesDate  null.Time   `db:"users_yourgames_date"`
	TurnEmail      string      `db:"users_awbw_turn_email"`
	Timezone       string      `db:"users_timezone"`
	EmailMessages  string      `db:"users_email_messages"`
	ShowWarnings   int         `db:"users_show_warnings"`
	CoPortraits    string      `db:"users_co_portraits"`
	TournamentDate null.Time   `db:"users_tournament_date"`
	Gridlines      string      `db:"users_gridlines"`
	Vacation       int         `db:"users_vacation"`
	LastVacation   null.Time   `db:"users_last_vacation"`
	MapCommittee   int         `db:"users_map_committee"`
	DorFC          null.String `db:"users_dor_fc"`
	GameAnimations bool        `db:"users_game_animations"`
	UserUnusedFields
}

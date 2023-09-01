package models

type PressTo struct {
	ID       int `db:"press_to_id"`
	PressID  int `db:"press_to_press_id,omitempty"`
	PlayerID int `db:"press_to_players_id"`
	Username int `db:"users_username"`
}

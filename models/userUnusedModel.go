package models

// Legacy unused fields in ofua_users
type UserUnusedFields struct {
	PnName           string `db:"pn_name" default:" "`
	PnUname          string `db:"pn_uname" default:" "`
	PnEmail          string `db:"pn_email" default:" "`
	PnFEmail         string `db:"pn_femail" default:" "`
	PnUrl            string `db:"pn_url" default:" "`
	PnUserAvatar     string `db:"pn_user_avatar" default:" "`
	PnUserRegdate    string `db:"pn_user_regdate" default:" "`
	PnUserIcq        string `db:"pn_user_icq" default:" "`
	PnUserOcc        string `db:"pn_user_occ" default:" "`
	PnUserFrom       string `db:"pn_user_from" default:" "`
	PnUserIntrest    string `db:"pn_user_intrest" default:" "`
	PnUserSig        string `db:"pn_user_sig" default:" "`
	PnUserViewEmail  string `db:"pn_user_viewemail" default:" "`
	PnUserTheme      string `db:"pn_user_theme" default:" "`
	PnUserAim        string `db:"pn_user_aim" default:" "`
	PnUserYim        string `db:"pn_user_yim" default:" "`
	PnUserMsnm       string `db:"pn_user_msnm" default:" "`
	PnPass           string `db:"pn_pass" default:" "`
	PnStorynum       string `db:"pn_storynum" default:" "`
	PnUMode          string `db:"pn_umode" default:" "`
	PnUOrder         string `db:"pn_uorder" default:" "`
	PnTHold          string `db:"pn_thold" default:" "`
	PnNoScore        string `db:"pn_noscore" default:" "`
	PnBio            string `db:"pn_bio" default:" "`
	PnUblockOn       string `db:"pn_ublockon" default:" "`
	PnUblock         string `db:"pn_ublock" default:" "`
	PnTheme          string `db:"pn_theme" default:" "`
	PnCommentMax     string `db:"pn_commentmax" default:" "`
	PnCounter        string `db:"pn_counter" default:" "`
	PnTimezoneOffset string `db:"pn_timezone_offset" default:" "`
	PnUid            string `db:"pn_uid" default:" "`
}

package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

// Nullable Bool that overrides sql.NullBool
type NullBool struct {
	sql.NullBool
}

func (nb NullBool) MarshalJSON() ([]byte, error) {
	if nb.Valid {
		return json.Marshal(nb.Bool)
	}
	return json.Marshal(nil)
}

func (nb *NullBool) UnmarshalJSON(data []byte) error {
	var b *bool
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	if b != nil {
		nb.Valid = true
		nb.Bool = *b
	} else {
		nb.Valid = false
	}
	return nil
}

// Nullable Float64 that overrides sql.NullFloat64
type NullString struct {
	sql.NullString
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		ns.Valid = true
		ns.String = *s
	} else {
		ns.Valid = false
	}
	return nil
}

// Nullable Int64 that overrides sql.NullInt64
type NullInt16 struct {
	sql.NullInt16
}

func (ni NullInt16) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int16)
	}
	return json.Marshal(nil)
}

func (ni *NullInt16) UnmarshalJSON(data []byte) error {
	var i *int16
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	if i != nil {
		ni.Valid = true
		ni.Int16 = *i
	} else {
		ni.Valid = false
	}
	return nil
}

// Nullable String that overrides sql.NullString
type NullTime struct {
	sql.NullTime
}

func (ns NullTime) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Time)
	}
	return json.Marshal(nil)
}

func (ns *NullTime) UnmarshalJSON(data []byte) error {
	var t *time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	if t != nil {
		ns.Valid = true
		ns.Time = *t
	} else {
		ns.Valid = false
	}
	return nil
}

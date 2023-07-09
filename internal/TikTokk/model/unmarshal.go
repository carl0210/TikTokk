package model

import (
	"database/sql"
	"strings"
	"time"
)

type ConvertTime time.Time
type ConvertNullTime sql.NullTime
type ConvertBool bool

func (c *ConvertBool) UnmarshalJSON(b []byte) error {
	str := strings.ToLower(strings.Trim(string(b), `""`))
	if str == "1" || str == "true" {
		*c = true
	} else {
		*c = false
	}
	return nil
}

func (c *ConvertTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `""`)
	time, err := time.Parse("2006-01-02 15:04:05.000", str)
	if err != nil {
		return err
	}
	*c = ConvertTime(time)
	return nil
}

func (c *ConvertNullTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `""`)
	if str == "" {
		*c = ConvertNullTime{Time: time.Time{}, Valid: false}
		return nil
	}
	time, err := time.Parse("2006-01-02 15:04:05.000", str)
	if err != nil {
		return err
	}
	*c = ConvertNullTime{Time: time, Valid: true}
	return nil
}

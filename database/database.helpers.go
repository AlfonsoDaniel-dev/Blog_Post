package database

import (
	"database/sql"
	"time"
)

func StringToNull(s string) sql.NullString {
	null := sql.NullString{
		String: s,
	}

	if null.String == "" {
		null.Valid = true
	}

	return null
}

func TimeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{
		Time: t,
	}

	if !null.Time.IsZero() {
		null.Valid = true
	}

	return null
}

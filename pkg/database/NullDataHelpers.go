package database

import (
	"database/sql"
	"time"
)

func TimeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{}
	if !null.Time.IsZero() {
		null.Valid = true
	}

	return null
}

func StringToNull(s string) sql.NullString {
	null := sql.NullString{
		String: s,
	}
	if null.String != "" {
		null.Valid = true
	}

	return null
}

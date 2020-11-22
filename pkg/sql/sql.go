package sql

import (
	"database/sql"
	"errors"
	"os"
	"strconv"
	"time"
)

func Connect() (*sql.DB, error) {
	_, portErr := strconv.Atoi(os.Getenv("DB_PORT"))
	if portErr != nil {
		return nil, errors.New("DB_PORT must be numeric")
	}
	if os.Getenv("DB_HOST") == "" {
		return nil, errors.New("DB_HOST required")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		return nil, errors.New("DB_PASSWORD required")
	}
	if os.Getenv("DB_USER") == "" {
		return nil, errors.New("DB_USER required")
	}
	conn, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?parseTime=true")
	if err != nil {
		panic("Cannot connect to DB")
	}
	return conn, err
}

func ConvertInt(value int) sql.NullInt64 {
	if value != 0 {
		return sql.NullInt64{
			Valid: true,
			Int64: int64(value),
		}
	}
	return sql.NullInt64{
		Valid: false,
	}
}

func ConvertString(value string) sql.NullString {
	if value != "" {
		return sql.NullString{
			Valid:  true,
			String: value,
		}
	}
	return sql.NullString{
		Valid: false,
	}
}

func ConvertTime(value time.Time) sql.NullTime {
	if !value.IsZero() {
		return sql.NullTime{
			Valid: true,
			Time:  value,
		}
	}
	return sql.NullTime{
		Valid: false,
	}
}

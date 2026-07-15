package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func PGTextParse(s string) pgtype.Text {
	return pgtype.Text{
		String: s,
		Valid:  s != "",
	}
}

func PGInt4Parse(n int) pgtype.Int4 {
	return pgtype.Int4{
		Int32: int32(n),
		Valid: n != 0,
	}
}

func PGDateParse(d string) pgtype.Date {
	date, err := time.Parse("01-2006", d)
	return pgtype.Date{
		Time:  date,
		Valid: err == nil,
	}
}

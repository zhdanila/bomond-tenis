package query

import "github.com/go-openapi/strfmt"

type GetCourtsQuery struct{}

type BookCourtQuery struct {
	CourtId  string      `db:"court_id"`
	Date     strfmt.Date `json:"date,omitempty"`
	Duration int64       `json:"duration,omitempty"`
	Time     string      `json:"time,omitempty"`
	UserID   string      `json:"userId,omitempty"`
}

type GetBookedCourtQuery struct {
	CourtId string `db:"court_id"`
}

type CancelCourtBookingQuery struct {
	CourtId string `db:"court_id"`
	BookId  string `db:"book_id"`
}

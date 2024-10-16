package query

import "github.com/go-openapi/strfmt"

type Court struct {
	CourtId string `db:"id"`
	Name    string `db:"name"`
}

type GetCourtsQuery struct {
	Out struct {
		Courts []Court `json:"courts"`
	}
}

type BookCourtQuery struct {
	CourtId  string      `db:"court_id"`
	Date     strfmt.Date `json:"date,omitempty"`
	Duration int64       `json:"duration,omitempty"`
	Time     string      `json:"time,omitempty"`
	UserID   string      `json:"userId,omitempty"`
	Out      struct {
		ID int `json:"id"`
	}
}

type GetBookedCourtQuery struct {
	CourtId string `db:"court_id"`
	Out     struct {
		Court Court `json:"court"`
	}
}

type CancelCourtBookingQuery struct {
	CourtId string `db:"court_id"`
	BookId  string `db:"book_id"`
}

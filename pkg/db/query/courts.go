package query

import "github.com/go-openapi/strfmt"

type Court struct {
	CourtId string `db:"id"`
	Name    string `db:"name"`
}

type BookedCourt struct {
	ID       string      `json:"id"`
	CourtId  string      `db:"court_id"`
	UserID   string      `json:"userId,omitempty"`
	Date     strfmt.Date `json:"date,omitempty"`
	Duration int64       `json:"duration,omitempty"`
	Time     string      `json:"time,omitempty"`
}

type GetCourtsQuery struct {
	Out struct {
		Courts []Court `json:"courts"`
	}
}

type BookCourtQuery struct {
	BookCourt BookedCourt `json:"bookCourt"`
	Out       struct {
		ID int `json:"id"`
	}
}

type GetBookedCourtQuery struct {
	CourtId string `db:"court_id"`
	Out     struct {
		Court        Court                 `json:"court"`
		BookedCourts []BookedCourtResponse `json:"bookedCourt"`
	}
}

type BookedCourtResponse struct {
	Account Account     `json:"account"`
	Court   BookedCourt `json:"court"`
}

type CourtsWithAccountsResponse struct {
	Responses []BookedCourtResponse `json:"responses"`
}

type CancelCourtBookingQuery struct {
	CourtId string `db:"court_id"`
	BookId  string `db:"book_id"`
}

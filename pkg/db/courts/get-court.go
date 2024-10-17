package courts

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	sql2 "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type GetCourtHandler struct {
	pool *sqlx.DB
}

func NewGetBookedCourtHandler(pool *sqlx.DB) *GetCourtHandler {
	return &GetCourtHandler{pool: pool}
}

func (h *GetCourtHandler) Exec(ctx context.Context, args *query.GetBookedCourtQuery) (err error) {
	sql := `SELECT c.id, b.user_id, b.date, b.duration, b.time, c.name 
            FROM courts c
            LEFT JOIN booked_court b ON c.id = b.court_id
            WHERE c.id = $1`

	rows, err := h.pool.Query(sql, args.CourtId)
	if err != nil {
		return err
	}
	defer rows.Close()

	var courtsWithAccounts []query.BookedCourtResponse
	var court query.Court

	for rows.Next() {
		var bookedCourt query.BookedCourt
		var userID, duration sql2.NullInt64
		var time sql2.NullString

		if err = rows.Scan(&court.CourtId, &userID, &bookedCourt.Date, &duration, &time, &court.Name); err != nil {
			return err
		}

		if userID.Valid {
			bookedCourt.UserID = strconv.FormatInt(userID.Int64, 10)
		} else {
			bookedCourt.UserID = strconv.Itoa(0)
		}

		if duration.Valid {
			bookedCourt.Duration = duration.Int64
		} else {
			bookedCourt.Duration = 0
		}

		if time.Valid {
			bookedCourt.Time = time.String
		} else {
			bookedCourt.Time = ""
		}

		account, err := getAccount(h.pool, int(userID.Int64))
		if err != nil {
			return err
		}

		courtsWithAccounts = append(courtsWithAccounts, query.BookedCourtResponse{
			Account: account,
			Court:   bookedCourt,
		})
	}

	if err = rows.Err(); err != nil {
		return err
	}

	args.Out.Court = court
	args.Out.BookedCourts = courtsWithAccounts

	return nil
}

func getAccount(pool *sqlx.DB, userID int) (query.Account, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", utils.UsersTable)

	rows := pool.QueryRow(sql, userID)

	var account query.Account
	if err := rows.Scan(&account.ID, &account.Email, &account.Password, &account.Username); err != nil {
		return query.Account{}, err
	}

	return account, nil
}

func (h *GetCourtHandler) Context() interface{} {
	return (*query.GetBookedCourtQuery)(nil)
}

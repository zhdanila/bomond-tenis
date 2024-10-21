package courts

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BookCourtHandler struct {
	pool *sqlx.DB
}

func NewBookCourtHandler(pool *sqlx.DB) *BookCourtHandler {
	return &BookCourtHandler{pool: pool}
}

func (h *BookCourtHandler) Exec(ctx context.Context, args *query.BookCourtQuery) (err error) {
	sql := fmt.Sprintf(`INSERT INTO %s (court_id, user_id, date, duration)
		VALUES ($1, $2, $3, $4) RETURNING id`, utils.BookedCourtTable)

	rows := h.pool.QueryRow(sql, args.BookCourt.CourtId, args.BookCourt.UserID,
		args.BookCourt.Date, args.BookCourt.Duration)

	var id int
	if err := rows.Scan(&id); err != nil {
		return err
	}

	args.Out.ID = id

	return nil
}

func (h *BookCourtHandler) Context() interface{} {
	return (*query.BookCourtQuery)(nil)
}

package courts

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type bookCourtHandler struct {
	pool *sqlx.DB
}

func NewBookCourtHandler(pool *sqlx.DB) *bookCourtHandler {
	return &bookCourtHandler{pool: pool}
}

func (h *bookCourtHandler) Exec(ctx context.Context, args *query.BookCourtQuery) (err error) {
	sql := fmt.Sprintf(`INSERT INTO %s (court_id, user_id, date, duration, time)
		VALUES ($1, $2, $3, $4, $5) RETURNING court_id`, utils.BookedCourtTable)

	rows := h.pool.QueryRow(sql, args.CourtId, args.UserID, args.Date, args.Duration, args.Time)

	var id int
	if err := rows.Scan(&id); err != nil {
		return err
	}

	args.Out.ID = id

	return nil
}

func (h *bookCourtHandler) Context() interface{} {
	return (*query.BookCourtQuery)(nil)
}

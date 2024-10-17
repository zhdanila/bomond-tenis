package courts

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CancelCourtBookingQueryHandler struct {
	pool *sqlx.DB
}

func NewCancelCourtBookingQueryHandler(pool *sqlx.DB) *CancelCourtBookingQueryHandler {
	return &CancelCourtBookingQueryHandler{pool: pool}
}

func (h *CancelCourtBookingQueryHandler) Exec(ctx context.Context, args *query.CancelCourtBookingQuery) (err error) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = $1", utils.BookedCourtTable)

	_, err = h.pool.Exec(sql, args.BookId)
	if err != nil {
		return err
	}

	return nil
}

func (h *CancelCourtBookingQueryHandler) Context() interface{} {
	return (*query.CancelCourtBookingQuery)(nil)
}

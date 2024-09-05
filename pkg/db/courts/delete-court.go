package courts

import (
	"bomond-tenis/pkg/db/query"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type cancelCourtBookingQueryHandler struct {
	pool *sqlx.DB
}

func NewCancelCourtBookingQueryHandler(pool *sqlx.DB) *cancelCourtBookingQueryHandler {
	return &cancelCourtBookingQueryHandler{pool: pool}
}

func (h *cancelCourtBookingQueryHandler) Exec(ctx context.Context, args *query.CancelCourtBookingQuery) (err error) {
	fmt.Println("cancel booking court")

	return nil
}

func (h *cancelCourtBookingQueryHandler) Context() interface{} {
	return (*query.CancelCourtBookingQuery)(nil)
}

package courts

import (
	"bomond-tenis/pkg/db/query"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type getBookedCourtHandler struct {
	pool *sqlx.DB
}

func NewGetBookedCourtHandler(pool *sqlx.DB) *getBookedCourtHandler {
	return &getBookedCourtHandler{pool: pool}
}

func (h *getBookedCourtHandler) Exec(ctx context.Context, args *query.GetBookedCourtQuery) (err error) {
	fmt.Println("get booked court")

	return nil
}

func (h *getBookedCourtHandler) Context() interface{} {
	return (*query.GetBookedCourtQuery)(nil)
}

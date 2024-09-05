package courts

import (
	"bomond-tenis/pkg/db/query"
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
	fmt.Println("book court")

	return nil
}

func (h *bookCourtHandler) Context() interface{} {
	return (*query.BookCourtQuery)(nil)
}

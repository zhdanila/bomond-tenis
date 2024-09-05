package courts

import (
	"bomond-tenis/pkg/db/query"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type getCourtsHandler struct {
	pool *sqlx.DB
}

func NewGetCourtsHandler(pool *sqlx.DB) *getCourtsHandler {
	return &getCourtsHandler{pool: pool}
}

func (h *getCourtsHandler) Exec(ctx context.Context, args *query.GetCourtsQuery) (err error) {
	fmt.Println("get courts")

	return nil
}

func (h *getCourtsHandler) Context() interface{} {
	return (*query.GetCourtsQuery)(nil)
}

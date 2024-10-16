package courts

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
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
	sql := fmt.Sprintf("SELECT * FROM %s", utils.CourtTable)

	rows, err := h.pool.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	var courts []query.Court

	for rows.Next() {
		var court query.Court
		if err := rows.Scan(&court.CourtId, &court.Name); err != nil {
			return err
		}
		courts = append(courts, court)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	args.Out.Courts = courts

	return nil
}

func (h *getCourtsHandler) Context() interface{} {
	return (*query.GetCourtsQuery)(nil)
}

package courts

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type GetCourtsHandler struct {
	pool *sqlx.DB
}

func NewGetCourtsHandler(pool *sqlx.DB) *GetCourtsHandler {
	return &GetCourtsHandler{pool: pool}
}

func (h *GetCourtsHandler) Exec(ctx context.Context, args *query.GetCourtsQuery) (err error) {
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

func (h *GetCourtsHandler) Context() interface{} {
	return (*query.GetCourtsQuery)(nil)
}

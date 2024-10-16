package courts

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
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
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", utils.CourtTable)

	rows := h.pool.QueryRow(sql, args.CourtId)

	var court query.Court
	if err = rows.Scan(&court.CourtId, &court.Name); err != nil {
		return err
	}

	args.Out.Court = court

	return nil
}

func (h *getBookedCourtHandler) Context() interface{} {
	return (*query.GetBookedCourtQuery)(nil)
}

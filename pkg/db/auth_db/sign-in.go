package auth_db

import (
	"bomond-tenis/pkg/db/query"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type signInHandler struct {
	pool *sqlx.DB
}

func NewSignInHandler(pool *sqlx.DB) *signInHandler {
	return &signInHandler{pool: pool}
}

func (h *signInHandler) Exec(ctx context.Context, args *query.SignInQuery) (err error) {
	//query := fmt.Sprintf(`INSERT INTO %s (uuid, account, type, parameters, data, status, created_at, send_at)
	//	VALUES (gen_random_uuid(), :account, :type, :parameters, :data, :status, :created_at, :send_at)
	//	RETURNING uuid`, NotificationTable)
	//
	//rows, err := h.pool.NamedQueryContext(ctx, query, params)
	//if err != nil {
	//	return err
	//}
	//
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var u string
	//
	//	if err = rows.Scan(&u); err != nil {
	//		return err
	//	}
	//
	//	uu, err := uuid.Parse(u)
	//	if err != nil {
	//		return err
	//	}
	//
	//	args.Uuid = uu
	//}

	fmt.Println("sign in")

	return nil
}

func (h *signInHandler) Context() interface{} {
	return (*query.SignInQuery)(nil)
}

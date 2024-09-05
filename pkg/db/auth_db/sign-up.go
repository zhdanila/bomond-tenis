package auth_db

import (
	db2 "bomond-tenis/pkg/db"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type signUpHandler struct {
	pool *sqlx.DB
}

func NewSignUpHandler(pool *sqlx.DB) *signUpHandler {
	return &signUpHandler{pool: pool}
}

func (h *signUpHandler) Exec(ctx context.Context, args *db2.SignUpQuery) (err error) {
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

	fmt.Println("logout")

	return nil
}

func (h *signUpHandler) Context() interface{} {
	return (*db2.SignUpQuery)(nil)
}

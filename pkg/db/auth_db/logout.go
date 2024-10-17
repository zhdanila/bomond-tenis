package auth_db

import (
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	redis2 "github.com/redis/go-redis/v9"
)

type LogoutHandler struct {
	pool  *sqlx.DB
	redis *redis2.Client
}

func NewLogoutHandler(pool *sqlx.DB, redis *redis2.Client) *LogoutHandler {
	return &LogoutHandler{
		pool:  pool,
		redis: redis,
	}
}

func (h *LogoutHandler) Exec(ctx context.Context, args *query.LogoutQuery) (err error) {
	err = h.redis.Set(ctx, args.Token, "blacklisted", utils.TokenTTL).Err()
	if err != nil {
		return fmt.Errorf("failed to add token to blacklist: %v", err)
	}

	return nil
}

func (h *LogoutHandler) Context() interface{} {
	return (*query.LogoutQuery)(nil)
}

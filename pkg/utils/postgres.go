package utils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func NewPostgresDB(host, user, dbname, password, sslmode string, port int) (*sqlx.DB, error) {
	connectUrl := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, password, sslmode)
	db, err := sqlx.Open("postgres", connectUrl)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Info().Msgf("connected to db")

	return db, nil
}

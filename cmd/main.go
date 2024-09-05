package main

import (
	"bomond-tenis/internal/api/restapi"
	config "bomond-tenis/internal/config"
	"bomond-tenis/pkg/api/http"
	"bomond-tenis/pkg/controller"
	"bomond-tenis/pkg/controller/configure"
	"bomond-tenis/pkg/utils"
	"context"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	redis2 "github.com/redis/go-redis/v9"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
)

const (
	appName = "Bomond-tenis"
)

func main() {
	var err error
	var cfg config.Environment
	var wg sync.WaitGroup

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	if err = env.Parse(&cfg); err != nil {
		panic(err)
	}

	ctx := context.Background()
	log.Logger = config.InitLog(appName, "debug")
	ctx = log.Logger.WithContext(ctx)

	var postgres *sqlx.DB
	postgres, err = utils.NewPostgresDB(
		cfg.PGDBHost,
		cfg.PGDBUser,
		cfg.PGDBName,
		cfg.PGDBPassword,
		cfg.PGDBSslMode,
		cfg.PGDBPort,
	)
	if err != nil {
		log.Panic().Err(err).Msgf("failed to connect to postgreSQL db")
	}
	defer func() {
		if err := postgres.Close(); err != nil {
			log.Error().Err(err).Msgf("failed to properly close postgreSQL db conn")
		}
	}()

	var redis *redis2.Client
	redis, err = utils.NewRedisDB(
		ctx,
		cfg.RedisHost,
		cfg.RedisPort,
		cfg.RedisPassword,
	)
	if err != nil {
		log.Panic().Err(err).Msgf("failed to connect to redis db")
	}
	defer func() {
		if err := redis.Close(); err != nil {
			log.Error().Err(err).Msgf("failed to properly close redis db conn")
		}
	}()

	ctrlImpl := controller.New()
	ctrl := wrapController(ctrlImpl)

	log.Info().Msg("configure controller")
	if err := configure.ControllerInit(
		ctrlImpl,
		ctrl,
		postgres,
		cfg,
	); err != nil {
		log.Panic().Err(err).Msg("failed to configure controller")
	}

	var serverHttp *restapi.Server
	if serverHttp, err = http.NewServer("", 8080, ctrl,
		func(ctx context.Context) error {
			return postgres.PingContext(ctx)
		},
	); err != nil {
		log.Panic().Err(err).Msg("unable to create http server")
	}

	errs := make(chan error, 4)
	go waitInterruptSignal(errs)

	wg.Add(1)
	go func() {
		defer wg.Done()
		errs <- serverHttp.Serve()
	}()

	log.Info().Msg("service started")
	err = <-errs
	log.Err(err).Msg("trying to shutdown gracefully")

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := serverHttp.Shutdown()
		log.Err(err).Msg("http server stopped")
	}()
}

func waitInterruptSignal(errs chan<- error) {
	c := make(chan os.Signal, 3)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	errs <- fmt.Errorf("%s", <-c)
	signal.Stop(c)
}

func wrapController(next controller.Controller) controller.Controller {
	return controller.RecoveryMiddleware(
		next,
	)
}

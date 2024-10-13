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
	"github.com/caarlos0/env/v11"
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
	appName = "Bomond-tennis"
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

	var pool *sqlx.DB
	pool, err = utils.NewPostgresDB(
		cfg.PGDB.PGDBHost,
		cfg.PGDB.User,
		cfg.PGDB.Name,
		cfg.PGDB.Password,
		cfg.PGDB.SslMode,
		cfg.PGDB.Port,
	)
	if err != nil {
		log.Panic().Err(err).Msgf("failed to connect to postgres db")
	}
	defer func() {
		if err := pool.Close(); err != nil {
			log.Error().Err(err).Msgf("failed to properly close db conn")
		}
	}()

	var redis *redis2.Client
	redis, err = utils.NewRedisDB(
		ctx,
		cfg.RedisDB.Host,
		cfg.RedisDB.Port,
		cfg.RedisDB.Password,
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
		pool,
		redis,
		cfg,
	); err != nil {
		log.Panic().Err(err).Msg("failed to configure controller")
	}

	var serverHttp *restapi.Server
	if serverHttp, err = http.NewServer("", cfg.HttpPort, ctrl, redis,
		func(ctx context.Context) error {
			return pool.PingContext(ctx)
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

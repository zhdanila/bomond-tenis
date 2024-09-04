package main

import (
	"bomond-tenis/internal/api/restapi"
	"bomond-tenis/internal/api/restapi/operations"
	config "bomond-tenis/internal/config"
	"bomond-tenis/internal/handlers"
	"context"
	"github.com/caarlos0/env"
	"github.com/go-openapi/loads"
	"github.com/joho/godotenv"

	"github.com/rs/zerolog/log"
)

const (
	appName = "MarketEvents"
)

func main() {
	var err error
	var cfg config.Environment
	//var wg sync.WaitGroup

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

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Panic().Err(err).Msg("failed to load swagger spec")
	}

	api := operations.NewBomondTenisAPI(swaggerSpec)
	handlers.ConfigureHandlers(api)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureFlags()
	server.ConfigureAPI()

	server.Host = cfg.HttpHost
	server.Port = cfg.HttpPort

	if err := server.Serve(); err != nil {
		log.Panic().Err(err).Msg("failed to run server")
	}

	//var pool *sqlx.DB
	//pool, err = utils.NewPostgresDB(cfg.PGDBHost,
	//	cfg.PGDBUser,
	//	cfg.PGDBName,
	//	cfg.PGDBPassword,
	//	cfg.PGDBSslMode,
	//	cfg.PGDBPort,
	//)
	//if err != nil {
	//	log.Panic().Err(err).Msgf("failed to connect to db")
	//}
	//defer func() {
	//	if err := pool.Close(); err != nil {
	//		log.Error().Err(err).Msgf("failed to properly close db conn")
	//	}
	//}()
	//
	//ctrlImpl := controller.New()
	//ctrl := wrapController(ctrlImpl)
	//
	//log.Info().Msg("configure controller")
	//if err := configure.ControllerInit(
	//	ctrlImpl,
	//	ctrl,
	//	pool,
	//	cfg,
	//); err != nil {
	//	log.Panic().Err(err).Msg("failed to configure controller")
	//}

	//var serverHttp *restapi.Server
	//if serverHttp, err = restapi.NewServer("", 8080, ctrl,
	//	func(ctx context.Context) error {
	//		return pool.PingContext(ctx)
	//	},
	//); err != nil {
	//	log.Panic().Err(err).Msg("unable to create http server")
	//}
	//
	//errs := make(chan error, 4)
	//go waitInterruptSignal(errs)
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	errs <- serverHttp.Serve()
	//}()
	//
	//log.Info().Msg("service started")
	//err = <-errs
	//log.Err(err).Msg("trying to shutdown gracefully")
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	err := serverHttp.Shutdown()
	//	log.Err(err).Msg("http server stopped")
	//}()
}

//func waitInterruptSignal(errs chan<- error) {
//	c := make(chan os.Signal, 3)
//	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
//	errs <- fmt.Errorf("%s", <-c)
//	signal.Stop(c)
//}
//
//func wrapController(next controller.Controller) controller.Controller {
//	return controller.RecoveryMiddleware(
//		next,
//	)
//}

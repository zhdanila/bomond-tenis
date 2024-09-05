package http

import (
	restapi2 "bomond-tenis/internal/api/restapi"
	operations2 "bomond-tenis/internal/api/restapi/operations"
	auth2 "bomond-tenis/pkg/api/http/handlers/auth"
	controller "bomond-tenis/pkg/controller"
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"runtime/debug"

	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/zerolog/log"
)

type Healthcheck func(ctx context.Context) error

func NewServer(host string, port int, ctrl controller.Controller, healthchecks ...Healthcheck) (*restapi2.Server, error) {
	spec, err := loads.Embedded(restapi2.SwaggerJSON, restapi2.FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	api := operations2.NewBomondTenisAPI(spec)

	//Users
	//api.UsersGetV1BomondVnUsersUserIDHandler = users.GetUserHandler(service.Users)
	//api.UsersPutV1BomondVnUsersUserIDHandler = users.PutUserHandler(service.Users)
	//api.UsersDeleteV1BomondVnUsersUserIDHandler = users.DeleteUserHandler(service.Users)

	//Authentication
	//api.AuthenticationPostV1BomondVnAuthSignInHandler = auth.SignInHandler(service.Authorization)
	//api.AuthenticationPostV1BomondVnAuthSignUpHandler = auth.SignUpHandler(service.Authorization)
	api.AuthenticationPostV1BomondVnAuthLogoutHandler = auth2.NewLogout(ctrl)

	//Courts
	//api.CourtsGetV1BomondVnCourtsHandler = courts.GetAllCourtsHandler(service.Courts)
	//api.CourtsGetV1BomondVnCourtIDBookHandler = courts.GetCourtByIdHandler(service.Courts)
	//api.CourtsPostV1BomondVnCourtIDBookHandler = courts.PostCourtHandler(service.Courts)
	//api.CourtsDeleteV1BomondVnCourtIDBookBookIDHandler = courts.DeleteCourtHandler(service.Courts)

	api.Logger = log.Logger.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.ServerShutdown = func() {}

	api.Middleware = func(builder middleware.Builder) http.Handler {
		middlewares := []func(h http.Handler) http.Handler{
			middlewarePprof,
			middlewareHealthz(healthchecks...),
			middlewareRecover,
			middlewareLogging,
		}

		return setupMiddleware(api.Context().RoutesHandler(builder), middlewares...)
	}

	server := restapi2.NewServer(api)
	server.Host = host
	server.Port = port

	return server, nil
}

func middlewarePprof(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/debug/pprof/cmdline") {
			pprof.Cmdline(rw, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/debug/pprof/profile") {
			pprof.Profile(rw, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/debug/pprof/symbol") {
			pprof.Symbol(rw, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/debug/pprof/trace") {
			pprof.Trace(rw, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/debug/pprof/") {
			pprof.Index(rw, r)
			return
		}

		h.ServeHTTP(rw, r)
	})
}

func middlewareHealthz(healthchecks ...Healthcheck) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/helthz" {
				ctx := r.Context()
				logger := log.Ctx(ctx)

				var err error
				for _, healthz := range healthchecks {
					if err = healthz(ctx); err != nil {
						logger.Err(err).Msg("helthz failed")
						break
					}
				}
				if err != nil {
					w.WriteHeader(500)
					//nolint:errcheck // there is no reason to ensure that error message gets fully dumped
					_, _ = w.Write([]byte(err.Error()))
				} else {
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte("ok"))
				}
				return
			}

			// define ping endpoint wich can fail only if service hangs completely
			if r.URL.Path == "/readyz" {
				w.WriteHeader(http.StatusOK)
				//nolint:errcheck // there is no reason to ensure that error message gets fully dumped
				_, _ = w.Write([]byte("ok"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func middlewareRecover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				logger := log.Ctx(r.Context())

				var err error
				switch errTyped := e.(type) {
				case error:
					err = errTyped
				default:
					err = fmt.Errorf("untyped panic: %v", e)
				}

				logger.Error().Err(err).Msgf("panic:\n%s", string(debug.Stack()))
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func middlewareLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var deviceId, userId string
		if idHeader := r.Header.Get("X-Request-Id"); len(idHeader) > 0 {
			ids := strings.Split(idHeader, "_")
			if len(ids) >= 2 {
				deviceId = ids[0]
				userId = ids[1]
			}
		}

		logger := log.Logger.With().Str("METHOD", r.Method).Str("PATH", r.URL.Path).Str("deviceId", deviceId).Str("userId", userId).Logger()
		h.ServeHTTP(w, r.WithContext(logger.WithContext(r.Context())))
	})
}

func setupMiddleware(init http.Handler, middlewares ...func(handler http.Handler) http.Handler) http.Handler {
	handler := init

	for _, h := range middlewares {
		handler = h(handler)
	}

	return handler
}
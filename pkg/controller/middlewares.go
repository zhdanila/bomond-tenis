package controller

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/cors"
	"net/http"
	"runtime/debug"
)

// controllerFunc provides a way to wrap singel function into controller.Controller type
type controllerFunc func(ctx context.Context, args interface{}) error

// Exec implements Controller.Exec
func (cf controllerFunc) Exec(ctx context.Context, args interface{}) error {
	return cf(ctx, args)
}

// RecoveryMiddleware is a controller that just call next one, wrapping with panic recover logic
// (e.g. wraps panic into error and return it)
func RecoveryMiddleware(next Controller) Controller {
	return controllerFunc(func(ctx context.Context, args interface{}) (e error) {
		defer func() {
			if r := recover(); r != nil {
				switch err := r.(type) {
				case error:
					e = fmt.Errorf("panic recovered: %w, %s", err, string(debug.Stack()))
				default:
					e = fmt.Errorf("untyped panic recovered: %v, %s", r, string(debug.Stack()))
				}
			}
		}()

		return next.Exec(ctx, args)
	})
}

// SetupGlobalMiddleware The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func SetupGlobalMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		Debug:            false,
		AllowedHeaders:   []string{"*"}, // TODO: config accordingly
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{},
		AllowCredentials: true,
		MaxAge:           1000,
	})
	h := corsHandler.Handler(handler)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var _, err = r.Cookie("guest_session")
		// cookie not set
		if err != nil {
			cookie := &http.Cookie{
				Name:     "guest_session",
				Value:    uuid.New().String(),
				SameSite: http.SameSiteDefaultMode,
				MaxAge:   260000}
			http.SetCookie(w, cookie)
		}
		h.ServeHTTP(w, r)
	})
}

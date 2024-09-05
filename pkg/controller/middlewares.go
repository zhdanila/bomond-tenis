package controller

import (
	"context"
	"fmt"
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

// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"bomond-tenis/internal/api/restapi/operations"
	"bomond-tenis/internal/api/restapi/operations/authentication"
	"bomond-tenis/internal/api/restapi/operations/courts"
	"bomond-tenis/internal/api/restapi/operations/users"
)

//go:generate swagger generate server --target ../../api --name BomondTenis --spec ../../../swagger/swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.BomondTenisAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BomondTenisAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CourtsDeleteV1BomondVnCourtIDBookBookIDHandler == nil {
		api.CourtsDeleteV1BomondVnCourtIDBookBookIDHandler = courts.DeleteV1BomondVnCourtIDBookBookIDHandlerFunc(func(params courts.DeleteV1BomondVnCourtIDBookBookIDParams) middleware.Responder {
			return middleware.NotImplemented("operation courts.DeleteV1BomondCourtIDBookBookID has not yet been implemented")
		})
	}
	if api.UsersDeleteV1BomondVnUsersUserIDHandler == nil {
		api.UsersDeleteV1BomondVnUsersUserIDHandler = users.DeleteV1BomondVnUsersUserIDHandlerFunc(func(params users.DeleteV1BomondVnUsersUserIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.DeleteV1BomondVnUsersUserID has not yet been implemented")
		})
	}
	if api.CourtsGetV1BomondVnCourtsHandler == nil {
		api.CourtsGetV1BomondVnCourtsHandler = courts.GetV1BomondVnCourtsHandlerFunc(func(params courts.GetV1BomondVnCourtsParams) middleware.Responder {
			return middleware.NotImplemented("operation courts.GetV1BomondCourts has not yet been implemented")
		})
	}
	if api.CourtsGetV1BomondVnCourtIDBookHandler == nil {
		api.CourtsGetV1BomondVnCourtIDBookHandler = courts.GetV1BomondVnCourtIDBookHandlerFunc(func(params courts.GetV1BomondVnCourtIDBookParams) middleware.Responder {
			return middleware.NotImplemented("operation courts.GetV1BomondVnCourtIDBook has not yet been implemented")
		})
	}
	if api.UsersGetV1BomondVnUsersUserIDHandler == nil {
		api.UsersGetV1BomondVnUsersUserIDHandler = users.GetV1BomondVnUsersUserIDHandlerFunc(func(params users.GetV1BomondVnUsersUserIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetV1BomondVnUsersUserID has not yet been implemented")
		})
	}
	if api.AuthenticationPostV1BomondVnAuthLogoutHandler == nil {
		api.AuthenticationPostV1BomondVnAuthLogoutHandler = authentication.PostV1BomondVnAuthLogoutHandlerFunc(func(params authentication.PostV1BomondVnAuthLogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.PostV1BomondVnAuthLogout has not yet been implemented")
		})
	}
	if api.AuthenticationPostV1BomondVnAuthSignInHandler == nil {
		api.AuthenticationPostV1BomondVnAuthSignInHandler = authentication.PostV1BomondVnAuthSignInHandlerFunc(func(params authentication.PostV1BomondVnAuthSignInParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.PostV1BomondVnAuthSignIn has not yet been implemented")
		})
	}
	if api.AuthenticationPostV1BomondVnAuthSignUpHandler == nil {
		api.AuthenticationPostV1BomondVnAuthSignUpHandler = authentication.PostV1BomondVnAuthSignUpHandlerFunc(func(params authentication.PostV1BomondVnAuthSignUpParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.PostV1BomondVnAuthSignUp has not yet been implemented")
		})
	}
	if api.CourtsPostV1BomondVnCourtIDBookHandler == nil {
		api.CourtsPostV1BomondVnCourtIDBookHandler = courts.PostV1BomondVnCourtIDBookHandlerFunc(func(params courts.PostV1BomondVnCourtIDBookParams) middleware.Responder {
			return middleware.NotImplemented("operation courts.PostV1BomondVnCourtIDBook has not yet been implemented")
		})
	}
	if api.UsersPutV1BomondVnUsersUserIDHandler == nil {
		api.UsersPutV1BomondVnUsersUserIDHandler = users.PutV1BomondVnUsersUserIDHandlerFunc(func(params users.PutV1BomondVnUsersUserIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PutV1BomondVnUsersUserID has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

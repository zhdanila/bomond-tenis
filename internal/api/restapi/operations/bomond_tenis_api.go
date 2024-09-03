// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"bomond-tenis/internal/api/restapi/operations/authentication"
	"bomond-tenis/internal/api/restapi/operations/courts"
	"bomond-tenis/internal/api/restapi/operations/users"
)

// NewBomondTenisAPI creates a new BomondTenis instance
func NewBomondTenisAPI(spec *loads.Document) *BomondTenisAPI {
	return &BomondTenisAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CourtsDeleteV1BomondCourtIDBookBookIDHandler: courts.DeleteV1BomondCourtIDBookBookIDHandlerFunc(func(params courts.DeleteV1BomondCourtIDBookBookIDParams) middleware.Responder {
			return middleware.NotImplemented("operation courts.DeleteV1BomondCourtIDBookBookID has not yet been implemented")
		}),
		UsersDeleteV1BomondVnUsersUserIDHandler: users.DeleteV1BomondVnUsersUserIDHandlerFunc(func(params users.DeleteV1BomondVnUsersUserIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.DeleteV1BomondVnUsersUserID has not yet been implemented")
		}),
		CourtsGetV1BomondCourtsHandler: courts.GetV1BomondCourtsHandlerFunc(func(params courts.GetV1BomondCourtsParams) middleware.Responder {
			return middleware.NotImplemented("operation courts.GetV1BomondCourts has not yet been implemented")
		}),
		CourtsGetV1BomondVnCourtIDBookHandler: courts.GetV1BomondVnCourtIDBookHandlerFunc(func(params courts.GetV1BomondVnCourtIDBookParams) middleware.Responder {
			return middleware.NotImplemented("operation courts.GetV1BomondVnCourtIDBook has not yet been implemented")
		}),
		UsersGetV1BomondVnUsersUserIDHandler: users.GetV1BomondVnUsersUserIDHandlerFunc(func(params users.GetV1BomondVnUsersUserIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetV1BomondVnUsersUserID has not yet been implemented")
		}),
		AuthenticationPostV1BomondVnAuthLogoutHandler: authentication.PostV1BomondVnAuthLogoutHandlerFunc(func(params authentication.PostV1BomondVnAuthLogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.PostV1BomondVnAuthLogout has not yet been implemented")
		}),
		AuthenticationPostV1BomondVnAuthSignInHandler: authentication.PostV1BomondVnAuthSignInHandlerFunc(func(params authentication.PostV1BomondVnAuthSignInParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.PostV1BomondVnAuthSignIn has not yet been implemented")
		}),
		AuthenticationPostV1BomondVnAuthSignUpHandler: authentication.PostV1BomondVnAuthSignUpHandlerFunc(func(params authentication.PostV1BomondVnAuthSignUpParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.PostV1BomondVnAuthSignUp has not yet been implemented")
		}),
		CourtsPostV1BomondVnCourtIDBookHandler: courts.PostV1BomondVnCourtIDBookHandlerFunc(func(params courts.PostV1BomondVnCourtIDBookParams) middleware.Responder {
			return middleware.NotImplemented("operation courts.PostV1BomondVnCourtIDBook has not yet been implemented")
		}),
		UsersPutV1BomondVnUsersUserIDHandler: users.PutV1BomondVnUsersUserIDHandlerFunc(func(params users.PutV1BomondVnUsersUserIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PutV1BomondVnUsersUserID has not yet been implemented")
		}),
	}
}

/*BomondTenisAPI the bomond tenis API */
type BomondTenisAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// CourtsDeleteV1BomondCourtIDBookBookIDHandler sets the operation handler for the delete v1 bomond court ID book book ID operation
	CourtsDeleteV1BomondCourtIDBookBookIDHandler courts.DeleteV1BomondCourtIDBookBookIDHandler
	// UsersDeleteV1BomondVnUsersUserIDHandler sets the operation handler for the delete v1 bomond vn users user ID operation
	UsersDeleteV1BomondVnUsersUserIDHandler users.DeleteV1BomondVnUsersUserIDHandler
	// CourtsGetV1BomondCourtsHandler sets the operation handler for the get v1 bomond courts operation
	CourtsGetV1BomondCourtsHandler courts.GetV1BomondCourtsHandler
	// CourtsGetV1BomondVnCourtIDBookHandler sets the operation handler for the get v1 bomond vn court ID book operation
	CourtsGetV1BomondVnCourtIDBookHandler courts.GetV1BomondVnCourtIDBookHandler
	// UsersGetV1BomondVnUsersUserIDHandler sets the operation handler for the get v1 bomond vn users user ID operation
	UsersGetV1BomondVnUsersUserIDHandler users.GetV1BomondVnUsersUserIDHandler
	// AuthenticationPostV1BomondVnAuthLogoutHandler sets the operation handler for the post v1 bomond vn auth logout operation
	AuthenticationPostV1BomondVnAuthLogoutHandler authentication.PostV1BomondVnAuthLogoutHandler
	// AuthenticationPostV1BomondVnAuthSignInHandler sets the operation handler for the post v1 bomond vn auth sign in operation
	AuthenticationPostV1BomondVnAuthSignInHandler authentication.PostV1BomondVnAuthSignInHandler
	// AuthenticationPostV1BomondVnAuthSignUpHandler sets the operation handler for the post v1 bomond vn auth sign up operation
	AuthenticationPostV1BomondVnAuthSignUpHandler authentication.PostV1BomondVnAuthSignUpHandler
	// CourtsPostV1BomondVnCourtIDBookHandler sets the operation handler for the post v1 bomond vn court ID book operation
	CourtsPostV1BomondVnCourtIDBookHandler courts.PostV1BomondVnCourtIDBookHandler
	// UsersPutV1BomondVnUsersUserIDHandler sets the operation handler for the put v1 bomond vn users user ID operation
	UsersPutV1BomondVnUsersUserIDHandler users.PutV1BomondVnUsersUserIDHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *BomondTenisAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *BomondTenisAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *BomondTenisAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BomondTenisAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BomondTenisAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BomondTenisAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BomondTenisAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BomondTenisAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BomondTenisAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BomondTenisAPI
func (o *BomondTenisAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CourtsDeleteV1BomondCourtIDBookBookIDHandler == nil {
		unregistered = append(unregistered, "courts.DeleteV1BomondCourtIDBookBookIDHandler")
	}
	if o.UsersDeleteV1BomondVnUsersUserIDHandler == nil {
		unregistered = append(unregistered, "users.DeleteV1BomondVnUsersUserIDHandler")
	}
	if o.CourtsGetV1BomondCourtsHandler == nil {
		unregistered = append(unregistered, "courts.GetV1BomondCourtsHandler")
	}
	if o.CourtsGetV1BomondVnCourtIDBookHandler == nil {
		unregistered = append(unregistered, "courts.GetV1BomondVnCourtIDBookHandler")
	}
	if o.UsersGetV1BomondVnUsersUserIDHandler == nil {
		unregistered = append(unregistered, "users.GetV1BomondVnUsersUserIDHandler")
	}
	if o.AuthenticationPostV1BomondVnAuthLogoutHandler == nil {
		unregistered = append(unregistered, "authentication.PostV1BomondVnAuthLogoutHandler")
	}
	if o.AuthenticationPostV1BomondVnAuthSignInHandler == nil {
		unregistered = append(unregistered, "authentication.PostV1BomondVnAuthSignInHandler")
	}
	if o.AuthenticationPostV1BomondVnAuthSignUpHandler == nil {
		unregistered = append(unregistered, "authentication.PostV1BomondVnAuthSignUpHandler")
	}
	if o.CourtsPostV1BomondVnCourtIDBookHandler == nil {
		unregistered = append(unregistered, "courts.PostV1BomondVnCourtIDBookHandler")
	}
	if o.UsersPutV1BomondVnUsersUserIDHandler == nil {
		unregistered = append(unregistered, "users.PutV1BomondVnUsersUserIDHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BomondTenisAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BomondTenisAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *BomondTenisAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *BomondTenisAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *BomondTenisAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *BomondTenisAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the bomond tenis API
func (o *BomondTenisAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BomondTenisAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v1/bomond/{court_id}/book/{book_id}"] = courts.NewDeleteV1BomondCourtIDBookBookID(o.context, o.CourtsDeleteV1BomondCourtIDBookBookIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v1/bomond.vn/users/{user_id}"] = users.NewDeleteV1BomondVnUsersUserID(o.context, o.UsersDeleteV1BomondVnUsersUserIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/bomond/courts"] = courts.NewGetV1BomondCourts(o.context, o.CourtsGetV1BomondCourtsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/bomond.vn/{court_id}/book"] = courts.NewGetV1BomondVnCourtIDBook(o.context, o.CourtsGetV1BomondVnCourtIDBookHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/bomond.vn/users/{user_id}"] = users.NewGetV1BomondVnUsersUserID(o.context, o.UsersGetV1BomondVnUsersUserIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/bomond.vn/auth/logout"] = authentication.NewPostV1BomondVnAuthLogout(o.context, o.AuthenticationPostV1BomondVnAuthLogoutHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/bomond.vn/auth/sign-in"] = authentication.NewPostV1BomondVnAuthSignIn(o.context, o.AuthenticationPostV1BomondVnAuthSignInHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/bomond.vn/auth/sign-up"] = authentication.NewPostV1BomondVnAuthSignUp(o.context, o.AuthenticationPostV1BomondVnAuthSignUpHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/bomond.vn/{court_id}/book"] = courts.NewPostV1BomondVnCourtIDBook(o.context, o.CourtsPostV1BomondVnCourtIDBookHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/bomond.vn/users/{user_id}"] = users.NewPutV1BomondVnUsersUserID(o.context, o.UsersPutV1BomondVnUsersUserIDHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BomondTenisAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *BomondTenisAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *BomondTenisAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *BomondTenisAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *BomondTenisAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}

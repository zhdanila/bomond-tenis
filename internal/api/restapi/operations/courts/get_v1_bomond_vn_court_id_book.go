// Code generated by go-swagger; DO NOT EDIT.

package courts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetV1BomondVnCourtIDBookHandlerFunc turns a function with the right signature into a get v1 bomond vn court ID book handler
type GetV1BomondVnCourtIDBookHandlerFunc func(GetV1BomondVnCourtIDBookParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetV1BomondVnCourtIDBookHandlerFunc) Handle(params GetV1BomondVnCourtIDBookParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetV1BomondVnCourtIDBookHandler interface for that can handle valid get v1 bomond vn court ID book params
type GetV1BomondVnCourtIDBookHandler interface {
	Handle(GetV1BomondVnCourtIDBookParams, interface{}) middleware.Responder
}

// NewGetV1BomondVnCourtIDBook creates a new http.Handler for the get v1 bomond vn court ID book operation
func NewGetV1BomondVnCourtIDBook(ctx *middleware.Context, handler GetV1BomondVnCourtIDBookHandler) *GetV1BomondVnCourtIDBook {
	return &GetV1BomondVnCourtIDBook{Context: ctx, Handler: handler}
}

/*
	GetV1BomondVnCourtIDBook swagger:route GET /v1/bomond.vn/{court_id}/book Courts getV1BomondVnCourtIdBook

GetV1BomondVnCourtIDBook get v1 bomond vn court ID book API
*/
type GetV1BomondVnCourtIDBook struct {
	Context *middleware.Context
	Handler GetV1BomondVnCourtIDBookHandler
}

func (o *GetV1BomondVnCourtIDBook) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetV1BomondVnCourtIDBookParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

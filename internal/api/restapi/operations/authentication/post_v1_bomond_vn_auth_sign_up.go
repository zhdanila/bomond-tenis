// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostV1BomondVnAuthSignUpHandlerFunc turns a function with the right signature into a post v1 bomond vn auth sign up handler
type PostV1BomondVnAuthSignUpHandlerFunc func(PostV1BomondVnAuthSignUpParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostV1BomondVnAuthSignUpHandlerFunc) Handle(params PostV1BomondVnAuthSignUpParams) middleware.Responder {
	return fn(params)
}

// PostV1BomondVnAuthSignUpHandler interface for that can handle valid post v1 bomond vn auth sign up params
type PostV1BomondVnAuthSignUpHandler interface {
	Handle(PostV1BomondVnAuthSignUpParams) middleware.Responder
}

// NewPostV1BomondVnAuthSignUp creates a new http.Handler for the post v1 bomond vn auth sign up operation
func NewPostV1BomondVnAuthSignUp(ctx *middleware.Context, handler PostV1BomondVnAuthSignUpHandler) *PostV1BomondVnAuthSignUp {
	return &PostV1BomondVnAuthSignUp{Context: ctx, Handler: handler}
}

/*
	PostV1BomondVnAuthSignUp swagger:route POST /v1/bomond.vn/auth/sign-up Authentication postV1BomondVnAuthSignUp

PostV1BomondVnAuthSignUp post v1 bomond vn auth sign up API
*/
type PostV1BomondVnAuthSignUp struct {
	Context *middleware.Context
	Handler PostV1BomondVnAuthSignUpHandler
}

func (o *PostV1BomondVnAuthSignUp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostV1BomondVnAuthSignUpParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostV1BomondVnAuthSignUpBody post v1 bomond vn auth sign up body
//
// swagger:model PostV1BomondVnAuthSignUpBody
type PostV1BomondVnAuthSignUpBody struct {

	// email
	Email string `json:"email,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this post v1 bomond vn auth sign up body
func (o *PostV1BomondVnAuthSignUpBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post v1 bomond vn auth sign up body based on context it is used
func (o *PostV1BomondVnAuthSignUpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostV1BomondVnAuthSignUpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostV1BomondVnAuthSignUpBody) UnmarshalBinary(b []byte) error {
	var res PostV1BomondVnAuthSignUpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

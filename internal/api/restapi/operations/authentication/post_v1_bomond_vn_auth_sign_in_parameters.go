// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"
)

// NewPostV1BomondVnAuthSignInParams creates a new PostV1BomondVnAuthSignInParams object
//
// There are no default values defined in the spec.
func NewPostV1BomondVnAuthSignInParams() PostV1BomondVnAuthSignInParams {

	return PostV1BomondVnAuthSignInParams{}
}

// PostV1BomondVnAuthSignInParams contains all the bound params for the post v1 bomond vn auth sign in operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostV1BomondVnAuthSignIn
type PostV1BomondVnAuthSignInParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	UserSignin PostV1BomondVnAuthSignInBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostV1BomondVnAuthSignInParams() beforehand.
func (o *PostV1BomondVnAuthSignInParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PostV1BomondVnAuthSignInBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("userSignin", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.UserSignin = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

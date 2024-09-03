// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"
)

// NewPostV1BomondVnAuthSignUpParams creates a new PostV1BomondVnAuthSignUpParams object
//
// There are no default values defined in the spec.
func NewPostV1BomondVnAuthSignUpParams() PostV1BomondVnAuthSignUpParams {

	return PostV1BomondVnAuthSignUpParams{}
}

// PostV1BomondVnAuthSignUpParams contains all the bound params for the post v1 bomond vn auth sign up operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostV1BomondVnAuthSignUp
type PostV1BomondVnAuthSignUpParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	UserSignup PostV1BomondVnAuthSignUpBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostV1BomondVnAuthSignUpParams() beforehand.
func (o *PostV1BomondVnAuthSignUpParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PostV1BomondVnAuthSignUpBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("userSignup", "body", ""))
			} else {
				res = append(res, errors.NewParseError("userSignup", "body", "", err))
			}
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
				o.UserSignup = body
			}
		}
	} else {
		res = append(res, errors.Required("userSignup", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

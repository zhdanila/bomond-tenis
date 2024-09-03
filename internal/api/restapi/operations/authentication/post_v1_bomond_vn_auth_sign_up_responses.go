// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bomond-tenis/internal/api/models"
)

// PostV1BomondVnAuthSignUpOKCode is the HTTP code returned for type PostV1BomondVnAuthSignUpOK
const PostV1BomondVnAuthSignUpOKCode int = 200

/*
PostV1BomondVnAuthSignUpOK User successfully signed up

swagger:response postV1BomondVnAuthSignUpOK
*/
type PostV1BomondVnAuthSignUpOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessResponse `json:"body,omitempty"`
}

// NewPostV1BomondVnAuthSignUpOK creates PostV1BomondVnAuthSignUpOK with default headers values
func NewPostV1BomondVnAuthSignUpOK() *PostV1BomondVnAuthSignUpOK {

	return &PostV1BomondVnAuthSignUpOK{}
}

// WithPayload adds the payload to the post v1 bomond vn auth sign up o k response
func (o *PostV1BomondVnAuthSignUpOK) WithPayload(payload *models.SuccessResponse) *PostV1BomondVnAuthSignUpOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 bomond vn auth sign up o k response
func (o *PostV1BomondVnAuthSignUpOK) SetPayload(payload *models.SuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1BomondVnAuthSignUpOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostV1BomondVnAuthSignUpBadRequestCode is the HTTP code returned for type PostV1BomondVnAuthSignUpBadRequest
const PostV1BomondVnAuthSignUpBadRequestCode int = 400

/*
PostV1BomondVnAuthSignUpBadRequest Bad request

swagger:response postV1BomondVnAuthSignUpBadRequest
*/
type PostV1BomondVnAuthSignUpBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewPostV1BomondVnAuthSignUpBadRequest creates PostV1BomondVnAuthSignUpBadRequest with default headers values
func NewPostV1BomondVnAuthSignUpBadRequest() *PostV1BomondVnAuthSignUpBadRequest {

	return &PostV1BomondVnAuthSignUpBadRequest{}
}

// WithPayload adds the payload to the post v1 bomond vn auth sign up bad request response
func (o *PostV1BomondVnAuthSignUpBadRequest) WithPayload(payload *models.ErrorResult) *PostV1BomondVnAuthSignUpBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 bomond vn auth sign up bad request response
func (o *PostV1BomondVnAuthSignUpBadRequest) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1BomondVnAuthSignUpBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

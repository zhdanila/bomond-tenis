// Code generated by go-swagger; DO NOT EDIT.

package courts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bomond-tenis/internal/api/models"
)

// GetV1BomondVnCourtIDBookOKCode is the HTTP code returned for type GetV1BomondVnCourtIDBookOK
const GetV1BomondVnCourtIDBookOKCode int = 200

/*
GetV1BomondVnCourtIDBookOK Booking details retrieved

swagger:response getV1BomondVnCourtIdBookOK
*/
type GetV1BomondVnCourtIDBookOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessResponse `json:"body,omitempty"`
}

// NewGetV1BomondVnCourtIDBookOK creates GetV1BomondVnCourtIDBookOK with default headers values
func NewGetV1BomondVnCourtIDBookOK() *GetV1BomondVnCourtIDBookOK {

	return &GetV1BomondVnCourtIDBookOK{}
}

// WithPayload adds the payload to the get v1 bomond vn court Id book o k response
func (o *GetV1BomondVnCourtIDBookOK) WithPayload(payload *models.SuccessResponse) *GetV1BomondVnCourtIDBookOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 bomond vn court Id book o k response
func (o *GetV1BomondVnCourtIDBookOK) SetPayload(payload *models.SuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1BomondVnCourtIDBookOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetV1BomondVnCourtIDBookBadRequestCode is the HTTP code returned for type GetV1BomondVnCourtIDBookBadRequest
const GetV1BomondVnCourtIDBookBadRequestCode int = 400

/*
GetV1BomondVnCourtIDBookBadRequest Bad request

swagger:response getV1BomondVnCourtIdBookBadRequest
*/
type GetV1BomondVnCourtIDBookBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewGetV1BomondVnCourtIDBookBadRequest creates GetV1BomondVnCourtIDBookBadRequest with default headers values
func NewGetV1BomondVnCourtIDBookBadRequest() *GetV1BomondVnCourtIDBookBadRequest {

	return &GetV1BomondVnCourtIDBookBadRequest{}
}

// WithPayload adds the payload to the get v1 bomond vn court Id book bad request response
func (o *GetV1BomondVnCourtIDBookBadRequest) WithPayload(payload *models.ErrorResult) *GetV1BomondVnCourtIDBookBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 bomond vn court Id book bad request response
func (o *GetV1BomondVnCourtIDBookBadRequest) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1BomondVnCourtIDBookBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

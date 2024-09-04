// Code generated by go-swagger; DO NOT EDIT.

package courts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bomond-tenis/internal/api/models"
)

// GetV1BomondVnCourtsOKCode is the HTTP code returned for type GetV1BomondVnCourtsOK
const GetV1BomondVnCourtsOKCode int = 200

/*
GetV1BomondVnCourtsOK Courts list retrieved

swagger:response getV1BomondVnCourtsOK
*/
type GetV1BomondVnCourtsOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessResponse `json:"body,omitempty"`
}

// NewGetV1BomondVnCourtsOK creates GetV1BomondVnCourtsOK with default headers values
func NewGetV1BomondVnCourtsOK() *GetV1BomondVnCourtsOK {

	return &GetV1BomondVnCourtsOK{}
}

// WithPayload adds the payload to the get v1 bomond vn courts o k response
func (o *GetV1BomondVnCourtsOK) WithPayload(payload *models.SuccessResponse) *GetV1BomondVnCourtsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 bomond vn courts o k response
func (o *GetV1BomondVnCourtsOK) SetPayload(payload *models.SuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1BomondVnCourtsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetV1BomondVnCourtsBadRequestCode is the HTTP code returned for type GetV1BomondVnCourtsBadRequest
const GetV1BomondVnCourtsBadRequestCode int = 400

/*
GetV1BomondVnCourtsBadRequest Bad request

swagger:response getV1BomondVnCourtsBadRequest
*/
type GetV1BomondVnCourtsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewGetV1BomondVnCourtsBadRequest creates GetV1BomondVnCourtsBadRequest with default headers values
func NewGetV1BomondVnCourtsBadRequest() *GetV1BomondVnCourtsBadRequest {

	return &GetV1BomondVnCourtsBadRequest{}
}

// WithPayload adds the payload to the get v1 bomond vn courts bad request response
func (o *GetV1BomondVnCourtsBadRequest) WithPayload(payload *models.ErrorResult) *GetV1BomondVnCourtsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 bomond vn courts bad request response
func (o *GetV1BomondVnCourtsBadRequest) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1BomondVnCourtsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
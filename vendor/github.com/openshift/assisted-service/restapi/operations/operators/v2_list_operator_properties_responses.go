// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2ListOperatorPropertiesOKCode is the HTTP code returned for type V2ListOperatorPropertiesOK
const V2ListOperatorPropertiesOKCode int = 200

/*
V2ListOperatorPropertiesOK Success.

swagger:response v2ListOperatorPropertiesOK
*/
type V2ListOperatorPropertiesOK struct {

	/*
	  In: Body
	*/
	Payload models.OperatorProperties `json:"body,omitempty"`
}

// NewV2ListOperatorPropertiesOK creates V2ListOperatorPropertiesOK with default headers values
func NewV2ListOperatorPropertiesOK() *V2ListOperatorPropertiesOK {

	return &V2ListOperatorPropertiesOK{}
}

// WithPayload adds the payload to the v2 list operator properties o k response
func (o *V2ListOperatorPropertiesOK) WithPayload(payload models.OperatorProperties) *V2ListOperatorPropertiesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list operator properties o k response
func (o *V2ListOperatorPropertiesOK) SetPayload(payload models.OperatorProperties) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListOperatorPropertiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.OperatorProperties{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// V2ListOperatorPropertiesUnauthorizedCode is the HTTP code returned for type V2ListOperatorPropertiesUnauthorized
const V2ListOperatorPropertiesUnauthorizedCode int = 401

/*
V2ListOperatorPropertiesUnauthorized Unauthorized.

swagger:response v2ListOperatorPropertiesUnauthorized
*/
type V2ListOperatorPropertiesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2ListOperatorPropertiesUnauthorized creates V2ListOperatorPropertiesUnauthorized with default headers values
func NewV2ListOperatorPropertiesUnauthorized() *V2ListOperatorPropertiesUnauthorized {

	return &V2ListOperatorPropertiesUnauthorized{}
}

// WithPayload adds the payload to the v2 list operator properties unauthorized response
func (o *V2ListOperatorPropertiesUnauthorized) WithPayload(payload *models.InfraError) *V2ListOperatorPropertiesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list operator properties unauthorized response
func (o *V2ListOperatorPropertiesUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListOperatorPropertiesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ListOperatorPropertiesForbiddenCode is the HTTP code returned for type V2ListOperatorPropertiesForbidden
const V2ListOperatorPropertiesForbiddenCode int = 403

/*
V2ListOperatorPropertiesForbidden Forbidden.

swagger:response v2ListOperatorPropertiesForbidden
*/
type V2ListOperatorPropertiesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2ListOperatorPropertiesForbidden creates V2ListOperatorPropertiesForbidden with default headers values
func NewV2ListOperatorPropertiesForbidden() *V2ListOperatorPropertiesForbidden {

	return &V2ListOperatorPropertiesForbidden{}
}

// WithPayload adds the payload to the v2 list operator properties forbidden response
func (o *V2ListOperatorPropertiesForbidden) WithPayload(payload *models.InfraError) *V2ListOperatorPropertiesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list operator properties forbidden response
func (o *V2ListOperatorPropertiesForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListOperatorPropertiesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ListOperatorPropertiesNotFoundCode is the HTTP code returned for type V2ListOperatorPropertiesNotFound
const V2ListOperatorPropertiesNotFoundCode int = 404

/*
V2ListOperatorPropertiesNotFound Error.

swagger:response v2ListOperatorPropertiesNotFound
*/
type V2ListOperatorPropertiesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ListOperatorPropertiesNotFound creates V2ListOperatorPropertiesNotFound with default headers values
func NewV2ListOperatorPropertiesNotFound() *V2ListOperatorPropertiesNotFound {

	return &V2ListOperatorPropertiesNotFound{}
}

// WithPayload adds the payload to the v2 list operator properties not found response
func (o *V2ListOperatorPropertiesNotFound) WithPayload(payload *models.Error) *V2ListOperatorPropertiesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list operator properties not found response
func (o *V2ListOperatorPropertiesNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListOperatorPropertiesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ListOperatorPropertiesInternalServerErrorCode is the HTTP code returned for type V2ListOperatorPropertiesInternalServerError
const V2ListOperatorPropertiesInternalServerErrorCode int = 500

/*
V2ListOperatorPropertiesInternalServerError Error.

swagger:response v2ListOperatorPropertiesInternalServerError
*/
type V2ListOperatorPropertiesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ListOperatorPropertiesInternalServerError creates V2ListOperatorPropertiesInternalServerError with default headers values
func NewV2ListOperatorPropertiesInternalServerError() *V2ListOperatorPropertiesInternalServerError {

	return &V2ListOperatorPropertiesInternalServerError{}
}

// WithPayload adds the payload to the v2 list operator properties internal server error response
func (o *V2ListOperatorPropertiesInternalServerError) WithPayload(payload *models.Error) *V2ListOperatorPropertiesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list operator properties internal server error response
func (o *V2ListOperatorPropertiesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListOperatorPropertiesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
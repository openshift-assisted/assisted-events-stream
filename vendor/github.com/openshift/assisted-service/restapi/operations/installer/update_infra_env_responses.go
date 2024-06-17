// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// UpdateInfraEnvCreatedCode is the HTTP code returned for type UpdateInfraEnvCreated
const UpdateInfraEnvCreatedCode int = 201

/*
UpdateInfraEnvCreated Success.

swagger:response updateInfraEnvCreated
*/
type UpdateInfraEnvCreated struct {

	/*
	  In: Body
	*/
	Payload *models.InfraEnv `json:"body,omitempty"`
}

// NewUpdateInfraEnvCreated creates UpdateInfraEnvCreated with default headers values
func NewUpdateInfraEnvCreated() *UpdateInfraEnvCreated {

	return &UpdateInfraEnvCreated{}
}

// WithPayload adds the payload to the update infra env created response
func (o *UpdateInfraEnvCreated) WithPayload(payload *models.InfraEnv) *UpdateInfraEnvCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update infra env created response
func (o *UpdateInfraEnvCreated) SetPayload(payload *models.InfraEnv) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInfraEnvCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateInfraEnvBadRequestCode is the HTTP code returned for type UpdateInfraEnvBadRequest
const UpdateInfraEnvBadRequestCode int = 400

/*
UpdateInfraEnvBadRequest Error.

swagger:response updateInfraEnvBadRequest
*/
type UpdateInfraEnvBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateInfraEnvBadRequest creates UpdateInfraEnvBadRequest with default headers values
func NewUpdateInfraEnvBadRequest() *UpdateInfraEnvBadRequest {

	return &UpdateInfraEnvBadRequest{}
}

// WithPayload adds the payload to the update infra env bad request response
func (o *UpdateInfraEnvBadRequest) WithPayload(payload *models.Error) *UpdateInfraEnvBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update infra env bad request response
func (o *UpdateInfraEnvBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInfraEnvBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateInfraEnvUnauthorizedCode is the HTTP code returned for type UpdateInfraEnvUnauthorized
const UpdateInfraEnvUnauthorizedCode int = 401

/*
UpdateInfraEnvUnauthorized Unauthorized.

swagger:response updateInfraEnvUnauthorized
*/
type UpdateInfraEnvUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUpdateInfraEnvUnauthorized creates UpdateInfraEnvUnauthorized with default headers values
func NewUpdateInfraEnvUnauthorized() *UpdateInfraEnvUnauthorized {

	return &UpdateInfraEnvUnauthorized{}
}

// WithPayload adds the payload to the update infra env unauthorized response
func (o *UpdateInfraEnvUnauthorized) WithPayload(payload *models.InfraError) *UpdateInfraEnvUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update infra env unauthorized response
func (o *UpdateInfraEnvUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInfraEnvUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateInfraEnvForbiddenCode is the HTTP code returned for type UpdateInfraEnvForbidden
const UpdateInfraEnvForbiddenCode int = 403

/*
UpdateInfraEnvForbidden Forbidden.

swagger:response updateInfraEnvForbidden
*/
type UpdateInfraEnvForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUpdateInfraEnvForbidden creates UpdateInfraEnvForbidden with default headers values
func NewUpdateInfraEnvForbidden() *UpdateInfraEnvForbidden {

	return &UpdateInfraEnvForbidden{}
}

// WithPayload adds the payload to the update infra env forbidden response
func (o *UpdateInfraEnvForbidden) WithPayload(payload *models.InfraError) *UpdateInfraEnvForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update infra env forbidden response
func (o *UpdateInfraEnvForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInfraEnvForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateInfraEnvNotFoundCode is the HTTP code returned for type UpdateInfraEnvNotFound
const UpdateInfraEnvNotFoundCode int = 404

/*
UpdateInfraEnvNotFound Error.

swagger:response updateInfraEnvNotFound
*/
type UpdateInfraEnvNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateInfraEnvNotFound creates UpdateInfraEnvNotFound with default headers values
func NewUpdateInfraEnvNotFound() *UpdateInfraEnvNotFound {

	return &UpdateInfraEnvNotFound{}
}

// WithPayload adds the payload to the update infra env not found response
func (o *UpdateInfraEnvNotFound) WithPayload(payload *models.Error) *UpdateInfraEnvNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update infra env not found response
func (o *UpdateInfraEnvNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInfraEnvNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateInfraEnvMethodNotAllowedCode is the HTTP code returned for type UpdateInfraEnvMethodNotAllowed
const UpdateInfraEnvMethodNotAllowedCode int = 405

/*
UpdateInfraEnvMethodNotAllowed Method Not Allowed.

swagger:response updateInfraEnvMethodNotAllowed
*/
type UpdateInfraEnvMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateInfraEnvMethodNotAllowed creates UpdateInfraEnvMethodNotAllowed with default headers values
func NewUpdateInfraEnvMethodNotAllowed() *UpdateInfraEnvMethodNotAllowed {

	return &UpdateInfraEnvMethodNotAllowed{}
}

// WithPayload adds the payload to the update infra env method not allowed response
func (o *UpdateInfraEnvMethodNotAllowed) WithPayload(payload *models.Error) *UpdateInfraEnvMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update infra env method not allowed response
func (o *UpdateInfraEnvMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInfraEnvMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateInfraEnvConflictCode is the HTTP code returned for type UpdateInfraEnvConflict
const UpdateInfraEnvConflictCode int = 409

/*
UpdateInfraEnvConflict Error.

swagger:response updateInfraEnvConflict
*/
type UpdateInfraEnvConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateInfraEnvConflict creates UpdateInfraEnvConflict with default headers values
func NewUpdateInfraEnvConflict() *UpdateInfraEnvConflict {

	return &UpdateInfraEnvConflict{}
}

// WithPayload adds the payload to the update infra env conflict response
func (o *UpdateInfraEnvConflict) WithPayload(payload *models.Error) *UpdateInfraEnvConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update infra env conflict response
func (o *UpdateInfraEnvConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInfraEnvConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateInfraEnvInternalServerErrorCode is the HTTP code returned for type UpdateInfraEnvInternalServerError
const UpdateInfraEnvInternalServerErrorCode int = 500

/*
UpdateInfraEnvInternalServerError Error.

swagger:response updateInfraEnvInternalServerError
*/
type UpdateInfraEnvInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateInfraEnvInternalServerError creates UpdateInfraEnvInternalServerError with default headers values
func NewUpdateInfraEnvInternalServerError() *UpdateInfraEnvInternalServerError {

	return &UpdateInfraEnvInternalServerError{}
}

// WithPayload adds the payload to the update infra env internal server error response
func (o *UpdateInfraEnvInternalServerError) WithPayload(payload *models.Error) *UpdateInfraEnvInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update infra env internal server error response
func (o *UpdateInfraEnvInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInfraEnvInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateInfraEnvNotImplementedCode is the HTTP code returned for type UpdateInfraEnvNotImplemented
const UpdateInfraEnvNotImplementedCode int = 501

/*
UpdateInfraEnvNotImplemented Not implemented.

swagger:response updateInfraEnvNotImplemented
*/
type UpdateInfraEnvNotImplemented struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateInfraEnvNotImplemented creates UpdateInfraEnvNotImplemented with default headers values
func NewUpdateInfraEnvNotImplemented() *UpdateInfraEnvNotImplemented {

	return &UpdateInfraEnvNotImplemented{}
}

// WithPayload adds the payload to the update infra env not implemented response
func (o *UpdateInfraEnvNotImplemented) WithPayload(payload *models.Error) *UpdateInfraEnvNotImplemented {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update infra env not implemented response
func (o *UpdateInfraEnvNotImplemented) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInfraEnvNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
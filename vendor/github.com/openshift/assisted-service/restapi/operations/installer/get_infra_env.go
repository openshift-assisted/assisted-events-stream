// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetInfraEnvHandlerFunc turns a function with the right signature into a get infra env handler
type GetInfraEnvHandlerFunc func(GetInfraEnvParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetInfraEnvHandlerFunc) Handle(params GetInfraEnvParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetInfraEnvHandler interface for that can handle valid get infra env params
type GetInfraEnvHandler interface {
	Handle(GetInfraEnvParams, interface{}) middleware.Responder
}

// NewGetInfraEnv creates a new http.Handler for the get infra env operation
func NewGetInfraEnv(ctx *middleware.Context, handler GetInfraEnvHandler) *GetInfraEnv {
	return &GetInfraEnv{Context: ctx, Handler: handler}
}

/*
	GetInfraEnv swagger:route GET /v2/infra-envs/{infra_env_id} installer getInfraEnv

Retrieves the details of the infra-env.
*/
type GetInfraEnv struct {
	Context *middleware.Context
	Handler GetInfraEnvHandler
}

func (o *GetInfraEnv) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetInfraEnvParams()
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

// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutPolicyHandlerFunc turns a function with the right signature into a put policy handler
type PutPolicyHandlerFunc func(PutPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutPolicyHandlerFunc) Handle(params PutPolicyParams) middleware.Responder {
	return fn(params)
}

// PutPolicyHandler interface for that can handle valid put policy params
type PutPolicyHandler interface {
	Handle(PutPolicyParams) middleware.Responder
}

// NewPutPolicy creates a new http.Handler for the put policy operation
func NewPutPolicy(ctx *middleware.Context, handler PutPolicyHandler) *PutPolicy {
	return &PutPolicy{Context: ctx, Handler: handler}
}

/*PutPolicy swagger:route PUT /policy policy putPolicy

Create or update a policy (sub)tree

*/
type PutPolicy struct {
	Context *middleware.Context
	Handler PutPolicyHandler
}

func (o *PutPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

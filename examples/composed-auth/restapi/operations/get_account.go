// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/France-ioi/go-swagger/examples/composed-auth/models"
)

// GetAccountHandlerFunc turns a function with the right signature into a get account handler
type GetAccountHandlerFunc func(GetAccountParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountHandlerFunc) Handle(params GetAccountParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetAccountHandler interface for that can handle valid get account params
type GetAccountHandler interface {
	Handle(GetAccountParams, *models.Principal) middleware.Responder
}

// NewGetAccount creates a new http.Handler for the get account operation
func NewGetAccount(ctx *middleware.Context, handler GetAccountHandler) *GetAccount {
	return &GetAccount{Context: ctx, Handler: handler}
}

/*
	GetAccount swagger:route GET /account getAccount

registered user account

Every registered user should be able to access this operation
*/
type GetAccount struct {
	Context *middleware.Context
	Handler GetAccountHandler
}

func (o *GetAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAccountParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

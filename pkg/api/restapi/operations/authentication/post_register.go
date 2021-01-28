// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostRegisterHandlerFunc turns a function with the right signature into a post register handler
type PostRegisterHandlerFunc func(PostRegisterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRegisterHandlerFunc) Handle(params PostRegisterParams) middleware.Responder {
	return fn(params)
}

// PostRegisterHandler interface for that can handle valid post register params
type PostRegisterHandler interface {
	Handle(PostRegisterParams) middleware.Responder
}

// NewPostRegister creates a new http.Handler for the post register operation
func NewPostRegister(ctx *middleware.Context, handler PostRegisterHandler) *PostRegister {
	return &PostRegister{Context: ctx, Handler: handler}
}

/*PostRegister swagger:route POST /register Authentication postRegister

Create a user

*/
type PostRegister struct {
	Context *middleware.Context
	Handler PostRegisterHandler
}

func (o *PostRegister) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostRegisterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

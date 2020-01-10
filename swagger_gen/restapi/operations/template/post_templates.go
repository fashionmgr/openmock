// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostTemplatesHandlerFunc turns a function with the right signature into a post templates handler
type PostTemplatesHandlerFunc func(PostTemplatesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTemplatesHandlerFunc) Handle(params PostTemplatesParams) middleware.Responder {
	return fn(params)
}

// PostTemplatesHandler interface for that can handle valid post templates params
type PostTemplatesHandler interface {
	Handle(PostTemplatesParams) middleware.Responder
}

// NewPostTemplates creates a new http.Handler for the post templates operation
func NewPostTemplates(ctx *middleware.Context, handler PostTemplatesHandler) *PostTemplates {
	return &PostTemplates{Context: ctx, Handler: handler}
}

/*PostTemplates swagger:route POST /templates template postTemplates

Modify templates in the base model by POSTing new ones

*/
type PostTemplates struct {
	Context *middleware.Context
	Handler PostTemplatesHandler
}

func (o *PostTemplates) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostTemplatesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

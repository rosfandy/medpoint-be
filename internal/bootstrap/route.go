// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden"
	"medpointbe/internal/controllers"
	"github.com/valyala/fasthttp"
)

func RegisterRoute(server *raiden.Server) {
	server.RegisterRoute([]*raiden.Route{
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/hello/{name}",
			Methods:    []string{fasthttp.MethodGet},
			Controller: &controllers.HelloWordController{},
		},
	})
}

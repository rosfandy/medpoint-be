// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden"
	raiden_controllers "github.com/sev-2/raiden/pkg/controllers"
	"github.com/valyala/fasthttp"
	"medpointbe/internal/controllers"
)

func RegisterRoute(server *raiden.Server) {
	server.RegisterRoute([]*raiden.Route{
		raiden.NewRouteFromController(&raiden_controllers.StateReadyController{}, []string{fasthttp.MethodPost}),
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/hello/{name}",
			Methods:    []string{fasthttp.MethodGet},
			Controller: &controllers.HelloWordController{},
		},
	})
}

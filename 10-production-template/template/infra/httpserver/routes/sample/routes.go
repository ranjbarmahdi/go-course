package sample

import (
	"net/http"

	"template/infra/httpserver/middlewares"
	"template/infra/httpserver/routes"
)

const Mount = "/samples"

func RegisterRoutes(mux *http.ServeMux, apiPrefix string, handler *Handler) {
	base := routes.Join(apiPrefix, Mount)

	routes.Register(mux, routes.POST, base, "", handler.Create, middlewares.UserHeader)
	routes.Register(mux, routes.GET, base, "/{id}", handler.GetByID, middlewares.UserHeader)
}

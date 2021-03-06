// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"datac/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: DatacHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/from2/:name",
				Handler: DatacHandler2Handler(serverCtx),
			},
		},
	)
}

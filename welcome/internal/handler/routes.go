// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"welcome-manage/welcome/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/active/title",
				Handler: SetTitleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/active/title",
				Handler: GetTitleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: UploadHandler(serverCtx),
			},
		},
	)
}

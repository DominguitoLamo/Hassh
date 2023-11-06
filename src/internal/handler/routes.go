// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	sshtask "hassh/src/internal/handler/sshtask"
	"hassh/src/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/echo",
				Handler: sshtask.EchoMsgHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/sshtask",
				Handler: sshtask.GetSshTaskInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sshtask",
				Handler: sshtask.AddSshTaskInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/sshtask",
				Handler: sshtask.UpdateSshTaskInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/sshtask",
				Handler: sshtask.DeleteSshTaskInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/switchBrand",
				Handler: sshtask.GetSwitchBrandHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/runCmd",
				Handler: sshtask.RunCmdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getCmdResult",
				Handler: sshtask.DownloadCmdResultHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
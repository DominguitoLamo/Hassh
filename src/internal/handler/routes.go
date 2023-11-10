// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	sshtask "hassh/src/internal/handler/sshtask"
	groupInfo "hassh/src/internal/handler/groupInfo"
	"hassh/src/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	addSshTaskRoutes(server, serverCtx)
	addGroupInfo(server, serverCtx)
}

func addSshTaskRoutes(server *rest.Server, serverCtx *svc.ServiceContext) {
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
			{
				Method:  http.MethodGet,
				Path:    "/fileExist",
				Handler: sshtask.FileExistHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}

func addGroupInfo(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/groupInfo",
				Handler: groupInfo.GetGroupInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/groupInfo",
				Handler: groupInfo.AddGroupInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/groupInfo",
				Handler: groupInfo.UpdateGroupInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/groupInfo",
				Handler: groupInfo.DeleteGroupInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/groupTask",
				Handler: groupInfo.AddGroupTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/groupTask",
				Handler: groupInfo.DeleteGroupTaskHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/groupNames",
				Handler: groupInfo.GetGroupNamesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/runGroupTasks",
				Handler: groupInfo.RunGroupTasksHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/groupFileExist",
				Handler: groupInfo.GroupFileExistHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/downloadGroupFile",
				Handler: groupInfo.DownloadGroupFileHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
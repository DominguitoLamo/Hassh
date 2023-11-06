package sshtask

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hassh/src/internal/logic/sshtask"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
)

func DownloadCmdResultHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadCmdResultReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sshtask.NewDownloadCmdResultLogic(r.Context(), svcCtx)
		err := l.DownloadCmdResult(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

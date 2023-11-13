package sshtask

import (
	"fmt"
	"net/http"

	"hassh/src/internal/logic/sshtask"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadCmdResultHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadCmdResultReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sshtask.NewDownloadCmdResultLogic(r.Context(), svcCtx)
		resp, err := l.DownloadCmdResult(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.Write(resp.Content)
			attach := fmt.Sprintf("attachment; filename=\"%s\"", resp.Name)
			w.Header().Add("content-disposition", attach)
			w.Header().Set("content-type", "text/plain")
			httpx.Ok(w)

			defer func() {
				svcCtx.Components.SSHResultManager.DeleteFile(req.Id)
			}()
		}
	}
}

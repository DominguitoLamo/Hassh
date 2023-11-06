package sshtask

import (
	"net/http"

	"hassh/src/internal/logic/sshtask"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RunCmdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RunCmdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sshtask.NewRunCmdLogic(r.Context(), svcCtx)
		resp, err := l.RunCmd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

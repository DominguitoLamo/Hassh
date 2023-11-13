package sshtask

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hassh/src/internal/logic/sshtask"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
)

func EchoMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EchoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := sshtask.NewEchoMsgLogic(r.Context(), svcCtx)
		resp, err := l.EchoMsg(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

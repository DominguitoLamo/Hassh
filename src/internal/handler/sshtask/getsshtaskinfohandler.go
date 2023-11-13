package sshtask

import (
	"net/http"

	"hassh/src/internal/logic/sshtask"
	"hassh/src/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSshTaskInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sshtask.NewGetSshTaskInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetSshTaskInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

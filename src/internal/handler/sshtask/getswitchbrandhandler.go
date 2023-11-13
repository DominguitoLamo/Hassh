package sshtask

import (
	"net/http"

	"hassh/src/internal/logic/sshtask"
	"hassh/src/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSwitchBrandHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sshtask.NewGetSwitchBrandLogic(r.Context(), svcCtx)
		resp, err := l.GetSwitchBrand()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

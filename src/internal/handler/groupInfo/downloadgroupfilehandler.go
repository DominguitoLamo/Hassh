package groupInfo

import (
	"net/http"

	"hassh/src/internal/logic/groupInfo"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadGroupFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadGroupFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := groupInfo.NewDownloadGroupFileLogic(r.Context(), svcCtx)
		err := l.DownloadGroupFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

package groupInfo

import (
	"net/http"

	"hassh/src/internal/logic/groupInfo"
	"hassh/src/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetGroupInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := groupInfo.NewGetGroupInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetGroupInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

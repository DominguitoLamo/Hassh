package groupInfo

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"

	"hassh/src/internal/components"
	"hassh/src/internal/logic/groupInfo"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/logger"

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
		result, err := l.DownloadGroupFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			zipWrite(w, result)
			httpx.Ok(w)
			defer func() {
				svcCtx.Components.GroupResultManager.DeleteFile(result.Key)
			}()
		}
	}
}

func zipWrite(w http.ResponseWriter, result *components.GroupTaskResult) (error) {
	attach := fmt.Sprintf("attachment; filename=\"%s.zip\"", result.Key)
	w.Header().Set("content-type", "application/zip")
	w.Header().Add("content-disposition", attach)

	zipWriter := zip.NewWriter(w)
	for _, item := range result.Details {
		f, err := zipWriter.Create(item.Name + ".txt")
		if (err != nil) {
			logger.ErrorLog("zip writer: %s", err)
			return err
		}
		io.WriteString(f, item.Content)
	}
	zipWriter.Close()
	return nil
}

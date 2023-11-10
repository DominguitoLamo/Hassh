package groupInfo

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadGroupFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadGroupFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadGroupFileLogic {
	return &DownloadGroupFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadGroupFileLogic) DownloadGroupFile(req *types.DownloadGroupFileReq) error {
	// todo: add your logic here and delete this line

	return nil
}

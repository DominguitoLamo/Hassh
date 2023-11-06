package sshtask

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadCmdResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadCmdResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadCmdResultLogic {
	return &DownloadCmdResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadCmdResultLogic) DownloadCmdResult(req *types.DownloadCmdResultReq) error {
	// todo: add your logic here and delete this line

	return nil
}

package sshtask

import (
	"context"

	"hassh/src/internal/components"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/internal/utils"

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

func (l *DownloadCmdResultLogic) DownloadCmdResult(req *types.DownloadCmdResultReq) (*components.CacheFile, error) {
	sshResult := l.svcCtx.Components.SSHResultManager
	file := sshResult.GetFile(req.Id)

	if (file == nil) {
		return nil, utils.SSHCmdError("no file exist")
	}

	return file, nil
}

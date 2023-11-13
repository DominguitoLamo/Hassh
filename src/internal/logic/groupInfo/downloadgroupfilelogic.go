package groupInfo

import (
	"context"

	"hassh/src/internal/components"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/internal/utils"

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

func (l *DownloadGroupFileLogic) DownloadGroupFile(req *types.DownloadGroupFileReq) (resp *components.GroupTaskResult, err error) {
	isExist := l.svcCtx.Components.GroupResultManager.IsExist(req.Id)
	if !isExist {
		err = utils.SSHCmdError("File doesn't exist")
		return
	}
	result := l.svcCtx.Components.GroupResultManager.GetFile(req.Id)
	if result.ErrMsg != "" {
		err = utils.SSHCmdError(result.ErrMsg)
		return
	}

	resp = &result
	return
}

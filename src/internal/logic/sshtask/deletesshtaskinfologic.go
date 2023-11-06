package sshtask

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSshTaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSshTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSshTaskInfoLogic {
	return &DeleteSshTaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSshTaskInfoLogic) DeleteSshTaskInfo(req *types.DELETESSHInfoReq) (resp *types.DELETESSHInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}

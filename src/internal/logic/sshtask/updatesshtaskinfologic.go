package sshtask

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSshTaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSshTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSshTaskInfoLogic {
	return &UpdateSshTaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSshTaskInfoLogic) UpdateSshTaskInfo(req *types.UpdateSSHInfoReq) (resp *types.UpdateSSHInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}

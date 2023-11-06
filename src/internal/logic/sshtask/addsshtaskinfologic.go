package sshtask

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSshTaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSshTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSshTaskInfoLogic {
	return &AddSshTaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSshTaskInfoLogic) AddSshTaskInfo(req *types.AddSSHInfoReq) (resp *types.AddSSHInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}

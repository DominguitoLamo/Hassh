package sshtask

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RunCmdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRunCmdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunCmdLogic {
	return &RunCmdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RunCmdLogic) RunCmd(req *types.RunCmdReq) (resp *types.RunCmdResp, err error) {
	// todo: add your logic here and delete this line

	return
}

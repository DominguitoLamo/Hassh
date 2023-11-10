package groupInfo

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RunGroupTasksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRunGroupTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunGroupTasksLogic {
	return &RunGroupTasksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RunGroupTasksLogic) RunGroupTasks(req *types.RunGroupTasksReq) (resp *types.RunGroupTasksResp, err error) {
	// todo: add your logic here and delete this line

	return
}

package groupInfo

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGroupTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupTaskLogic {
	return &DeleteGroupTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGroupTaskLogic) DeleteGroupTask(req *types.DeleteGroupTaskReq) (resp *types.DeleteGroupTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}

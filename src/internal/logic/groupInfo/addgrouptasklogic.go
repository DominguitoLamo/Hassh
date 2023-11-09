package groupInfo

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGroupTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupTaskLogic {
	return &AddGroupTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGroupTaskLogic) AddGroupTask(req *types.AddGroupTaskReq) (resp *types.AddGroupTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}

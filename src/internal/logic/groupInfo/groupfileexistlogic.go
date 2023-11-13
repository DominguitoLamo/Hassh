package groupInfo

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupFileExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupFileExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupFileExistLogic {
	return &GroupFileExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupFileExistLogic) GroupFileExist(req *types.GroupFileExistReq) (resp *types.GroupFileExistResp, err error) {
	resp = new(types.GroupFileExistResp)
	resp.IsExist = l.svcCtx.Components.GroupResultManager.IsExist(req.Id)
	return
}

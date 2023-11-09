package groupInfo

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupInfoLogic {
	return &DeleteGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGroupInfoLogic) DeleteGroupInfo(req *types.DeleteGroupInfoReq) (resp *types.DeleteGroupInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}

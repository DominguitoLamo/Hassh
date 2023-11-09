package groupInfo

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"


	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupInfoLogic {
	return &AddGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGroupInfoLogic) AddGroupInfo(req *types.AddGroupInfoReq) (resp *types.AddGroupInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}

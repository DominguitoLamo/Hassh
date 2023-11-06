package sshtask

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSshTaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSshTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSshTaskInfoLogic {
	return &GetSshTaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSshTaskInfoLogic) GetSshTaskInfo(req *types.GETSSHInfoReq) (resp []types.GETSSHInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}

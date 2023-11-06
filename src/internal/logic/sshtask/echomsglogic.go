package sshtask

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EchoMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEchoMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EchoMsgLogic {
	return &EchoMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EchoMsgLogic) EchoMsg(req *types.EchoReq) (resp *types.EchoResp, err error) {
	// todo: add your logic here and delete this line

	return
}

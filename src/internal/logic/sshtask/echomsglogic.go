package sshtask

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/internal/utils"

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
	resp = new(types.EchoResp)

	if (req.Msg != "echo") {
		err = utils.ParameterError()
	} else {
		resp.Msg = req.Msg
	}
	return
}

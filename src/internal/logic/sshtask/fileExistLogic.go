package sshtask

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileExistLogic {
	return &FileExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileExistLogic) FileExistLogic(req *types.FileExistReq) (resp *types.FileExistResp, err error) {
	id := req.Id
	sshResults := l.svcCtx.Components.SSHResultManager
	resp = new(types.FileExistResp)
	resp.IsExist = sshResults.IsExist(id)

	return resp, err
}

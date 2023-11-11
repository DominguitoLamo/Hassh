package sshtask

import (
	"context"

	sshtask "hassh/src/internal/model/sshTaskModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/jinzhu/copier"
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

func (l *GetSshTaskInfoLogic) GetSshTaskInfo() (resp *[]*types.GETSSHInfoResp, err error) {
	dao := sshtask.NewSshTaskModel(l.svcCtx.Components.DbConnection)
	sqlResult, sqlErr := dao.SelectItems(l.ctx)
	if sqlErr != nil {
		logx.Error("sql error: ", sqlErr.Error())
		err = sqlErr
	}

	resp = new([]*types.GETSSHInfoResp)
	copier.Copy(resp, sqlResult)
	return
}

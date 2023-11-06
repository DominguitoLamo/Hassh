package sshtask

import (
	"context"

	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSwitchBrandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSwitchBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSwitchBrandLogic {
	return &GetSwitchBrandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSwitchBrandLogic) GetSwitchBrand() (resp *types.SwitchBrandResp, err error) {
	// todo: add your logic here and delete this line

	return
}

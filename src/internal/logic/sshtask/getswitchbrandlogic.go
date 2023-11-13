package sshtask

import (
	"context"
	switchbrand "hassh/src/internal/model/switchBrandModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/jinzhu/copier"
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

func (l *GetSwitchBrandLogic) GetSwitchBrand() (resp []*types.SwitchBrandResp, err error) {
	switchDao := switchbrand.NewSwitchBrandModel(l.svcCtx.Components.DbConnection)
	result, err := switchDao.FindAll(l.ctx)
	err = copier.Copy(&resp, &result)
	return
}

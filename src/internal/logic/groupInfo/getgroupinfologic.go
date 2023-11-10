package groupInfo

import (
	"context"

	groupInfo "hassh/src/internal/model/groupInfoModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoLogic {
	return &GetGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupInfoLogic) GetGroupInfo() (resp *[]types.GetGroupInfoResp, err error) {
	resp = new([]types.GetGroupInfoResp)
	dao := groupInfo.NewGroupInfoModel(l.svcCtx.Components.DbConnection)
	dataResult, dbErr := dao.SelectGroupDetail(l.ctx)
	if (dbErr != nil) {
		err = dbErr
		return
	}

	for _, item := range dataResult {
		var i types.GetGroupInfoResp
		copier.Copy(&i, &item)
		*resp = append(*resp, i)
	}
	return
}

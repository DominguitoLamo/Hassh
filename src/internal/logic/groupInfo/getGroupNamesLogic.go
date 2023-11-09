package groupInfo

import (
	"context"

	groupInfo "hassh/src/internal/model/groupInfoModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/logger"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupNamesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupNamesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoLogic {
	return &GetGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupInfoLogic) GetGroupNamesLogic() (resp *[]*types.GetGroupName, err error) {
	dao := groupInfo.NewGroupInfoModel(l.svcCtx.Components.DbConnection)
	result, dbErr := dao.SelectAll(l.ctx)
	if (dbErr != nil) {
		err =dbErr
		logger.ErrorLog("db err %s", err.Error())
	}
	resp = new([]*types.GetGroupName)
	copier.Copy(resp, result)
	return
}

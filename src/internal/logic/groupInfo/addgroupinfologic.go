package groupInfo

import (
	"context"

	groupInfo "hassh/src/internal/model/groupInfoModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/internal/utils"
	"hassh/src/logger"

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
	if req.Name == "" {
		err = utils.ParameterError()
	}
	dao := groupInfo.NewGroupInfoModel(l.svcCtx.Components.DbConnection)
	item := new(groupInfo.GroupInfo)
	item.Name = req.Name
	result, sqlErr := dao.Insert(l.ctx, item)
	if sqlErr != nil {
		logger.ErrorLog("db error: %s", sqlErr.Error())
	}
	
	resp = new(types.AddGroupInfoResp)
	id, _ := result.LastInsertId()
	resp.Id = id
	return
}

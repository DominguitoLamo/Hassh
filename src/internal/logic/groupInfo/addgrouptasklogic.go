package groupInfo

import (
	"context"

	groupInfo "hassh/src/internal/model/groupInfoModel"
	groupTask "hassh/src/internal/model/groupTasksModel"
	sshtask "hassh/src/internal/model/sshTaskModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/internal/utils"
	"hassh/src/logger"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGroupTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupTaskLogic {
	return &AddGroupTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGroupTaskLogic) AddGroupTask(req *types.AddGroupTaskReq) (resp *types.AddGroupTaskResp, err error) {
	sshDao := sshtask.NewSshTaskModel(l.svcCtx.Components.DbConnection)
	groupDao := groupInfo.NewGroupInfoModel(l.svcCtx.Components.DbConnection)

	sshItem, _ := sshDao.FindOne(l.ctx, req.TaskId)
	groupItem, _ := groupDao.FindOne(l.ctx, req.GroupId)

	if sshItem == nil || groupItem == nil {
		err = utils.ParameterError()
		return
	}

	groupTaskDao := groupTask.NewGroupTasksModel(l.svcCtx.Components.DbConnection)
	dataInserted := new(groupTask.GroupTasks)
	copier.Copy(dataInserted, req)
	result, dbErr := groupTaskDao.Insert(l.ctx, dataInserted)

	if (dbErr != nil) {
		err = dbErr
		logger.ErrorLog(dbErr.Error())
	}
	id, _ := result.LastInsertId()
	resp = new(types.AddGroupTaskResp)
	resp.Id = id
	return
}

package groupInfo

import (
	"context"

	groupTask "hassh/src/internal/model/groupTasksModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGroupTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupTaskLogic {
	return &DeleteGroupTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGroupTaskLogic) DeleteGroupTask(req *types.DeleteGroupTaskReq) (err error) {
	dao := groupTask.NewGroupTasksModel(l.svcCtx.Components.DbConnection)
	item := new(groupTask.GroupTasks)
	copier.Copy(item, req)
	dao.Delete(l.ctx, item)
	return
}

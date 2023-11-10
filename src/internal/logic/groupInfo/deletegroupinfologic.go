package groupInfo

import (
	"context"

	groupInfo "hassh/src/internal/model/groupInfoModel"
	groupTask "hassh/src/internal/model/groupTasksModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DeleteGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupInfoLogic {
	return &DeleteGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGroupInfoLogic) DeleteGroupInfo(req *types.DeleteGroupInfoReq) (resp *types.DeleteGroupInfoResp, err error) {
	conn := l.svcCtx.Components.DbConnection
	transErr := conn.TransactCtx(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		groupTaskDao := groupTask.NewGroupTasksModel(conn)
		taskErr := groupTaskDao.DeleteByGroupId(ctx, req.Id)
		if taskErr != nil {
			return taskErr
		}

		groupInfoDao := groupInfo.NewGroupInfoModel(conn)
		infoErr := groupInfoDao.Delete(ctx, req.Id)
		if infoErr != nil {
			return infoErr
		}

		return nil
	})

	if (transErr != nil) {
		err = transErr
		return
	}

	resp = new(types.DeleteGroupInfoResp)
	resp.Id = req.Id
	return
}

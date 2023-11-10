package sshtask

import (
	"context"

	groupTask "hassh/src/internal/model/groupTasksModel"
	sshtask "hassh/src/internal/model/sshTaskModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DeleteSshTaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSshTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSshTaskInfoLogic {
	return &DeleteSshTaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSshTaskInfoLogic) DeleteSshTaskInfo(req *types.DELETESSHInfoReq) (resp *types.DELETESSHInfoResp, err error) {
	conn := l.svcCtx.Components.DbConnection
	transErr := conn.TransactCtx(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		groupTaskDao := groupTask.NewGroupTasksModel(conn)
		taskErr := groupTaskDao.DeleteByTaskId(ctx, req.Id)
		if taskErr != nil {
			return taskErr
		}

		sshTaskDao := sshtask.NewSshTaskModel(conn)
		sshErr := sshTaskDao.Delete(ctx, req.Id)
		if sshErr != nil {
			return taskErr
		}
		return nil
	})
	if (transErr != nil) {
		err = transErr
		return
	}

	resp = new(types.DELETESSHInfoResp)
	resp.Id = req.Id
	return
}

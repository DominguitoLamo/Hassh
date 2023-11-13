package sshtask

import (
	"context"

	sshtask "hassh/src/internal/model/sshTaskModel"
	switchbrand "hassh/src/internal/model/switchBrandModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSshTaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSshTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSshTaskInfoLogic {
	return &UpdateSshTaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSshTaskInfoLogic) UpdateSshTaskInfo(req *types.UpdateSSHInfoReq) (resp *types.UpdateSSHInfoResp, err error) {
	// ip valid
	if err = utils.IpFormatValid(req.Ip); err != nil {
		return
	}

	// brand valid
	if brandErr := l.brandValid(req); brandErr != nil {
		err = brandErr
		return
	}

	resp = new(types.UpdateSSHInfoResp)
	dao := sshtask.NewSshTaskModel(l.svcCtx.Components.DbConnection)
	err = dao.Update(l.ctx, (*sshtask.SshTask)(req))
	if (err == nil) {
		resp.Id = req.Id
		return
	}
	return
}

func (l *UpdateSshTaskInfoLogic) brandValid(req *types.UpdateSSHInfoReq) error {
	switchDao := switchbrand.NewSwitchBrandModel(l.svcCtx.Components.DbConnection)
	if result, err := switchDao.FindAll(l.ctx); err != nil {
		return err
	} else {
		for _, item := range result {
			if item.Name == req.Brand {
				return nil
			}
		}
	}

	return utils.ParameterError()
}
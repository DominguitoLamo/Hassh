package sshtask

import (
	"context"

	sshtask "hassh/src/internal/model/sshTaskModel"
	switchbrand "hassh/src/internal/model/switchBrandModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/internal/utils"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddSshTaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSshTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSshTaskInfoLogic {
	return &AddSshTaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSshTaskInfoLogic) AddSshTaskInfo(req *types.AddSSHInfoReq) (resp *types.AddSSHInfoResp, err error) {
	// ip valid
	if err = utils.IpFormatValid(req.Ip); err != nil {
		return
	}

	// brand valid
	if brandErr := l.brandValid(req); brandErr != nil {
		err = brandErr
		return
	}

	insertedItem := new(sshtask.SshTask)
	copier.Copy(insertedItem, req)

	dao := sshtask.NewSshTaskModel(l.svcCtx.Components.DbConnection)
	result, sqlErr := dao.Insert(l.ctx, insertedItem)
	if (sqlErr != nil) {
		logx.Error(sqlErr.Error())
		err = sqlErr
		return
	}

	resp = new(types.AddSSHInfoResp)
	if id, idErr := result.LastInsertId(); idErr != nil {
		logx.Error(idErr.Error())
		err = idErr
	} else {
		resp.Id = id
	}
	return
}

func (l *AddSshTaskInfoLogic) brandValid(req *types.AddSSHInfoReq) error {
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
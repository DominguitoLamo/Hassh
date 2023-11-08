package sshtask

import (
	"context"
	"strings"

	"hassh/src/internal/constant"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/internal/utils"

	switchgo "github.com/DominguitoLamo/switchGo"
	"github.com/zeromicro/go-zero/core/logx"
)

type RunCmdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRunCmdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunCmdLogic {
	return &RunCmdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RunCmdLogic) RunCmd(req *types.RunCmdReq) (resp *types.RunCmdResp, err error) {
	sshManager := l.svcCtx.Components.SSHManager

	config, configErr := switchgo.SSHConfigCreate(req.Account, req.Password, req.Ip, constant.SSH_DEFAULT_PORT, req.Brand)
	if configErr != nil {
		err = utils.ConfigError(configErr.Error())
		return
	}

	session, sessionErr := sshManager.GetSSHSession(config)
	if sessionErr != nil {
		err = utils.SSHConnectionError(sessionErr.Error())
		return
	}
	
	cmdResult, cmdError := session.RunCmdsAndClose(strings.Split(req.Script, ";")...)
	if cmdError != nil {
		err = utils.SSHCmdError(cmdError.Error())
		return
	}

	sshResultManager := l.svcCtx.Components.SSHResultManager
	saveKey, saveErr := sshResultManager.SaveResult(req, session, cmdResult)
	if saveErr != nil {
		err = utils.SSHCmdError(saveErr.Error())
		return
	}

	resp = new(types.RunCmdResp)
	resp.Id = saveKey
	
	return
}

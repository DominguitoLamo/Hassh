package sshtask

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"time"

	"hassh/src/internal/constant"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/internal/utils"
	taskqueue "hassh/src/taskQueue"

	switchgo "github.com/DominguitoLamo/switchGo"
	"github.com/zeromicro/go-zero/core/logc"
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

	saveKey, hashErr := genhashKey(req)

	if (hashErr != nil) {
		err = utils.SSHCmdError("key generated error")
	}

	resp = new(types.RunCmdResp)
	resp.Id = saveKey

	taskReq := taskqueue.NewSshQueueRequest(req, sshManager)
	taskReq.ExecuteTask = func() {
		result := ""
		session, sessionErr := sshManager.GetSSHSession(config)
		if sessionErr != nil {
			result = sessionErr.Error()
			logx.Error("session error: ", result)
			l.saveFile(session, req, saveKey, result)
			return
		}
		
		cmdResult, cmdError := session.RunCmds(strings.Split(req.Script, ";")...)
		if cmdError != nil {
			result = cmdError.Error()
			logx.Error("cmd error: ", result)
			l.saveFile(session, req, saveKey, result)
			return
		}
		result = cmdResult
		// logx.Debug("cmd return result: ", result)
		logc.Debug(l.ctx, "cmd return result: %s", result)
		l.saveFile(session, req, saveKey, result)
	}
	l.svcCtx.Queues.SSHQueue.AddTask(taskReq)
	
	return
}

func (l *RunCmdLogic) saveFile(session *switchgo.SSHSession, req *types.RunCmdReq, key string, result string) {
	sshResultManager := l.svcCtx.Components.SSHResultManager
	saveErr := sshResultManager.SaveResult(session, req, key, result)
	if saveErr != nil {
		logx.Error("Write File Error: ", saveErr.Error())
	}
}

func genhashKey(req *types.RunCmdReq) (string, error) {
	hasher := md5.New()
	keyStr := fmt.Sprintf("%s%s%s", req.Name, req.Ip, time.Now().String())
	_, err := io.WriteString(hasher, keyStr)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil))[:8], nil
}

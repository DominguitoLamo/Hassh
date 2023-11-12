package groupInfo

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
	"time"

	"hassh/src/internal/components"
	"hassh/src/internal/constant"
	groupInfo "hassh/src/internal/model/groupInfoModel"
	sshtask "hassh/src/internal/model/sshTaskModel"
	"hassh/src/internal/svc"
	"hassh/src/internal/types"
	"hassh/src/logger"

	switchgo "github.com/DominguitoLamo/switchGo"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RunGroupTasksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRunGroupTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunGroupTasksLogic {
	return &RunGroupTasksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RunGroupTasksLogic) RunGroupTasks(req *types.RunGroupTasksReq) (resp *types.RunGroupTasksResp, err error) {
	resp = new(types.RunGroupTasksResp)
	resp.Id = getHashKey()

	executeTask := func() {
		logger.DebugLog("%s start to run group task", resp.Id)
		var result components.GroupTaskResult
		result.Key = resp.Id
		result.ErrMsg = ""

		groupId := req.Id
		dao := groupInfo.NewGroupInfoModel(l.svcCtx.Components.DbConnection)
		detail, dbErr := dao.SelectGroupDetailById(context.Background(), groupId)
		if (dbErr != nil) {
			result.ErrMsg = dbErr.Error()
		}
		detailLen := len(*detail.Tasks)
		result.Details = make([]components.SSHTaskDetail, 0)
		resultChan := make(chan components.SSHTaskDetail, detailLen)

		for _, item := range *detail.Tasks {
			var clone sshtask.SshTask
			copier.Copy(&clone, item)
			go func() {
				logger.DebugLog("%s start to run cmd task", clone.Name)
				var sshTask components.SSHTaskDetail
				sshTask.Name = clone.Name

				sshManager := l.svcCtx.Components.SSHManager
				config, _ := switchgo.SSHConfigCreate(clone.Account, clone.Password, clone.Ip, constant.SSH_DEFAULT_PORT, clone.Brand)
				session, sessionErr := sshManager.GetSSHSession(config)
				if (sessionErr != nil) {
					sshTask.Content += sessionErr.Error() + "\n"
					resultChan <- sshTask
					return
				}

				cmds := strings.Split(clone.Script, ";")
				cmdResult, cmdErr := session.RunCmds(cmds...)
				if (cmdErr != nil) {
					sshTask.Content += cmdErr.Error() + "\n"
					resultChan <- sshTask
					return
				}
				sshTask.Content = cmdResult
				// logger.DebugLog("%s result: %s", clone.Name, cmdResult)
				resultChan <- sshTask
			}()
		}

		for item := range resultChan {
			logger.DebugLog("%s received: %s", item.Name, item.Content)
			result.Details = append(result.Details, item)

			if len(result.Details) == detailLen {
				break
			}
		}
		logger.DebugLog("%s task finished", result.Key)
		l.svcCtx.Components.GroupResultManager.SaveResult(result)
	}

	l.svcCtx.Queues.GroupQueue.AddTask(executeTask)
	return
}

func getHashKey() (string) {
	hasher := md5.New()
	keyStr := time.Now().String()
	io.WriteString(hasher, keyStr)

	return hex.EncodeToString(hasher.Sum(nil))[:8]
}
package svc

import (
	"hassh/src/internal/components"
	"hassh/src/internal/config"
	taskqueue "hassh/src/taskQueue"
	"os"
	"path/filepath"

	switchgo "github.com/DominguitoLamo/switchGo"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type CustomConfigStruct struct {
	DatabaseKey string
}

type ComponentRegister struct {
	DbConnection sqlx.SqlConn
	SSHManager *switchgo.SessionManager
	SSHResultManager *components.SshResultManager
}

type QueueRegister struct {
	SSHQueue taskqueue.SshTaskQueue
}

type ServiceContext struct {
	execPath string
	Config config.Config
	CustomConfig CustomConfigStruct
	Components *ComponentRegister
	Queues *QueueRegister
}

func NewServiceContext(c config.Config, custom CustomConfigStruct) *ServiceContext {
	component := new(ComponentRegister)
	queues := new(QueueRegister)
	return &ServiceContext{
		Config: c,
		CustomConfig: custom,
		Components: component,
		Queues: queues,
	}
}

func (ctx *ServiceContext) SetExecPath() {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exePath := filepath.Dir(ex)
	ctx.execPath = exePath
}

func (ctx *ServiceContext) GetTmpPath() string {
	return filepath.Join(ctx.execPath, "tmp")
}

func InitQueues(ctx *ServiceContext) {
	ctx.Queues.SSHQueue = *taskqueue.NewSshTaskQueue()
	ctx.Queues.SSHQueue.RunTask()
}

func InitComponents(ctx *ServiceContext) {
	InitDBConnection(ctx)
	InitSSHSession(ctx)
	InitSSHResultManager(ctx)
}

func InitSSHResultManager(ctx *ServiceContext) {
	manager := components.NewSSHResultManager()
	ctx.Components.SSHResultManager = manager
}

func InitDBConnection(ctx *ServiceContext) {
	dbConnection := sqlx.NewMysql(ctx.CustomConfig.DatabaseKey)
	ctx.Components.DbConnection = dbConnection
}

func InitSSHSession(ctx *ServiceContext) {
	sshManager := switchgo.NewSessionManager()
	ctx.Components.SSHManager = sshManager
}

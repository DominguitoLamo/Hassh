package svc

import (
	"hassh/src/internal/components"
	"hassh/src/internal/config"

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

type ServiceContext struct {
	Config config.Config
	CustomConfig CustomConfigStruct
	Components *ComponentRegister
}

func NewServiceContext(c config.Config, custom CustomConfigStruct) *ServiceContext {
	component := new(ComponentRegister)
	return &ServiceContext{
		Config: c,
		CustomConfig: custom,
		Components: component,
	}
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

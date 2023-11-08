package svc

import (
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

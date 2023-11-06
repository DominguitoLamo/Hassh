package components

import (
	"hassh/src/internal/svc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func InitDBConnection(ctx *svc.ServiceContext) {
	dbConnection := sqlx.NewMysql(ctx.CustomConfig.DatabaseKey)
	ctx.Components.DbConnection = dbConnection
}


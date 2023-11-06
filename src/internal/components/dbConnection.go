package components

import (
	"hassh/src/internal/svc"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func InitDBConnection(ctx *svc.ServiceContext) {
	dbConnection := sqlx.NewMysql("root:moto9171@/demo?parseTime=true")
	ctx.Components.DbConnection = dbConnection
}


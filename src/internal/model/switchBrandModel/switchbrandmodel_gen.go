package switchbrand

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	switchBrandFieldNames          = builder.RawFieldNames(&SwitchBrand{})
	switchBrandRows                = strings.Join(switchBrandFieldNames, ",")
)

type (
	switchBrandModel interface {
		FindAll(ctx context.Context) ([]*SwitchBrand, error)
	}

	defaultSwitchBrandModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SwitchBrand struct {
		Id   int64  `db:"id"` // Primary Key
		Name string `db:"name"`
	}
)

func newSwitchBrandModel(conn sqlx.SqlConn) *defaultSwitchBrandModel {
	return &defaultSwitchBrandModel{
		conn:  conn,
		table: "`switch_brand`",
	}
}

func (m *defaultSwitchBrandModel) FindAll(ctx context.Context) ([]*SwitchBrand, error) {
	sql := fmt.Sprintf("select %s from %s", switchBrandRows, m.table)
	var resp []*SwitchBrand
	err := m.conn.QueryRowsCtx(ctx, &resp, sql)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}


func (m *defaultSwitchBrandModel) tableName() string {
	return m.table
}

package switchbrand

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	switchBrandFieldNames          = builder.RawFieldNames(&SwitchBrand{})
	switchBrandRows                = strings.Join(switchBrandFieldNames, ",")
	switchBrandRowsExpectAutoSet   = strings.Join(stringx.Remove(switchBrandFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	switchBrandRowsWithPlaceHolder = strings.Join(stringx.Remove(switchBrandFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	switchBrandModel interface {
		Insert(ctx context.Context, data *SwitchBrand) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SwitchBrand, error)
		Update(ctx context.Context, data *SwitchBrand) error
		Delete(ctx context.Context, id int64) error
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

func (m *defaultSwitchBrandModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSwitchBrandModel) FindOne(ctx context.Context, id int64) (*SwitchBrand, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", switchBrandRows, m.table)
	var resp SwitchBrand
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSwitchBrandModel) Insert(ctx context.Context, data *SwitchBrand) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, switchBrandRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name)
	return ret, err
}

func (m *defaultSwitchBrandModel) Update(ctx context.Context, data *SwitchBrand) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, switchBrandRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Id)
	return err
}

func (m *defaultSwitchBrandModel) tableName() string {
	return m.table
}

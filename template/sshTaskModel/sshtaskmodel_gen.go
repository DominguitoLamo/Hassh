// Code generated by goctl. DO NOT EDIT.

package template

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
	sshTaskFieldNames          = builder.RawFieldNames(&SshTask{})
	sshTaskRows                = strings.Join(sshTaskFieldNames, ",")
	sshTaskRowsExpectAutoSet   = strings.Join(stringx.Remove(sshTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sshTaskRowsWithPlaceHolder = strings.Join(stringx.Remove(sshTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	sshTaskModel interface {
		Insert(ctx context.Context, data *SshTask) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SshTask, error)
		Update(ctx context.Context, data *SshTask) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSshTaskModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SshTask struct {
		Id       int64  `db:"id"` // Primary Key
		Ip       string `db:"ip"`
		Name     string `db:"name"`
		Account  string `db:"account"`
		Password string `db:"password"`
		Brand    string `db:"brand"`  // only for cisco, huawei
		Script   string `db:"script"` // script for task. Statements are separated with ;
		Desc     string `db:"desc"`   // task desc
	}
)

func newSshTaskModel(conn sqlx.SqlConn) *defaultSshTaskModel {
	return &defaultSshTaskModel{
		conn:  conn,
		table: "`ssh_task`",
	}
}

func (m *defaultSshTaskModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSshTaskModel) FindOne(ctx context.Context, id int64) (*SshTask, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sshTaskRows, m.table)
	var resp SshTask
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

func (m *defaultSshTaskModel) Insert(ctx context.Context, data *SshTask) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, sshTaskRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Ip, data.Name, data.Account, data.Password, data.Brand, data.Script, data.Desc)
	return ret, err
}

func (m *defaultSshTaskModel) Update(ctx context.Context, data *SshTask) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sshTaskRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Ip, data.Name, data.Account, data.Password, data.Brand, data.Script, data.Desc, data.Id)
	return err
}

func (m *defaultSshTaskModel) tableName() string {
	return m.table
}

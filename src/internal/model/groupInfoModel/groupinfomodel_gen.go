// Code generated by goctl. DO NOT EDIT.

package groupInfo

import (
	"context"
	"database/sql"
	"fmt"
	groupTask "hassh/src/internal/model/groupTasksModel"
	sshtask "hassh/src/internal/model/sshTaskModel"
	"hassh/src/logger"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	groupInfoFieldNames          = builder.RawFieldNames(&GroupInfo{})
	groupInfoRows                = strings.Join(groupInfoFieldNames, ",")
	groupInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(groupInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	groupInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(groupInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	groupInfoModel interface {
		Insert(ctx context.Context, data *GroupInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*GroupInfo, error)
		Update(ctx context.Context, data *GroupInfo) error
		Delete(ctx context.Context, id int64) error
		SelectAll(ctx context.Context) (*[]*GroupInfo, error)
		SelectGroupDetail(ctx context.Context) (resp []GroupDetail, err error)
		SelectGroupDetailById(ctx context.Context, groupId int64) (resp GroupDetail, err error)
	}

	defaultGroupInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	GroupInfo struct {
		Id   int64  `db:"id"` // Primary Key
		Name string `db:"name"`
	}

	GroupDetail struct {
		Id   int64
		Name string
		Tasks *[]*sshtask.SshTask 
	}
)

func newGroupInfoModel(conn sqlx.SqlConn) *defaultGroupInfoModel {
	return &defaultGroupInfoModel{
		conn:  conn,
		table: "`group_info`",
	}
}

func (m *defaultGroupInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGroupInfoModel) FindOne(ctx context.Context, id int64) (*GroupInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupInfoRows, m.table)
	var resp GroupInfo
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

func (m *defaultGroupInfoModel) SelectAll(ctx context.Context) (*[]*GroupInfo, error) {
	query := fmt.Sprintf("select %s from %s", groupInfoRows, m.table)
	resp := new([]*GroupInfo)
	err := m.conn.QueryRowsCtx(ctx, resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGroupInfoModel) SelectGroupDetail(ctx context.Context) (resp []GroupDetail, err error) {
	groupInfo, groupErr := m.SelectAll(ctx)
	if (groupErr != nil) {
		logger.ErrorLog(groupErr.Error())
		err = groupErr
		return
	}

	resp = make([]GroupDetail, 0)
	for _, item := range *groupInfo {
		var detail GroupDetail
		copier.Copy(&detail, &item)

		taskDao := groupTask.NewGroupTasksModel(m.conn)
		taskIds, taskErr := taskDao.SelectTaskIds(ctx, item.Id)
		if (taskErr != nil) {
			logger.ErrorLog(taskErr.Error())
			err = taskErr
			return
		}

		if len(taskIds) == 0 {
			continue
		}

		sshDao := sshtask.NewSshTaskModel(m.conn)
		sshTasks, sshErr := sshDao.SelectItemsByIds(ctx, taskIds)
		if (sshErr != nil) {
			logger.ErrorLog(sshErr.Error())
			err = sshErr
			return
		}
		detail.Tasks = sshTasks
		resp = append(resp, detail)
	}
	return
}

func (m *defaultGroupInfoModel) SelectGroupDetailById(ctx context.Context, groupId int64) (resp GroupDetail, err error) {
	var detail GroupDetail
	detail.Id = groupId

	taskDao := groupTask.NewGroupTasksModel(m.conn)
	taskIds, taskErr := taskDao.SelectTaskIds(ctx, groupId)
	if (taskErr != nil) {
		logger.ErrorLog(taskErr.Error())
		err = taskErr
		return
	}

	if len(taskIds) == 0 {
		return detail, nil
	}

	sshDao := sshtask.NewSshTaskModel(m.conn)
	sshTasks, sshErr := sshDao.SelectItemsByIds(ctx, taskIds)
	if (sshErr != nil) {
		logger.ErrorLog(sshErr.Error())
		err = sshErr
		return
	}
	detail.Tasks = sshTasks
	return detail, nil
}

func (m *defaultGroupInfoModel) Insert(ctx context.Context, data *GroupInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, groupInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name)
	return ret, err
}

func (m *defaultGroupInfoModel) Update(ctx context.Context, data *GroupInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, groupInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Id)
	return err
}

func (m *defaultGroupInfoModel) tableName() string {
	return m.table
}

package groupTask

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GroupTasksModel = (*customGroupTasksModel)(nil)

type (
	// GroupTasksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupTasksModel.
	GroupTasksModel interface {
		groupTasksModel
		withSession(session sqlx.Session) GroupTasksModel
	}

	customGroupTasksModel struct {
		*defaultGroupTasksModel
	}
)

// NewGroupTasksModel returns a model for the database table.
func NewGroupTasksModel(conn sqlx.SqlConn) GroupTasksModel {
	return &customGroupTasksModel{
		defaultGroupTasksModel: newGroupTasksModel(conn),
	}
}

func (m *customGroupTasksModel) withSession(session sqlx.Session) GroupTasksModel {
	return NewGroupTasksModel(sqlx.NewSqlConnFromSession(session))
}

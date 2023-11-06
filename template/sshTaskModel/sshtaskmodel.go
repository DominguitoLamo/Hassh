package template

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SshTaskModel = (*customSshTaskModel)(nil)

type (
	// SshTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSshTaskModel.
	SshTaskModel interface {
		sshTaskModel
		withSession(session sqlx.Session) SshTaskModel
	}

	customSshTaskModel struct {
		*defaultSshTaskModel
	}
)

// NewSshTaskModel returns a model for the database table.
func NewSshTaskModel(conn sqlx.SqlConn) SshTaskModel {
	return &customSshTaskModel{
		defaultSshTaskModel: newSshTaskModel(conn),
	}
}

func (m *customSshTaskModel) withSession(session sqlx.Session) SshTaskModel {
	return NewSshTaskModel(sqlx.NewSqlConnFromSession(session))
}

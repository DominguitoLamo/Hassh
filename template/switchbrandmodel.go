package template

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SwitchBrandModel = (*customSwitchBrandModel)(nil)

type (
	// SwitchBrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSwitchBrandModel.
	SwitchBrandModel interface {
		switchBrandModel
		withSession(session sqlx.Session) SwitchBrandModel
	}

	customSwitchBrandModel struct {
		*defaultSwitchBrandModel
	}
)

// NewSwitchBrandModel returns a model for the database table.
func NewSwitchBrandModel(conn sqlx.SqlConn) SwitchBrandModel {
	return &customSwitchBrandModel{
		defaultSwitchBrandModel: newSwitchBrandModel(conn),
	}
}

func (m *customSwitchBrandModel) withSession(session sqlx.Session) SwitchBrandModel {
	return NewSwitchBrandModel(sqlx.NewSqlConnFromSession(session))
}

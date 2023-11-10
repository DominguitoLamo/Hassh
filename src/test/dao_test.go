package test

import (
	"context"
	"fmt"
	groupInfo "hassh/src/internal/model/groupInfoModel"
	groupTask "hassh/src/internal/model/groupTasksModel"
	sshtask "hassh/src/internal/model/sshTaskModel"
	"testing"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CONN_URL = "root:moto9171@/hassh?parseTime=true"
)

func TestSelectGroupDetail(t *testing.T) {
	dbConnection := sqlx.NewMysql(CONN_URL)
	dao := groupInfo.NewGroupInfoModel(dbConnection)
	items, _ := dao.SelectGroupDetail(context.Background())
	fmt.Println(items)
}

func TestTaskIds(t *testing.T) {
	dbConnection := sqlx.NewMysql(CONN_URL)
	dao := groupTask.NewGroupTasksModel(dbConnection)
	items, _ := dao.SelectTaskIds(context.Background(), 1)
	fmt.Println(items)
}

func TestSshTasksByTaskIds(t *testing.T) {
	dbConnection := sqlx.NewMysql(CONN_URL)
	dao := groupTask.NewGroupTasksModel(dbConnection)
	items, _ := dao.SelectTaskIds(context.Background(), 1)
	
	sshDao := sshtask.NewSshTaskModel(dbConnection)
	result, _ := sshDao.SelectItemsByIds(context.Background(), items)
	for _, item := range *result {
		fmt.Println(item)
	}
}
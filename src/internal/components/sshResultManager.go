package components

import (
	"context"
	"fmt"
	"hassh/src/internal/types"
	"os"
	"sync"

	switchgo "github.com/DominguitoLamo/switchGo"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SshResultManager struct {
	resultMap map[string]*os.File
	mutex sync.Mutex
}

func NewSSHResultManager() *SshResultManager {
	manager := new(SshResultManager)
	manager.mutex = sync.Mutex{}
	manager.resultMap = make(map[string]*os.File)
	return manager
}

func (s *SshResultManager) SaveResult(session *switchgo.SSHSession, req *types.RunCmdReq, key, result string) error {
	// create file
	t := session.GetLastUseTime()
	date := fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day()) 
	fileName := req.Name + " " + req.Ip + " " + date
	file, fileErr := os.Create(fileName)
	if (fileErr != nil) {
		logx.Error(fileErr.Error())
		return fileErr
	}
	bytes, err := file.Write([]byte(result))
	if (err != nil) {
		logx.Error("Write file error: ", err.Error())
	} else {
		logc.Debug(context.Background(), "Write file success. %s bytes written", bytes)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.resultMap[key] = file

	return nil
}
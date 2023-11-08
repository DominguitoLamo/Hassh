package components

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hassh/src/internal/types"
	"io"
	"os"
	"sync"

	switchgo "github.com/DominguitoLamo/switchGo"
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

func (s *SshResultManager) SaveResult(req *types.RunCmdReq, session *switchgo.SSHSession, result string) (string, error) {
	// gen key
	key, hasherr := genhashKey(req, session)
	if (hasherr != nil) {
		logx.Error(hasherr.Error())
		return "", hasherr
	}

	// create file
	t := session.GetLastUseTime()
	date := fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day()) 
	fileName := req.Name + " " + req.Ip + " " + date
	file, fileErr := os.Create(fileName)
	if (fileErr != nil) {
		logx.Error(fileErr.Error())
		return "", fileErr
	}
	file.Write([]byte(result))

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.resultMap[key] = file

	return key, nil
}

func genhashKey(req *types.RunCmdReq, session *switchgo.SSHSession) (string, error) {
	hasher := md5.New()
	keyStr := fmt.Sprintf("%s%s%s", req.Name, req.Ip, session.GetLastUseTime().String())
	_, err := io.WriteString(hasher, keyStr)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil))[:8], nil
}
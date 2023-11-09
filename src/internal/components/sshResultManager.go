package components

import (
	"fmt"
	"hassh/src/internal/types"
	"sync"

	switchgo "github.com/DominguitoLamo/switchGo"
)

type CacheFile struct {
	Content []byte
	Name string
}

type SshResultManager struct {
	resultMap map[string]*CacheFile
	mutex sync.Mutex
}

func NewCacheFile(content []byte, name string) *CacheFile {
	f := new(CacheFile)
	f.Content = content
	f.Name = name
	return f
}

func NewSSHResultManager() *SshResultManager {
	manager := new(SshResultManager)
	manager.mutex = sync.Mutex{}
	manager.resultMap = make(map[string]*CacheFile)
	return manager
}

func (s *SshResultManager) SaveResult(session *switchgo.SSHSession, req *types.RunCmdReq, key, result string) error {
	// create file
	t := session.GetLastUseTime()
	date := fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day()) 
	fileName := req.Name + "_" + req.Ip + "_" + date + ".txt"

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.resultMap[key] = NewCacheFile([]byte(result), fileName)

	return nil
}

func (s *SshResultManager) IsExist(id string) bool {
	result := s.resultMap[id]
	return result != nil
}

func (s *SshResultManager) GetFile(id string) *CacheFile {
	result := s.resultMap[id]
	return result
}

func (s *SshResultManager) DeleteFile(id string) {
	s.resultMap[id] = nil
}
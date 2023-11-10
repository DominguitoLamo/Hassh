package components

import "sync"

type GroupTaskResult struct {
	Key     string
	Details []SSHTaskDetail
	ErrMsg  string
}

type SSHTaskDetail struct {
	Name    string
	Content string
}

type GroupResultManager struct {
	resultMap map[string]GroupTaskResult
	mutex     *sync.Mutex
}

func NewGroupResultManager() GroupResultManager {
	var manager GroupResultManager
	manager.resultMap = make(map[string]GroupTaskResult)
	manager.mutex = &sync.Mutex{}
	return manager
}

func (s *GroupResultManager) SaveResult(result GroupTaskResult) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.resultMap[result.Key] = result

	return nil
}

func (s *GroupResultManager) IsExist(id string) bool {
	_, ok := s.resultMap[id]
	return ok
}

func (s *GroupResultManager) GetFile(id string) GroupTaskResult {
	result := s.resultMap[id]
	return result
}

func (s *GroupResultManager) DeleteFile(id string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.resultMap, id)
}
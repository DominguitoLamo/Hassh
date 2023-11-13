package taskqueue

import (
	"hassh/src/internal/types"

	switchgo "github.com/DominguitoLamo/switchGo"
	"github.com/zeromicro/go-zero/core/logx"
)

type SshQueueRequest struct {
	Req *types.RunCmdReq
	SshManager *switchgo.SessionManager
	ExecuteTask func()
}

func NewSshQueueRequest(req *types.RunCmdReq, sshManager *switchgo.SessionManager) *SshQueueRequest {
	l := new(SshQueueRequest)
	l.Req = req
	l.SshManager = sshManager
	return l
}

type SshTaskQueue struct {
	queue chan *SshQueueRequest
}

func NewSshTaskQueue() *SshTaskQueue {
	s := new(SshTaskQueue)
	s.queue = make(chan *SshQueueRequest)
	return s
}

func (s *SshTaskQueue) RunTask() {
	go func() {
		for {
			r := <- s.queue
			logx.Debug("ssh task request: ", r)
			r.ExecuteTask()
		}
	}()
}

func (s *SshTaskQueue) AddTask(r *SshQueueRequest) {
	s.queue <- r
	logx.Debug("add finish:", r)
}
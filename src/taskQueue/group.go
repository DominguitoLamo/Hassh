package taskqueue

type GroupExecuteTask func()

type GroupTaskQueue struct {
	queue chan GroupExecuteTask
}

func NewGroupTaskQueue() *GroupTaskQueue {
	s := new(GroupTaskQueue)
	s.queue = make(chan GroupExecuteTask)
	return s
}

func (s *GroupTaskQueue) RunTask() {
	go func() {
		for {
			r := <- s.queue
			r()
		}
	}()
}

func (s *GroupTaskQueue) AddTask(r GroupExecuteTask) {
	s.queue <- r
}
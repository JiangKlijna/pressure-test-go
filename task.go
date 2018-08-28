package main

type TaskServer struct {
	person  []TaskPerson
	setting *TaskSetting
}

// start server
func (t *TaskServer) start() {

}

// stop server
func (t *TaskServer) stop() {
	for _, p := range t.person {
		p.stop()
	}
}

// real statistics
func (t *TaskServer) real_statistics() {

}

// notify statistics
func (t *TaskServer) notify_statistics() {
	for _, p := range t.person {
		if !p.isStop {
			return
		}
	}
	t.real_statistics()
}

func NewTaskServer(setting *TaskSetting) *TaskServer {
	return &TaskServer{make([]TaskPerson, setting.Final_person), setting}
}

type TaskPerson struct {
	isRun  bool
	isStop bool
	task   *TaskServer
}

// single request
func (t *TaskPerson) run() {

}

// start multi-request
func (t *TaskPerson) start() {
	go func() {
		t.isRun = true
		for t.isRun {
			t.run()
		}
		t.isStop = true
	}()
}

// stop multi-request
func (t *TaskPerson) stop() {
	t.isRun = false
}

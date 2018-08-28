package main

type TaskServer struct {
	person  []TaskPerson
	setting *TaskSetting
}

func (t *TaskServer) start() {

}

func (t *TaskServer) stop() {
	for _, p := range t.person {
		p.stop()
	}
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

// start server
func (t *TaskPerson) start() {
	go func() {
		t.isRun = true
		for t.isRun {
			t.run()
		}
		t.isStop = true
	}()
}

// stop server
func (t *TaskPerson) stop() {
	t.isRun = false
}

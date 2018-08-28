package main

type TaskServer struct {
	person  []TaskPerson
	setting *TaskSetting
}

func NewTaskServer(setting *TaskSetting) *TaskServer {
	return &TaskServer{make([]TaskPerson, setting.Final_person), setting}
}

type TaskPerson struct {
	isRun bool
	urls  []map[string]interface{}
}

// single request
func (t *TaskPerson) run() {

}

// start server
func (t *TaskPerson) start() {
	go func() {
		for t.isRun {
			t.run()
		}
	}()
}

// stop server
func (t *TaskPerson) stop() {
	t.isRun = false
}

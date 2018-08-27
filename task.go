package main

type TaskServer struct {
	isRun   bool
	setting *TaskSetting
}

func NewTaskServer(setting *TaskSetting) *TaskServer {
	return &TaskServer{true, setting}
}

// single request
func (t *TaskServer) run() {

}

// start server
func (t *TaskServer) start() {
	go func() {
		for t.isRun {
			t.run()
		}
	}()
}

// stop server
func (t *TaskServer) stop() {
	t.isRun = false
}

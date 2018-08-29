package main

import "sync"

type TaskServer struct {
	mutex   sync.Mutex
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
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for _, p := range t.person {
		if !p.isStop {
			return
		}
	}
	t.real_statistics()
}

func NewTaskServer(setting *TaskSetting) *TaskServer {
	return &TaskServer{sync.Mutex{}, make([]TaskPerson, setting.Final_person), setting}
}

type TaskPerson struct {
	isRun  bool
	isStop bool
	task   *TaskServer
	result *PressureTestResult
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
		t.task.notify_statistics()
	}()
}

// stop multi-request
func (t *TaskPerson) stop() {
	t.isRun = false
}

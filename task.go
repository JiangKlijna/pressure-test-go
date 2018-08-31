package main

import (
	"sync"
	"time"
)

type TaskServer struct {
	mutex   sync.Mutex
	person  []TaskPerson
	setting *TaskSetting
}

// start server
func (t *TaskServer) start() {
	for i := t.setting.Init_person; i <= t.setting.Final_person; i += t.setting.Add_person {
		for j := i; j <= t.setting.Add_person; j++ {
			t.person[j].start()
		}
		time.Sleep(time.Duration(t.setting.Duration_time * 1000))
	}
	t.stop()
}

// stop server
func (t *TaskServer) stop() {
	for _, p := range t.person {
		p.stop()
	}
}

// real statistics
func (t *TaskServer) real_statistics() {
	result := &PressureTestResult{}
	for _, p := range t.person {
		result.add(p.result);
	}
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

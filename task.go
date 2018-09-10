package main

import (
	"sync"
	"time"
	"net/http"
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
	client *http.Client
	result *PressureTestResult
}

// single request
func (t *TaskPerson) run() {
	url := t.task.setting.random_url()
	req, err := http.NewRequest(url.method(), url.url(), url.data())
	if err != nil {
		t.stop()
		return
	}
	for k, v := range t.task.setting.Headers {
		req.Header.Add(k, v)
	}
	start := time.Now()
	res, err := t.client.Do(req)
	end := time.Since(start)
	if err != nil {
		t.stop()
		return
	}
	println(res, end)
}

// start multi-request
func (t *TaskPerson) start() {
	go func() {
		t.client = &http.Client{}
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

// mark PressureTestResult
func (t *TaskPerson) mark(isFailure bool, start time.Time) {
	t.result.request_number++
	if isFailure {
		t.result.failure_number++
	}
	t.result.duration_time += time.Since(start)
}

package main

import (
	"log"
	"sync"
	"time"
	"net/http"
)

type TaskService struct {
	tag     string
	mutex   sync.Mutex
	person  []TaskPerson
	setting *TaskSetting
}

// start server
func (t *TaskService) start() {
	i := 0
	setting := t.setting
	// init start server
	for ; i < setting.Init_person; i++ {
		t.person[i].start(i, t)
	}
	for i < setting.Final_person {
		time.Sleep(time.Duration(setting.Duration_time) * time.Second)
		for j := 0; j < setting.Add_person; j++ {
			log.Println(i)
			t.person[i].start(i, t)
			i++
		}
	}
	t.stop()
}

// stop server
func (t *TaskService) stop() {
	for _, p := range t.person {
		p.stop()
	}
}

// real statistics
func (t *TaskService) real_statistics() {
	result := &PressureTestResult{}
	for _, p := range t.person {
		result.add(p.result);
	}
}

// notify statistics
func (t *TaskService) notify_statistics() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for _, p := range t.person {
		if !p.isStop {
			return
		}
	}
	t.real_statistics()
}

func NewTaskService(tag string, setting *TaskSetting) *TaskService {
	return &TaskService{tag, sync.Mutex{}, make([]TaskPerson, setting.Final_person), setting}
}

type TaskPerson struct {
	index  int
	isRun  bool
	isStop bool
	task   *TaskService
	client *http.Client
	result *PressureTestResult
}

// single request
func (t *TaskPerson) run() {
	start := time.Now()
	url := t.task.setting.random_url()
	req, err := http.NewRequest(url.method(), url.url(), url.data())
	log.Printf("TaskService[%s][%d]->%s->%s\n", t.task.tag, t.index, url.method(), url.url())
	if err != nil {
		t.stop()
		t.result.mark(false, start)
		return
	}
	for k, v := range t.task.setting.Headers {
		req.Header.Add(k, v)
	}
	res, err := t.client.Do(req)
	if err != nil {
		t.stop()
		t.result.mark(false, start)
		return
	}
	if res.StatusCode >= 200 || res.StatusCode < 300 {
		t.result.mark(true, start)
	} else {
		t.result.mark(false, start)
	}
}

// start multi-request and init
func (t *TaskPerson) start(index int, service *TaskService) {
	go func() {
		t.index = index
		t.task = service
		t.client = &http.Client{}
		t.result = &PressureTestResult{}
		t.isRun = true
		log.Printf("TaskService[%s][%d]->start\n", service.tag, index)
		for t.isRun {
			t.run()
		}
		log.Printf("TaskService[%s][%d]->stop\n", service.tag, index)
		t.isStop = true
		t.task.notify_statistics()
	}()
}

// stop multi-request
func (t *TaskPerson) stop() {
	t.isRun = false
}

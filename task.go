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
	person  []SubTask
	setting *TaskSetting
}

// start server
func (t *TaskService) start() {
	i := 0
	setting := t.setting
	// init start server
	for ; i < setting.Init_person; i++ {
		go t.person[i].start()
	}
	for i < setting.Final_person {
		time.Sleep(time.Duration(setting.Duration_time) * time.Second)
		for j := 0; j < setting.Add_person; j++ {
			log.Println(i)
			go t.person[i].start()
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
		result.add(p.result());
	}
}

// notify statistics
func (t *TaskService) notify_statistics() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for _, p := range t.person {
		if !p.isStop() {
			return
		}
	}
	t.real_statistics()
}

// single request
func (t *TaskService) request(url *Url, client *http.Client) bool {
	req, err := http.NewRequest(url.method(), url.url(), url.data())
	if err != nil {
		return false
	}
	for k, v := range t.setting.Headers {
		req.Header.Add(k, v)
	}
	res, err := client.Do(req)
	if err != nil {
		return false
	}
	return res.StatusCode >= 200 || res.StatusCode < 300
}

func NewTaskService(tag string, setting *TaskSetting) *TaskService {
	return &TaskService{tag, sync.Mutex{}, make([]SubTask, setting.Final_person), setting}
}

type SubTask struct {
	start  func()
	stop   func()
	isRun  func() bool
	isStop func() bool
	result func() *PressureTestResult
}

func NewSubTask(index int, task *TaskService) *SubTask {
	isRun, isStop := false, false
	client := &http.Client{}
	result := &PressureTestResult{}
	return &SubTask{
		start: func() {
			isRun = true
			isStop = false
			log.Printf("TaskService[%s][%d]->start\n", task.tag, index)
			for isRun {
				url := task.setting.random_url()
				task.request(&url, client)
				log.Printf("TaskService[%s][%d]->%s->%s\n", task.tag, index, url.method(), url.url())
			}
			log.Printf("TaskService[%s][%d]->stop\n", task.tag, index)
			isStop = true
			task.notify_statistics()
		},
		stop:   func() { isRun = false },
		isRun:  func() bool { return isRun },
		isStop: func() bool { return isStop },
		result: func() *PressureTestResult { return result },
	}
}

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
	persons []*SubTask
	setting *TaskSetting
}

// start server
func (t *TaskService) start() {
	i := 0
	setting := t.setting
	// init start server
	for ; i < setting.Init_person; i++ {
		go t.persons[i].start()
	}
	for i < setting.Final_person {
		time.Sleep(time.Duration(setting.Duration_time) * time.Second)
		for j := 0; j < setting.Add_person && i < setting.Final_person; j++ {
			go t.persons[i].start()
			i++
		}
	}
	time.Sleep(time.Duration(setting.Duration_time) * time.Second)
	t.stop()
}

// stop server
func (t *TaskService) stop() {
	for _, p := range t.persons {
		p.stop()
	}
}

// real statistics
func (t *TaskService) real_statistics() {
	res := make([]*PressureTestResult, len(t.persons) + 1)
	result := &PressureTestResult{}
	for i, p := range t.persons {
		result.add(p.result());
		res[i] = p.result()
	}
	res[len(t.persons)] = result
	OutputResult(res, t.setting.Output_format)
}

// notify statistics
func (t *TaskService) notify_statistics() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for _, p := range t.persons {
		if !p.isStop() {
			return
		}
	}
	t.real_statistics()
}

// single request and return isFailure
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
	return res.StatusCode < 200 || res.StatusCode >= 300
}

func NewTaskService(tag string, setting *TaskSetting) *TaskService {
	service := &TaskService{tag, sync.Mutex{}, make([]*SubTask, setting.Final_person), setting}
	for i := 0; i < setting.Final_person; i++ {
		service.persons[i] = NewSubTask(i, service)
	}
	return service
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
	result := &PressureTestResult{Id:index+1}
	return &SubTask{
		start: func() {
			isRun = true
			isStop = false
			log.Printf("TaskService[%s][%d]->start\n", task.tag, index)
			for isRun {
				url := task.setting.random_url()
				start := time.Now()
				isFailure := task.request(&url, client)
				result.mark(isFailure, start)
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

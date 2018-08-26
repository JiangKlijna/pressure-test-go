package main

import (
	"io/ioutil"
	"os"
	"encoding/json"
)

const defaultJson = `{
	"task0": {
		"urls": [
			{"method": "GET", "url": "http://baidu.com", "params": {}},
			{"method": "POST", "url": "http://baidu.com", "data": {}}
		],
		"headers": {"user-agent": "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko"},
		"init_person": 10,
		"add_persion": 10,
		"final_person": 50,
		"duration_time": 5,
		"output_format": "csv"
	},
    "task1": {
		"urls": [
			{"method": "GET", "url": "http://baidu.com", "params": {}},
			{"method": "POST", "url": "http://baidu.com", "data": {}}
		],
		"headers": {"user-agent": "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko"},
		"init_person": 10,
		"add_persion": 10,
		"final_person": 50,
		"duration_time": 5,
		"output_format": "html"
    }
}`

type TaskSetting struct {
	Init_person   int                      `json:"init_person"`
	Add_persion   int                      `json:"add_persion"`
	Final_person  int                      `json:"final_person"`
	Duration_time int                      `json:"duration_time"`
	Output_format string                   `json:"output_format"`
	Urls          []map[string]interface{} `json:"urls"`
	Headers       map[string]string        `json:"headers"`
}

// New creates a new Setting
func NewTaskSetting(filename string) (map[string]TaskSetting, error) {
	var bytes []byte
	var err error
	if !FileExists(filename) {
		bytes = []byte(defaultJson)
		ioutil.WriteFile(filename, bytes, os.ModePerm)
	} else {
		bytes, err = ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
	}
	data := make(map[string]TaskSetting)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

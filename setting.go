package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"math/rand"
	"strings"
)

const defaultJson = `{
	"task0": {
		"urls": [
			{"method": "GET", "url": "http://baidu.com", "params": {}},
			{"method": "POST", "url": "http://baidu.com", "data": {}}
		],
		"headers": {"user-agent": "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko"},
		"init_person": 10,
		"add_person": 10,
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
		"add_person": 10,
		"final_person": 50,
		"duration_time": 5,
		"output_format": "html"
    }
}`

var empty_body = strings.NewReader("")

type Url map[string]interface{}

type TaskSetting struct {
	Init_person   int               `json:"init_person"`
	Add_person    int               `json:"add_person"`
	Final_person  int               `json:"final_person"`
	Duration_time int               `json:"duration_time"`
	Output_format string            `json:"output_format"`
	Urls          []Url             `json:"urls"`
	Headers       map[string]string `json:"headers"`
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

// get random url
func (t TaskSetting) random_url() Url {
	return t.Urls[rand.Intn(len(t.Urls))]
}

// get method
func (u Url) method() string {
	if method, is := u["method"]; is {
		if method, is := method.(string); is {
			return method
		}
	}
	return "GET"
}

// get url path
func (u Url) url() string {
	if path, is := u["path"]; is {
		return path.(string)
	}
	if params, is := u["params"]; is {
		url := u["url"].(string)
		s := params_string(params.(map[string]interface{}))
		if s == "" {
			u["path"] = url
		} else if strings.LastIndex(url, "?") > 0 {
			u["path"] = url + "&" + s
		} else {
			u["path"] = url + "?" + s
		}
	} else {
		u["path"] = u["url"]
	}
	return u.url()
}

// get data
func (u Url) data() *strings.Reader {
	if dat, is := u["dat"]; is {
		if dat == "" {
			return empty_body
		} else {
			return strings.NewReader(dat.(string))
		}
	}
	if data, is := u["data"]; is {
		u["dat"] = params_string(data.(map[string]interface{}))
		return u.data()
	} else {
		return empty_body
	}
}

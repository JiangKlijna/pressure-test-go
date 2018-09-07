package main

import "testing"

func TestNewSetting(t *testing.T) {
	setting, err := NewTaskSetting("setting.json")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(setting)
	for _, s := range setting {
		for _, url := range s.Urls {
			t.Log(url)
		}
	}

	u := Url(map[string]interface{}{
		//"method": "GET",
		"url":    "http://baidu.com/s",
		"params": map[string]string{
			"wd": "abc",
		},
		"data": map[string]string{
			"wd": "abc",
		},
	})

	if u.url() != u["path"] {
		t.Error(u.url())
		return
	}
	t.Log(u.method())
	t.Log(u.url())
	t.Log(u.data())
}


package main

import "testing"

func TestNewSetting(t *testing.T) {
	setting, err := NewTaskSetting("setting.json")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(setting)
}

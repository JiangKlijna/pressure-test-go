package main

import (
	"os"
	"os/signal"
)

type Application struct {
	services map[string]TaskService
	settings map[string]TaskSetting
}

// Init log & set
func (app *Application) Init() error {
	ss, err := NewTaskSetting("setting.json")
	if err != nil {
		return err;
	}
	app.settings = ss
	app.services = make(map[string]TaskService)
	for k, s := range ss {
		err = s.is_valid()
		if err != nil {
			return err
		}
		app.services[k] = *NewTaskService(k, &s)
	}
	return nil
}

// check error and exit
func (app *Application) check(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

// Start all of server
func (app *Application) Start() {
	for _, s := range app.services {
		s.start()
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

// Stop all of server
func (app *Application) Stop() {
	for _, s := range app.services {
		s.stop()
	}
}

func main() {
	app := &Application{}
	err := app.Init()
	app.check(err)
	app.Start()
	app.Stop()
}

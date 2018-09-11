package main

import(
	"os"
)

type Application struct {
	settings *map[string]TaskSetting
}

// Init log & set
func (app *Application) Init() error {
	ss, err := NewTaskSetting("setting.json")
	if err != nil {
		return err;
	}
	app.settings = &ss
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

}

// Stop all of server
func (app *Application) Stop() {
}

func main() {
	app := &Application{}
	err := app.Init()
	app.check(err)
	app.Start()
	app.Stop()
}

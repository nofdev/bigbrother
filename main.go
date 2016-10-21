package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func getExampleConf() (conf string) {
	conf = `[DEFAULT]
configration_path = /etc/bigbrother/bigbrother.conf

[monitor]
check_url = /status
check_interval = 3s
receive_string = Running`
	return
}

var (
	app                 = kingpin.New("bigbrother", "Bit brother is watching you!")
	daemon              = app.Command("daemon", "running as daemon")
	show                = app.Command("show", "show bigbrother info")
	showExampleConf     = show.Command("example-conf", "show example configration")
	showExampleConfSave = showExampleConf.Flag("save", "save configuration to /etc/bigbrother/bigbrother.conf").Bool()
)

func main() {
	app.Version("v0.0.1")
	app.Author("nofdev")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case "daemon":
		// TODO running as daemon
	case "show example-conf":
		fmt.Println(getExampleConf())
		if *showExampleConfSave {
			// TODO save file
		}
	}

}

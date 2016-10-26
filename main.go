package main

import (
	"fmt"
	"os"

	"github.com/nofdev/bigbrother/conf"
	"gopkg.in/alecthomas/kingpin.v2"
)

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
		fmt.Println(conf.GetExampleConf())
		if *showExampleConfSave {
			// TODO save file
		}
	}

}

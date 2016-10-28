package conf

import (
	"log"

	"github.com/alyu/configparser"
)

// ReadConf reads configration from /etc/bigbrother/bigbrother.conf and returns the full configuration
func ReadConf(path string) (config *configparser.Configuration, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
            log.Fatal("Cannot read bigbrother.conf in /etc/bigbrother")
		}
	}()

	config, err = configparser.Read(path)
	if err != nil {
		panic(err)
	}
    return config, err
}

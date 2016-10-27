package conf

import "io/ioutil"

// GetExampleConf returns example configration.
func GetExampleConf() (conf string) {
	conf = `[DEFAULT]
configration_path = /etc/bigbrother/bigbrother.conf

[monitor]
check_url = /status
check_interval = 3s
receive_string = Running`

	return
}

// SaveConf save string to a file.
func SaveConf(s string, path string) error {
	err := ioutil.WriteFile(path, []byte(s), 0644)
	if err != nil {
		panic(err)
	} else {
		return nil
	}
}

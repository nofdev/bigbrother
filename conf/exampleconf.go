package conf

// GetExampleConf returns example configration
func GetExampleConf() (conf string) {
	conf = `[DEFAULT]
configration_path = /etc/bigbrother/bigbrother.conf

[monitor]
check_url = /status
check_interval = 3s
receive_string = Running`

	return
}

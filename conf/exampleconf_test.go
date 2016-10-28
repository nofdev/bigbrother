package conf

import "testing"

func TestSaveConf(t *testing.T) {
	conf := GetExampleConf()
	err := SaveConf(conf, "bigbrother.conf")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("save conf passed")
	}
}

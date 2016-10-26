package conf

import "testing"
import "os"

func TestSaveConf(t *testing.T) {
	conf := GetExampleConf()
	err := SaveConf(conf, "bigbrother.conf")
	if err != nil {
		t.Error(err)
	} else {

		if err := os.Remove("bigbrother.conf"); err != nil {
			t.Error(err)
		}
		t.Log("save conf passed and cleaned")
	}
}

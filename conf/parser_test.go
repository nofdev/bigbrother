package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConf(t *testing.T) {
	config, err := ReadConf("bigbrother.conf")
	if err != nil {
		t.Error(err)
	}
	section, err := config.Section("DEFAULT")
	if err != nil {
		t.Error(err)
	}
	path := section.ValueOf("configration_path")
	assert.Equal(t, "/etc/bigbrother/bigbrother.conf", path)
}

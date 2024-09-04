package conf_test

import (
	"testing"
	"vblog/conf"
)

func TestLoadConfigFromToml(t *testing.T) {

	err := conf.LoadConfigFromToml("../etc/config.toml")
	if err != nil {
		t.Errorf("want nil, but got %v", err)
	}

	t.Log(conf.C())

}

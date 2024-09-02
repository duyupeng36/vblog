package conf_test

import (
	"testing"
	"vblog/conf"

	"github.com/BurntSushi/toml"
)

func TestTomlDecodeFile(t *testing.T) {
	configObj := &conf.Config{}

	_, err := toml.DecodeFile("../etc/config.toml", configObj)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(configObj)
	t.Log(configObj.MySQL)
}

func init() {
	err := conf.LoadConfigFromToml("../etc/config.toml")
	if err != nil {
		panic(err)
	}
}

func TestMySQLGetConnPool(t *testing.T) {
	db, err := conf.C().MySQL.GetConnPool()
	if err != nil {
		t.Errorf("want error to be nil, but get %v\n", err.Error())
	}

	t.Log(db)
}

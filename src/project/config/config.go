package config

import (
	"encoding/json"
	"io/ioutil"
)

type CommandStruct struct {
	Cmd string `json:cmd`
}

type Config struct {
	Dir    string        `json:dir`
	Log    string        `json:log`
	Create CommandStruct `json:create`
	Change CommandStruct `json:change`
	Delete CommandStruct `json:delete`
}

var StaticConfig = Config{Dir: "./"}

func GetConfig(configFile string) Config {

	StaticConfig = Config{Dir: "./"}

	if configFile != "" {
		data, _ := ioutil.ReadFile(configFile)
		json.Unmarshal([]byte(data), &StaticConfig)
	}

	return StaticConfig
}

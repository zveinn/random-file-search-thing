package files

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var GlobalConfig *Config

func LoadConfig() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println("Could not find config file ...")
		os.Exit(1)
	}
	GlobalConfig = new(Config)
	err = json.Unmarshal(data, GlobalConfig)
	if err != nil {
		log.Println("Could not read/parse the config file ...")
		os.Exit(1)
	}

}

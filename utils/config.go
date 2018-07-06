package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type Config struct {
	App struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"app"`

	Db struct {
		Host string `json:"host"`
		Port string `json:"port"`
		Name string `json:"name"`
		User string `json:"user"`
		Pass string `json:"pass"`
	} `json:"db"`

	Redis struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"redis"`

	OAuth struct {
		URL string `json:"url"`
	}
}

func GetConfig() Config {
	cfg := Config{}

	cfgDir := os.Getenv("CONFIGDIR")

	if cfgDir == "" {
		cfgDir = "./src/configs"
	}

	env := os.Getenv("ENV")

	if env == "" {
		env = gin.DebugMode
	}

	cfgFile := path.Join(cfgDir, "config."+env+".json")

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		log.Fatal("no config file", err)
	}

	if content, err := ioutil.ReadFile(cfgFile); err == nil {
		if err := json.Unmarshal(content, &cfg); err != nil {
			log.Fatal("no config file", err)
		}
	}

	return cfg
}

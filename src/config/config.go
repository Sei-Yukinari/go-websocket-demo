package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App   App   `toml:"app"`
	Redis Redis `toml:"redis"`
}

type App struct {
	Port string `toml:"port"`
}

type Redis struct {
	URL      string `toml:"url"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

var Conf *Config

const repositoryName = "go-websocket-demo"

func getProjectRootPath() string {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, repositoryName) && !strings.HasSuffix(wd, "app") {
		wd = filepath.Dir(wd)
	}

	return wd
}

func init() {
	Conf = new(Config)

	GoEnv := os.Getenv("GO_ENV")
	if GoEnv == "" {
		GoEnv = "development"
	}

	viper.SetConfigName(GoEnv)
	viper.AddConfigPath(getProjectRootPath() + "/src/config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s \n", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("failed to unmarshal err: %s \n", err))
	}
}

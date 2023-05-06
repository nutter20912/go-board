package config

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Env  string `yaml:"env"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"app"`

	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"password"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"db"`
}

// 載入配置
func Init() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	files, err := filepath.Glob("./config/*.yaml")
	if err != nil {
		log.Fatal(err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, file := range files {
		viper.SetConfigFile(file)
		if err := viper.MergeInConfig(); err != nil {
			log.Fatal(err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return config
}

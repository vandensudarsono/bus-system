package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// type Config struct {
// 	Server struct {
// 		Host string `json:"host"`
// 		Port string `json:"port"`
// 	} `json:"server"`
// 	MongoDB struct {
// 		User     string `json:"user"`
// 		Password string `json:"password"`
// 		Hostname string `json:"hostname"`
// 		Port     int    `json:"port"`
// 		Database string `json:"database"`
// 	} `json:"mongoDB"`
// 	Requester struct {
// 		Urls              []string `json:"urls"`
// 		TimeWindowDefault string   `json:"timeWindow"`
// 		TimeWindowEachUrl []string `json:"timeWindowEachUrl"`
// 	} `json:"requester"`
// }

// var Cfg Config

// func GetConfig() *Config {
// 	return &Cfg
// }

// Initilize viper
func LoadConfig(path string) {
	readConfig(path)
}

func readConfig(path string) {
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	if path == "" {
		path = "."

	}

	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("reading config error ", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config update, automatically reloadding services..")
	})
}

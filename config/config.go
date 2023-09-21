package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

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

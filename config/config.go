package config

import (
	"github.com/spf13/viper"
	"log"
)

var defaults = map[string]string{
	"port":            "8080",
	"service_name":    "todo-api",
	"environment":     "QA",
	"project_id":      "reddit-clone-599b7",
	"collection_name": "todo",
}

func init() {
	log.Println("Initializing configs")
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
}

func InitializeConfig() {
	viper.SetEnvPrefix("TD")
	viper.AutomaticEnv()
}

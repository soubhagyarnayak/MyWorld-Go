package main

import (
	"log"

	"github.com/soubhagyarnayak/MyWorld-Go/internal"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	service := internal.Service{}
	service.Run()
}

func initConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %v", err)
	}
}

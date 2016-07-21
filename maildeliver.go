package main

import (
	"log"
	"maildeliver/api"
	"maildeliver/service"
	"maildeliver/utils"
)

const (
	APP_VERSION   = "0.1"
	CONFIG_FOLDER = "./config"
)

func main() {
	log.Println("welcome use Emailer")
	log.Printf("Application version %s", APP_VERSION)
	utils.LoadConfig(CONFIG_FOLDER)
	//store.Init()
	service.InitService()
	api.InitApi()
}

package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	ServiceConfigFileName = "service.json"
)

type ServiceSettings struct {
	EmailHost string `json:"email_host"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FromEmail string `json:"from_email"`
	TicketMax int    `json:"ticket_max"`
}

type Config struct {
	*ServiceSettings
}

var Cfg *Config

func LoadConfig(configFolderPath string) {
	if _, err := os.Stat(configFolderPath); err != nil {
		panic("Config folder not exists " + err.Error())
	}
	serviceConfigFilePath := filepath.Join(configFolderPath, ServiceConfigFileName)
	if _, err := os.Stat(serviceConfigFilePath); err != nil {
		panic("Service config file " + serviceConfigFilePath + " does not exist")
	}
	serviceConfigFile, err := os.Open(serviceConfigFilePath)
	if err != nil {
		panic("error opening file " + serviceConfigFilePath + ", error is " + err.Error())
	}
	defer serviceConfigFile.Close()

	Cfg = &Config{}
	var serviceSettings ServiceSettings
	err = json.NewDecoder(serviceConfigFile).Decode(&serviceSettings)
	if err != nil {
		panic("error parsing config file " + serviceConfigFilePath + " error is " + err.Error())
	}
	Cfg.ServiceSettings = &serviceSettings
}

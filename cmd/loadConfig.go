package cmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
) 

var loaded = false

func LoadViperConfig() (viper.Viper){
	if loaded {
		log.Debug("env already loaded, just return")
		return *viper.GetViper()
	}
	dir, _ := os.Getwd()	
	log.Info("Read .env file in", "dir", dir)
	log.Debug("LoadViperConfig\n")
	viper.SetConfigFile(".env")	
	if err := viper.ReadInConfig(); err != nil {
		log.Warn("Error reading config file, %s", err)
	}
	siteId := viper.Get("LP_SITE")
	log.Info("Load LP Site id:", "id", siteId)
	loaded = true

	return *viper.GetViper()
}
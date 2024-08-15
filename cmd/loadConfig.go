package cmd

import (
	"github.com/spf13/viper"
	"fmt"
) 

func LoadViperConfig() (viper.Viper){
	fmt.Printf("LoadViperConfig\n")
	viper.SetConfigFile(".env")	
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	siteId := viper.Get("LP_SITE")
	fmt.Println("siteid directory:", siteId)

	return *viper.GetViper()
}
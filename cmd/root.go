/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"
	"os"
	"github.com/charmbracelet/log"
	// "github.com/wlai-lp/bo-botflow/internal/lpbot"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var name, input, account, bearer, debug string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bo-botflow",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:

	// Run: func(cmd *cobra.Command, args []string) { 
	// 	fmt.Printf("hello %s\n", name)
	// 	// viper.AutomaticEnv()

	// 	viper.SetConfigFile(".env")	
	// 	// Read the .env file
	// 	if err := viper.ReadInConfig(); err != nil {
	// 		fmt.Printf("Error reading config file, %s", err)
	// 	}
	// 	// Bind flags to Viper
	// 	if err := viper.BindPFlags(cmd.Flags()); err != nil {
	// 		log.Error("Error binding flags: %v", err)
	// 	}
	// 	home := viper.Get("HOME")
	// 	siteId := viper.Get("LP_SITE")
	// 	nvName := viper.Get("name")
	// 	name := viper.Get("name")
    // 	fmt.Println("Home directory:", home)
    // 	fmt.Println("siteid directory:", siteId)
    // 	fmt.Println("name is directory:", name)
    // 	fmt.Println("nvname is directory:", nvName)
	// 	// lpbot.Hello()
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	log.SetReportCaller(true)
	// log.WithPrefix("root").Info("init")
	log.Debug("init")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	LoadViperConfig()
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bo-botflow.yaml)")

	// Cobra also supports local flags, which will only run
	// // when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")
	// // rootCmd.Flags().StringVarP(&debug, "debug", "d", "Debug", "Enable debug log level")
	// rootCmd.Flags().StringVarP(&input, "input", "i", "", "input bot json file")
	// rootCmd.Flags().StringVarP(&bearer, "bearer", "b", "", "bearer token")
	// rootCmd.Flags().StringVarP(&account, "account", "a", "", "LP Account Id / Site ID")
	// rootCmd.MarkFlagRequired("input")
	// rootCmd.Flags().StringVarP(&input, "input", "i", "World", "Name to greet")
}



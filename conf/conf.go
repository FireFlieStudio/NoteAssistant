package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func init() {
	InitConfig()
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(workDir + "/conf")
	viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// config file not found
			fmt.Println("No config file found,default config will be used")
		} else {
			// file found,but error
			panic(fmt.Errorf("Fatal error conf file: %s \n", err))
		}
	}
}

func OnWatchConfigFile() {
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}

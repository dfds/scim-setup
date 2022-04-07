package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

// ReadConfigFile reads environment variables from a file
func ReadConfigFile(configFile string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	viper.AddConfigPath(homeDir)
	viper.AddConfigPath(workDir)
	viper.SetConfigName(configFile)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Configuration file %s wasn't found in neither %s or %s\n", configFile, homeDir, workDir)
		} else {
			log.Fatalf("Unknown error when reading configuration file: %v", err.Error())
		}
	}

	s := viper.AllKeys()
	for _, v := range s {
		err := os.Setenv(strings.ToUpper(v), fmt.Sprint(viper.Get(v)))
		if err != nil {
			return
		}
	}
}

package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/dfds/scim-setup/pkg/logging"
	"github.com/spf13/viper"
)

// ReadConfigFile reads environment variables from a file
func ReadConfigFile(configFile string) (bool, error) {
	log := logging.GetLogger()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	workDir, err := os.Getwd()
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	viper.AddConfigPath(homeDir)
	viper.AddConfigPath(workDir)
	viper.SetConfigName(configFile)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Errorf("Configuration file %s wasn't found in neither %s or %s", configFile, homeDir, workDir)
			return false, err
		} else {
			log.Errorf("Unknown error when reading configuration file: %v", err.Error())
			return false, err
		}
	}

	s := viper.AllKeys()
	for _, v := range s {
		err := os.Setenv(strings.ToUpper(v), fmt.Sprint(viper.Get(v)))
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

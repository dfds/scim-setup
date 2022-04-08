package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"strings"
)

// ReadConfigFile reads environment variables from a file
func ReadConfigFile(configFile string) (bool, error) {
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
		}
	}(logger)
	sugar := logger.Sugar()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		sugar.Error(err.Error())
		return false, err
	}
	workDir, err := os.Getwd()
	if err != nil {
		sugar.Error(err.Error())
		return false, err
	}
	viper.AddConfigPath(homeDir)
	viper.AddConfigPath(workDir)
	viper.SetConfigName(configFile)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			sugar.Errorf("Configuration file %s wasn't found in neither %s or %s", configFile, homeDir, workDir)
			return false, err
		} else {
			sugar.Errorf("Unknown error when reading configuration file: %v", err.Error())
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

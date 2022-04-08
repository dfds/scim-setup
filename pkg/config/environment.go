package config

import (
	"go.uber.org/zap"
	"os"
)

// envVar checks if environment variable has been defined. If it has, its value gets returned.
func envVar(variable string) string {
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
		}
	}(logger)
	sugar := logger.Sugar()
	value, exist := os.LookupEnv(variable)
	if !exist {
		sugar.Errorf("You need to export environment variable %s", variable)
	}
	return value
}

// TenantId returns a tenant id from the environment variable AZURE_TENANT_ID
func TenantId() string {
	return envVar("AZURE_TENANT_ID")
}

// ClientId returns a client id from the environment variable AZURE_CLIENT_ID
func ClientId() string {
	return envVar("AZURE_CLIENT_ID")
}

// ClientSecret returns a client secret from the environment variable AZURE_CLIENT_SECRET
func ClientSecret() string {
	return envVar("AZURE_CLIENT_SECRET")
}

// GroupId returns a group's ObjectID
func GroupId() string {
	return envVar("AZURE_GROUP_OBJECT_ID")
}

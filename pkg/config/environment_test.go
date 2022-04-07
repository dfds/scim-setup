package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	TestId string = "00000-00000-00000-00000"
)

func TestEnvVar(t *testing.T) {
	a := assert.New(t)
	err := os.Setenv("DUMMY_VAR", "dummy")
	if err != nil {
		return
	}
	a.Equal("dummy", envVar("DUMMY_VAR"))
}

func TestGroupId(t *testing.T) {
	a := assert.New(t)
	err := os.Setenv("AZURE_GROUP_OBJECT_ID", TestId)
	if err != nil {
		return
	}
	a.Equal(TestId, GroupId())
}

func TestClientId(t *testing.T) {
	a := assert.New(t)
	err := os.Setenv("AZURE_CLIENT_ID", TestId)
	if err != nil {
		return
	}
	a.Equal(TestId, ClientId())
}

func TestClientSecret(t *testing.T) {
	a := assert.New(t)
	err := os.Setenv("AZURE_CLIENT_SECRET", TestId)
	if err != nil {
		return
	}
	a.Equal(TestId, ClientSecret())
}

func TestTenantId(t *testing.T) {
	a := assert.New(t)
	err := os.Setenv("AZURE_TENANT_ID", TestId)
	if err != nil {
		return
	}
	a.Equal(TestId, TenantId())
}

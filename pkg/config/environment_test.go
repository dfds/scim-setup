package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	TestId string = "00000-00000-00000-00000"
)

func init() {
	ReadConfigFile(".scim-setup")
}

func TestEnvVar(t *testing.T) {
	a := assert.New(t)
	err := os.Setenv("DUMMY_VAR", TestId)
	if err != nil {
		return
	}
	a.Equal(TestId, envVar("DUMMY_VAR"))
}

func TestGroupId(t *testing.T) {
	a := assert.New(t)
	a.NotEqualf(TestId, GroupId(), "They are not supposed to be equal, but they are.")
}

func TestClientId(t *testing.T) {
	a := assert.New(t)
	a.NotEqualf(TestId, ClientId(), "They are not supposed to be equal, but they are.")
}

func TestClientSecret(t *testing.T) {
	a := assert.New(t)
	a.NotEqualf(TestId, ClientSecret(), "They are not supposed to be equal, but they are.")
}

func TestTenantId(t *testing.T) {
	a := assert.New(t)
	a.NotEqualf(TestId, TenantId(), "They are not supposed to be equal, but they are.")
}

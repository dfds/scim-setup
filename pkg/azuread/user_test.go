package azuread

import (
	"github.com/dfds/scim-setup/pkg/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	_, err := config.ReadConfigFile(".scim-setup")
	if err != nil {
		return
	}
}

func TestGetUser(t *testing.T) {
	a := assert.New(t)
	a.Panics(func() {
		_, _, _, err := GetUser(nil, "dummy@example.com")
		if err != nil {
			return
		}
	})
}

func TestGetUserId(t *testing.T) {
	a := assert.New(t)
	a.Panics(func() {
		_, err := GetUserId(nil, "dummy@example.com")
		if err != nil {
			return
		}
	})
}

func TestGetUserName(t *testing.T) {
	a := assert.New(t)
	a.Panics(func() {
		_, err := GetUserName(nil, "dummy@example.com")
		if err != nil {
			return
		}
	})
}

func TestGetUserEmail(t *testing.T) {
	a := assert.New(t)
	a.Panics(func() {
		_, err := GetUserEmail(nil, "dummy@example.com")
		if err != nil {
			return
		}
	})
}

package azuread

import (
	"github.com/dfds/scim-setup/pkg/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	config.ReadConfigFile(".scim-setup")
}

func TestGetBearerToken(t *testing.T) {
	a := assert.New(t)
	token := GetBearerToken()
	a.Equal("Bearer ", token[0:7])
}

package config

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReadConfigFile(t *testing.T) {
	a := assert.New(t)
	retVal, err := ReadConfigFile("dummy.txt")
	a.Equal(false, retVal)
	a.True(strings.Contains(err.Error(), "Config File \"dummy.txt\" Not Found in"))
}

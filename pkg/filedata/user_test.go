package filedata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUsers(t *testing.T) {
	a := assert.New(t)
	s := GetUsers("fixtures/test-users.txt")
	a.NotEmpty(s)
	a.Len(s, 4)
	a.Equal("daisy@example.com", s[1])
}

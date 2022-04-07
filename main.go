package main

import (
	"github.com/dfds/scim-setup/pkg/cmd"
	"github.com/dfds/scim-setup/pkg/config"
	"os"
)

func main() {
	// TODO: Replace this dirty hack with the cobra CLI
	usersFile := "users.txt"
	args := os.Args
	if len(args) >= 2 {
		usersFile = args[1]
	}
	config.ReadConfigFile(".scim-setup")
	cmd.LoadUsersIntoGroups(usersFile)
}

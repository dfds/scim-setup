package filedata

import (
	"bufio"
	"os"

	"github.com/dfds/scim-setup/pkg/logging"
)

// GetUsers return a slice of user emails from a file
func GetUsers(fileName string) []string {
	log := logging.GetLogger()
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(f)

	var users []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		users = append(users, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
	}

	return users
}

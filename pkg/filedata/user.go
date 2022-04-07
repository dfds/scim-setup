package filedata

import (
	"bufio"
	"log"
	"os"
)

// GetUsers return a slice of user emails from a file
func GetUsers(fileName string) []string {
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

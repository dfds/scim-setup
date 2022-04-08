package filedata

import (
	"bufio"
	"go.uber.org/zap"
	"os"
)

// GetUsers return a slice of user emails from a file
func GetUsers(fileName string) []string {
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
		}
	}(logger)
	sugar := logger.Sugar()
	f, err := os.Open(fileName)
	if err != nil {
		sugar.Fatal(err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			sugar.Fatal(err.Error())
		}
	}(f)

	var users []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		users = append(users, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		sugar.Fatal(err.Error())
	}

	return users
}

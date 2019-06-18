package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	_, err = getGitDir(dir)
	if err != nil {
		fmt.Println("This is not a git repository")
	} else {
		fmt.Println("This is a Git repository!")
	}

}

func getGitDir(path string) (string, error) {
	for curDir := path; curDir != filepath.Dir(curDir); curDir = filepath.Dir(curDir) {
		_, err := os.Stat(filepath.Join(curDir, ".git"))
		if err == nil {
			return curDir, nil
		}
	}

	return "", errors.New("No Git directory found")
}

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
	gitDir, err := getGitDir(dir)
	if err != nil {
		fmt.Printf("Git dir is located in: %s\n", gitDir)
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

func ignorePath(ignoreFile string, path string) error {
	f, err := os.OpenFile(ignoreFile, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintln(f, path)
	return nil
}

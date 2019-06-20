package main

import (
	"errors"
	"flag"
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
		fmt.Println("This is not a Git repository")
		return
	}
	flag.Parse()
	pathToIgnore := flag.Arg(0)
	if pathToIgnore == "" {
		return
	}
	ignoreFile := filepath.Join(gitDir, ".gitignore")
	ignorePath(ignoreFile, pathToIgnore)
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

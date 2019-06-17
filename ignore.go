package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		return
	}

	if isGitDir(dir) {
		fmt.Println("This is not a git repository")
	} else {
		fmt.Println("This is a Git repository!")
	}

}

func isGitDir(path string) bool {
	for curDir := path; curDir != filepath.Dir(curDir); curDir = filepath.Dir(curDir) {
		_, err := os.Stat(filepath.Join(curDir, ".git"))
		if err == nil {
			return true
		}
	}

	return false
}

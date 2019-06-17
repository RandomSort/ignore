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
	if _, err := os.Stat(filepath.Join(path, ".git")); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

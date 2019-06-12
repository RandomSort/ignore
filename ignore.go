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
	if _, err := os.Stat(filepath.Join(dir, ".git")); os.IsNotExist(err) {
		fmt.Println("This is not a git repository")
	} else {
		fmt.Println("This is a Git repository!")
	}

}

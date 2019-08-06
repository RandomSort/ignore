package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
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
	iFile := filepath.Join(gitDir, ".gitignore")
	ignorePath(iFile, pathToIgnore)
	iFileObj := new(ignoreFile)
	iFileObj.path = iFile
	iFileObj.LoadLines()
	iFileObj.PrintLines()

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

type ignoreFile struct {
	path  string
	lines []string
}

func (iFile *ignoreFile) PrintLines() {
	for _, line := range iFile.lines {
		fmt.Println(line)
	}
}
func (iFile *ignoreFile) LoadLines() {
	f, err := os.Open(iFile.path)
	defer f.Close()
	if err != nil {
		log.Fatalf("Failed to open file: %s", iFile.path)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		line := scanner.Text()
		iFile.lines = append(iFile.lines, line)
	}

}

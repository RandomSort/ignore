package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestNoGitDir(t *testing.T) {
	testDir := getTestDir(t)
	defer os.RemoveAll(testDir)

	gitDir, err := getGitDir(testDir)
	if err == nil {
		t.Errorf("%s shouldn't be a Git repo, but found one at %s", testDir, gitDir)
	}
}

func TestParentGitDit(t *testing.T) {
	testDir := getTestDir(t)
	defer os.RemoveAll(testDir)

	err := os.Mkdir(filepath.Join(testDir, "subdir"), 644)
	err = os.Mkdir(filepath.Join(testDir, ".git"), 644)
	gitDir, err := getGitDir(filepath.Join(testDir, "subdir"))
	if err != nil {
		t.Errorf("Expected to find .gitdir in parent folder of: %v", filepath.Join(testDir, "subdir"))
	}
	if gitDir != testDir {
		t.Errorf("Expected %s, got %s", testDir, gitDir)
	}
}

func TestCWDGitDir(t *testing.T) {
	testDir := getTestDir(t)
	defer os.RemoveAll(testDir)

	err := os.Mkdir(filepath.Join(testDir, ".git"), 644)
	gitDir, err := getGitDir(testDir)
	if err != nil {
		t.Errorf("testDir should be a Git repo, but seems not to be")
	}
	if gitDir != testDir {
		t.Errorf("Expected %v, got %v", testDir, gitDir)
	}
}

func TestAppendPath(t *testing.T) {
	testDir := getTestDir(t)
	defer os.RemoveAll(testDir)

	ignoreFilePath := filepath.Join(testDir, ".gitignore")
	f, err := os.OpenFile(ignoreFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if ( err != nil ){
		t.Errorf("Couldn't open ignoreFile at path %s", ignoreFilePath)
	}
	_, err = fmt.Fprintln(f, "ignoredfile")
	if(err != nil) {
		t.Errorf("Couldn't write to path %s caused by %s", ignoreFilePath, err)
	}
	expectedFilePath := filepath.Join(testDir, "expectedfile")
	expectedFile, err := os.OpenFile(expectedFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if (err != nil ) {
		t.Errorf("Couldn't open expected file at path %s", expectedFilePath)
	}
	fmt.Fprintln(expectedFile, "ignoredfile")
	fmt.Fprintln(expectedFile, "filename.go")
	err = ignorePath(ignoreFilePath, "filename.go")
	if (err != nil) {
		t.Errorf("Couldn't ignore path caused by %s", err)
	}
	expected, err := ioutil.ReadFile(expectedFilePath)
	if (err != nil) {
		t.Errorf("Couldn't readfile at path %s", expectedFilePath)
	}
	dat, err := ioutil.ReadFile(ignoreFilePath)
	if (err != nil) {
		t.Errorf("Couldn't ReadFile at path %s", ignoreFilePath)
	}
	if fmt.Sprintf("%s", dat) != fmt.Sprintf("%s", expected) {
		t.Errorf("Expected: %s got %s", expected, dat)
	}
}

func TestLoadLines(t *testing.T) {
	testDir := getTestDir(t)
	defer os.RemoveAll(testDir)

	ignoreFilePath := filepath.Join(testDir, ".gitignore")
	f, _ := os.OpenFile(ignoreFilePath, os.O_APPEND|os.O_CREATE, 0644)
	expected := []string{"1", "2", "3"}
	for _, v := range expected {
		fmt.Fprintln(f, v)
	}
	iFile := new(ignoreFile)
	iFile.path = ignoreFilePath
	iFile.LoadLines()
	for i := range iFile.lines {
		if iFile.lines[i] != expected[i] {
			t.Errorf("Expected %s, found %s", expected[i], iFile.lines[i])
		}

	}

}

func getTestDir(t *testing.T) string {
	testDir, err := ioutil.TempDir("", "AppendPath")
	if err != nil {
		t.Errorf("Was unable to get workdir(%v). Should never fail. Got Error: %v", testDir, err)
	}
	return testDir
}

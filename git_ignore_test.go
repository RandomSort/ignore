package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestNoGitDir(t *testing.T) {
	testDir, err := ioutil.TempDir("", "TestGitDir")
	defer os.RemoveAll(testDir)
	if err != nil {
		t.Errorf("Was unable to get workdir(%v). Should never fail. Got error: %v", testDir, err)
	}
	gitDir, err := getGitDir(testDir)
	if err == nil {
		t.Errorf("%s shouldn't be a Git repo, but found one at %s", testDir, gitDir)
	}
}

func TestParentGitDit(t *testing.T) {
	testDir, err := ioutil.TempDir("", "GitDir")
	defer os.RemoveAll(testDir)
	if err != nil {
		t.Errorf("Was unable to get workdir(%v). Should never fail. Got error: %v", testDir, err)
	}
	err = os.Mkdir(filepath.Join(testDir, "subdir"), 644)
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
	testDir, err := ioutil.TempDir("", "NoGitDir")
	defer os.RemoveAll(testDir)
	if err != nil {
		t.Errorf("Was unable to get workdir(%v). Should never fail. Got error: %v", testDir, err)
	}

	err = os.Mkdir(filepath.Join(testDir, ".git"), 644)
	gitDir, err := getGitDir(testDir)
	if err != nil {
		t.Errorf("testDir should be a Git repo, but seems not to be")
	}
	if gitDir != testDir {
		t.Errorf("Expected %v, got %v", testDir, gitDir)
	}
}

func TestAppendPath(t *testing.T) {
	testDir, err := ioutil.TempDir("", "AppendPath")
	t.Log(testDir)
	if err != nil {
		t.Errorf("Was unable to get workdir(%v). Should never fail. Got Error: %v", testDir, err)
	}
	defer os.RemoveAll(testDir)
	f, err := os.OpenFile(filepath.Join(testDir, ".gitignore"), os.O_APPEND|os.O_CREATE, 0644)
	fmt.Fprintln(f, "ignoredfile")
	expectedFile, _ := os.OpenFile(filepath.Join(testDir, "expectedfile"), os.O_APPEND|os.O_CREATE, 0664)
	fmt.Fprintln(expectedFile, "ignoredfile")
	fmt.Fprintln(expectedFile, "filename.go")
	ignorePath(filepath.Join(testDir, ".gitignore"), "filename.go")
	expected, _ := ioutil.ReadFile(filepath.Join(testDir, "expectedfile"))
	dat, err := ioutil.ReadFile(filepath.Join(testDir, ".gitignore"))
	if fmt.Sprintf("%s", dat) != fmt.Sprintf("%s", expected) {
		t.Errorf("Expected: %s got %s", expected, dat)
	}
}

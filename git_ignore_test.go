package main

import (
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

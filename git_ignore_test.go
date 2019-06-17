package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGitDir(t *testing.T) {
	testDir, err := ioutil.TempDir("", "GitDir")
	if err != nil {
		t.Errorf("Was unable to get workdir(%v). Should never fail. Got error: %v", testDir, err)
	}
	defer os.RemoveAll(testDir)
	if isGitDir(testDir) {
		t.Errorf("testDir shouldn't be a Git repo, but seems not to be")
	}
}

func TestNoGitDir(t *testing.T) {
	testDir, err := ioutil.TempDir("", "NoGitDir")
	if err != nil {
		t.Errorf("Was unable to get workdir(%v). Should never fail. Got error: %v", testDir, err)
	}
	defer os.RemoveAll(testDir)
	err = os.Mkdir(filepath.Join(testDir, ".git"), 644)
	if !isGitDir(testDir) {
		t.Errorf("testDir should be a Git repo, but seems not to be")
	}
}

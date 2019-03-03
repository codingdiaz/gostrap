package main

import (
	"io/ioutil"
	"testing"
	"fmt"
)

const (
	testDir string = "test_playground"
)

func TestSaveFile(t *testing.T) {
	err := ioutil.WriteFile(
		fmt.Sprintf("%s/%s",testDir, "example"),
		[]byte("hello\n"),
		0644,
	)
	if err != nil {
		t.Errorf("Something is really bad!")
	}
	file := ProjectFile{
		fmt.Sprintf("%s/%s",testDir, "example"),
		fmt.Sprintf("example_saved"),
		false,
	}
	project := Project {LocalRepoPath: testDir}
	err = saveFile(&file, &project)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Was unable to save a file! failing!")
	}
}
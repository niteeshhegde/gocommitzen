package main

import (
	"testing"
)

func TestFileExists(t *testing.T) {
	exists := fileExists("config.json")
	if !exists {
		t.Errorf("File config.json exists.\n Expected ---> true, \n Recieved ---> %v", exists)
	}

	exists = fileExists("config_test.json")
	if exists {
		t.Errorf("File config.json exists.\n Expected ---> false, \n Recieved ---> %v", exists)
	}

}

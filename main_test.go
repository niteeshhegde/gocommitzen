package main

import (
	"testing"
)

func TestFileExists(t *testing.T) {
	t.Run("config file exists", func(t *testing.T) {
		exists := fileExists("commit.json")
		if !exists {
			t.Errorf("File config.json exists.\n Expected ---> true, \n Recieved ---> %v", exists)
		}
	})

	t.Run("config file doesnot exists", func(t *testing.T) {
		exists := fileExists("config_test.json")
		if exists {
			t.Errorf("File config.json exists.\n Expected ---> false, \n Recieved ---> %v", exists)
		}
	})

}

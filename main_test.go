package main

import (
	"testing"
)

func TestNewMain(t *testing.T) {
tests := map[string]struct {
	name string
	valid  bool
}{
	"config file exits": {"commit.json", true},
	"config file doesnot exists": {"config_test.json", false},
}

t.Parallel()
for name, file := range tests {
	file := file
	t.Run(name, func(t *testing.T) {
		exists := fileExists(file.name)
		if !exists && file.valid {
			t.Errorf("Error while checking if file exists\n Expected ---> %v to exist, \n Recieved ---> %v", file.name,exists)
		}
		if exists && !file.valid {
			t.Errorf("Error while checking if file exists\n Expected ---> %v to not exist, \n Recieved ---> %v", file.name, exists)
		}
	})
	}
}
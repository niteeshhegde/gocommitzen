package main

import (
	"reflect"
	"testing"
)

var fileConfig = Config{
	Type: Type{
		MinLength:   1,
		MaxLength:   5,
		AcceptExtra: true,
		Required:    true,
		Values:      []string{"feat", "fix"},
	},
	Scope: Scope{
		MinLength:   0,
		MaxLength:   10,
		AcceptExtra: true,
		Required:    true,
		Values:      []string{},
	},
	Description: Description{
		MinLength: 1,
		MaxLength: 44,
		Required:  true,
	},
	Subject: Subject{
		MinLength: 1,
		MaxLength: 50,
		Required:  true,
	},
	Body: Body{
		Wrap:     72,
		Required: true,
	},
	Footer: Footer{
		Wrap:     72,
		Required: true,
	},
}

func TestNewConfig(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		config, err := newConfig("")
		if !reflect.DeepEqual(config, defaultConfig) {
			t.Errorf("Default Config was not created properly.\n Expected ---> %v, \n Recieved ---> %v", defaultConfig, config)
		}
		if err != nil {
			t.Errorf("Error while creating Default Config %s", err)
		}
	})

	t.Run("custom config", func(t *testing.T) {
		config, err := newConfig("commit.json")
		if !reflect.DeepEqual(config, fileConfig) {
			t.Errorf("Config from commit.json was not created properly.\n Expected ---> %v, \n Recieved ---> %v", defaultConfig, config)
		}
		if err != nil {
			t.Errorf("Error while creating Config from file %s", err)
		}
	})

	t.Run("wrong custom config", func(t *testing.T) {
		_, err := newConfig("commit_test.json")
		expectedErr := "open commit_test.json: no such file or directory"
		if err.Error() != expectedErr {
			t.Errorf("Error while creating Config from file %s.\n Execting error %s", err, expectedErr)
		}
	})
}

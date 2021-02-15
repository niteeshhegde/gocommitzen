package main

import (
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

func arraysAreEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestNewConfig(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		config, err := newConfig("")
		if config.Type.AcceptExtra != defaultConfig.Type.AcceptExtra || config.Type.Required != defaultConfig.Type.Required || config.Type.MinLength != defaultConfig.Type.MinLength || config.Type.MaxLength != defaultConfig.Type.MaxLength || !arraysAreEqual(config.Type.Values, defaultConfig.Type.Values) {
			t.Errorf("Default Config Type was not created properly.\n Expected ---> %v, \n Recieved ---> %v", defaultConfig, config)
		}
		if config.Scope.AcceptExtra != defaultConfig.Scope.AcceptExtra || config.Scope.Required != defaultConfig.Scope.Required || config.Scope.MinLength != defaultConfig.Scope.MinLength || config.Scope.MaxLength != defaultConfig.Scope.MaxLength || !arraysAreEqual(config.Scope.Values, defaultConfig.Scope.Values) {
			t.Errorf("Default Config Scope was not created properly.\n Expected ---> %v, \n Recieved ---> %v", defaultConfig, config)
		}
		if config.Description != defaultConfig.Description || config.Body != defaultConfig.Body || config.Footer != defaultConfig.Footer {
			t.Errorf("Default Config was not created properly.\n Expected ---> %v, \n Recieved ---> %v", defaultConfig, config)
		}
		if err != nil {
			t.Errorf("Error while creating Default Config %s", err)
		}
	})

	t.Run("custom config", func(t *testing.T) {
		config, err := newConfig("commit.json")
		if config.Type.AcceptExtra != fileConfig.Type.AcceptExtra || config.Type.Required != fileConfig.Type.Required || config.Type.MinLength != fileConfig.Type.MinLength || config.Type.MaxLength != fileConfig.Type.MaxLength || !arraysAreEqual(config.Type.Values, fileConfig.Type.Values) {
			t.Errorf("File Config Type was not created properly.\n Expected ---> %v, \n Recieved ---> %v", fileConfig, config)
		}
		if config.Scope.AcceptExtra != fileConfig.Scope.AcceptExtra || config.Scope.Required != fileConfig.Scope.Required || config.Scope.MinLength != fileConfig.Scope.MinLength || config.Scope.MaxLength != fileConfig.Scope.MaxLength || !arraysAreEqual(config.Scope.Values, fileConfig.Scope.Values) {
			t.Errorf("File Config Scope was not created properly.\n Expected ---> %v, \n Recieved ---> %v", fileConfig, config)
		}
		if config.Description != fileConfig.Description || config.Body != fileConfig.Body || config.Footer != fileConfig.Footer {
			t.Errorf("File Config was not created properly.\n Expected ---> %v, \n Recieved ---> %v", fileConfig, config)
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

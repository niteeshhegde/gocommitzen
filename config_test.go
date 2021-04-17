package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var wantFileConfig = Config{
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
	config, err := newConfig("")
	if err != nil {
		t.Fatalf("Error while creating Default Config %s", err)
	}
	givenFileConfig, err := newConfig("commit.json")
	if err != nil {
		t.Fatalf("Error while creating file Config %s", err)
	}

	tests := map[string]struct {
		given Config
		want  Config
	}{
		"default config": {
			given: config,
			want:  defaultConfig,
		},
		"file config": {
			given: givenFileConfig,
			want:  wantFileConfig,
		},
	}

	t.Parallel()
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			compareTypes(t, test.given.Type, test.want.Type)
			if test.given.Scope.AcceptExtra != test.want.Scope.AcceptExtra || test.given.Scope.Required != test.want.Scope.Required || test.given.Scope.MinLength != test.want.Scope.MinLength || test.given.Scope.MaxLength != test.want.Scope.MaxLength || !arraysAreEqual(test.given.Scope.Values, test.want.Scope.Values) {
				t.Errorf("Default Config Scope was not created properly.\n Expected ---> %v, \n Recieved ---> %v", test.want, test.given)
			}
			if test.given.Description != test.want.Description || test.given.Body != test.want.Body || test.given.Footer != test.want.Footer {
				t.Errorf("Default Config was not created properly.\n Expected ---> %v, \n Recieved ---> %v", test.want, test.given)
			}
		})
	}

	t.Run("wrong custom config", func(t *testing.T) {
		_, err := newConfig("commit_test.json")
		var wantErr *os.PathError
		if !errors.As(err, &wantErr) || wantErr.Op != "open" || wantErr.Path != "commit_test.json" {
			t.Errorf("Error while creating Config from file %s.\n Execting error open error", err)
		}
	})
}

func compareTypes(t *testing.T, a, b Type) {
	t.Helper()

	if a.AcceptExtra != b.AcceptExtra {
		t.Errorf("could not match accept extra: %#v, %#v", a.AcceptExtra, b.AcceptExtra)
	}
	if a.Required != b.Required {
		t.Errorf("could not match required: %#v, %#v", a.Required, b.Required)
	}
	if a.MinLength != b.MinLength {
		t.Errorf("could not match min length: %#v, %#v", a.MinLength, b.MinLength)
	}
	if a.MaxLength != b.MaxLength {
		t.Errorf("could not match max length: %#v, %#v", a.MaxLength, b.MaxLength)
	}
	if fmt.Sprintf("%s", a.Values) != fmt.Sprintf("%s", b.Values) {
		t.Errorf("could not match values: %#v, %#v", a.Values, b.Values)
	}
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

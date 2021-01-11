package main

import (
	"encoding/json"
	"os"
)

var defaultConfig = Config{
	Type: struct {
		MinLength   int      `json:"minLength"`
		MaxLength   int      `json:"maxLength"`
		AcceptExtra bool     `json:"acceptExtra"`
		Values      []string `json:"values"`
	}{
		MinLength:   1,
		MaxLength:   5,
		AcceptExtra: true,
		Values:      []string{"feat", "fix"},
	},
	Scope: struct {
		MinLength   int      `json:"minLength"`
		MaxLength   int      `json:"maxLength"`
		AcceptExtra bool     `json:"acceptExtra"`
		Values      []string `json:"values"`
	}{
		MinLength:   0,
		MaxLength:   10,
		AcceptExtra: false,
		Values:      []string{},
	},
	Description: struct {
		MinLength int `json:"minLength"`
		MaxLength int `json:"maxLength"`
	}{
		MinLength: 1,
		MaxLength: 44,
	},
	Subject: struct {
		MinLength int `json:"minLength"`
		MaxLength int `json:"maxLength"`
	}{
		MinLength: 1,
		MaxLength: 50,
	},
	Body: struct {
		Wrap int `json:"wrap"`
	}{
		Wrap: 72,
	},
	Footer: struct {
		Wrap int `json:"wrap"`
	}{
		Wrap: 72,
	},
} //add from config

type Config struct {
	Type struct {
		MinLength   int      `json:"minLength"`
		MaxLength   int      `json:"maxLength"`
		AcceptExtra bool     `json:"acceptExtra"`
		Values      []string `json:"values"`
	} `json:"type"`
	Scope struct {
		MinLength   int      `json:"minLength"`
		MaxLength   int      `json:"maxLength"`
		AcceptExtra bool     `json:"acceptExtra"`
		Values      []string `json:"values"`
	} `json:"scope"`
	Description struct {
		MinLength int `json:"minLength"`
		MaxLength int `json:"maxLength"`
	} `json:"description"`
	Subject struct {
		MinLength int `json:"minLength"`
		MaxLength int `json:"maxLength"`
	} `json:"subject"`
	Body struct {
		Wrap int `json:"wrap"`
	} `json:"body"`
	Footer struct {
		Wrap int `json:"wrap"`
	} `json:"footer"`
}

func newConfig(path string) (Config, error) {

	var cng Config
	if len(path) > 0 {
		file, err := os.Open(path)
		if err != nil {
			return cng, err
		}
		defer file.Close()
		if err := json.NewDecoder(file).Decode(&cng); err != nil {
			return cng, err
		}
		return cng, err
	}
	return defaultConfig, nil
}

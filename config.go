package main

import (
	"encoding/json"
	"os"
)

var defaultConfig = Config{
	Type: Type{
		MinLength:   1,
		MaxLength:   5,
		AcceptExtra: true,
		Required:    false,
		Values:      []string{"feat", "fix"},
	},
	Scope: Scope{
		MinLength:   1,
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
		Wrap:     2,
		Required: true,
	},
	Footer: Footer{
		Wrap:     2,
		Required: true,
	},
} //add from config

type Config struct {
	Type        Type        `json:"type"`
	Scope       Scope       `json:"scope"`
	Description Description `json:"description"`
	Subject     Subject     `json:"subject"`
	Body        Body        `json:"body"`
	Footer      Footer      `json:"footer"`
}

type Type struct {
	MinLength   int      `json:"minLength"`
	MaxLength   int      `json:"maxLength"`
	AcceptExtra bool     `json:"acceptExtra"`
	Required    bool     `json:"required"`
	Values      []string `json:"values"`
}

// Scope fs
type Scope struct {
	MinLength   int      `json:"minLength"`
	MaxLength   int      `json:"maxLength"`
	AcceptExtra bool     `json:"acceptExtra"`
	Required    bool     `json:"required"`
	Values      []string `json:"values"`
}

type Description struct {
	MinLength int  `json:"minLength"`
	MaxLength int  `json:"maxLength"`
	Required  bool `json:"required"`
}

type Subject struct {
	MinLength int  `json:"minLength"`
	MaxLength int  `json:"maxLength"`
	Required  bool `json:"required"`
}

type Body struct {
	Wrap     int  `json:"wrap"`
	Required bool `json:"required"`
}

type Footer struct {
	Wrap     int  `json:"wrap"`
	Required bool `json:"required"`
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

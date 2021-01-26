package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var messages = map[string]string{
	"type":        "Commits MUST be prefixed with a type, which consists of a noun, feat, fix, etc.",
	"scope":       "An optional scope MAY be provided after a type. A scope is a phrase describing a section of the codebase.",
	"description": "A description MUST immediately follow the type/scope prefix. The description is a short description of the changes",
	"body":        "A longer commit body MAY be provided after the short description.",
	"footer":      "A footer MAY be provided one blank line after the body. The footer SHOULD contain additional meta-information about the changes(such as the issues it fixes, e.g., fixes #13, #5).",
}

func createMessage(config interface{}, personalized bool, name string, reader bufio.Reader) string {
	var values []string
	var valueInput string
	var acceptExtra, required bool
	var minLength, maxLength int

	printHeader(fmt.Sprintf("\n-------%s-------", strings.ToUpper(name)))
	printDescrition(messages[name])

	v := reflect.ValueOf(config)
	if v.FieldByName("MinLength").IsValid() && v.FieldByName("MaxLength").IsValid() {
		minLength = int(v.FieldByName("MinLength").Int())
		maxLength = int(v.FieldByName("MaxLength").Int())
		printDescrition(fmt.Sprintf("Length of this field must be between %d and %d", minLength, maxLength))
	} else if v.FieldByName("MaxLength").IsValid() {
		minLength = int(v.FieldByName("MinLength").Int())
		printDescrition(fmt.Sprintf("Minimum length of this field must be %d", minLength))
	} else if v.FieldByName("MaxLength").IsValid() {
		maxLength = int(v.FieldByName("MaxLength").Int())
		printDescrition(fmt.Sprintf("Maximum length of this field must be %d", maxLength))
	} else if v.FieldByName("Wrap").IsValid() {
		maxLength = int(v.FieldByName("Wrap").Int())
		minLength = 1
		printDescrition(fmt.Sprintf("Maximum length of this field must be %d", maxLength))
	}

	if v.FieldByName("AcceptExtra").IsValid() && v.FieldByName("AcceptExtra").Bool() == true {
		acceptExtra = true
	}

	if v.FieldByName("Required").IsValid() && v.FieldByName("Required").Bool() == true {
		printDescrition("This field is Required")
		required = true
	} else {
		printDescrition("This field is Optional")
	}

	if v.FieldByName("Values").IsValid() && v.FieldByName("Values").Len() > 0 && !personalized {

		values = v.FieldByName("Values").Interface().([]string)

		if v.FieldByName("AcceptExtra").IsValid() && v.FieldByName("AcceptExtra").Bool() == true {
			values = append(values, extraKeyName)
		}

		if !required {
			values = append(values, skipKeyName)
		}

		for valueInput == "" {
			printInput(fmt.Sprintf("Select commit's %s from the below -", name))
			for index, element := range values {
				printInput(fmt.Sprintf("%d - %s", index, element))
			}
			input, _, err := reader.ReadLine()
			if err != nil {
				panic(err)
			}

			idx, err := strconv.Atoi(string(input))
			if int(idx) < len(values) && idx >= 0 && err == nil {
				valueInput = values[idx]
				if valueInput == skipKeyName {
					printSkipping(fmt.Sprintf("Skipping %s as no values entered!", name))
					return ""
				}
			} else {
				printError(fmt.Sprintf("Wrong Inupt! Please enter a number between 0 and %d.",
					len(values)-1))
			}
		}

	}

	if (!v.FieldByName("Values").IsValid() || acceptExtra) && (valueInput == extraKeyName || valueInput == "") {
		printInput(fmt.Sprintf("Enter commit's %s:", name))
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		valueInput = string(line)
		if !required && (valueInput == "") {
			printSkipping(fmt.Sprintf("Skipping %s as no values entered!", name))
			return ""
		}
		for len(valueInput) > maxLength || len(valueInput) < minLength {
			printError(fmt.Sprintf("Input Length must be between %d and %d characters. Please enter again", minLength, maxLength))
			line, _, err := reader.ReadLine()
			if err != nil {
				panic(err)
			}
			valueInput = string(line)
			if !required && (valueInput == "") {
				printSkipping(fmt.Sprintf("Skipping %s as no values entered!", name))
				return ""
			}
		}
	}

	return valueInput
}

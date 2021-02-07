package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	wordwrap "github.com/mitchellh/go-wordwrap"
)

var messages = map[string]string{
	"type":        "Commits MUST be prefixed with a type, which consists of a noun, feat, fix, etc.",
	"scope":       "An optional scope MAY be provided after a type. A scope is a phrase describing a section of the codebase.",
	"description": "A description MUST immediately follow the type/scope prefix. The description is a short description of the changes",
	"body":        "A longer commit body MAY be provided after the short description.",
	"footer":      "A footer MAY be provided one blank line after the body. The footer SHOULD contain additional meta-information about the changes(such as the issues it fixes, e.g., fixes #13, #5).",
}

func createMessage(cnf interface{}, personalized bool, name string, reader bufio.Reader) string {
	var values []string
	var valueInput string
	var minLength, maxLength, wrap int
	var required, acceptExtra bool
	switch cnf.(type) {

	case Type:
		minLength = cnf.(Type).MinLength
		maxLength = cnf.(Type).MaxLength
		values = cnf.(Type).Values
		acceptExtra = cnf.(Type).AcceptExtra
		required = cnf.(Type).Required

	case Scope:
		minLength = cnf.(Scope).MinLength
		maxLength = cnf.(Scope).MaxLength
		values = cnf.(Scope).Values
		acceptExtra = cnf.(Scope).AcceptExtra
		required = cnf.(Scope).Required

	case Description:
		minLength = cnf.(Description).MinLength
		maxLength = cnf.(Description).MaxLength
		required = cnf.(Description).Required

	case Body:
		wrap = cnf.(Body).Wrap
		required = cnf.(Body).Required

	case Footer:
		wrap = cnf.(Footer).Wrap
		required = cnf.(Footer).Required

	default:
		panic("whatever 'i' is, it is not a Dog or Cat")
	}
	printHeader(fmt.Sprintf("\n-------%s-------", strings.ToUpper(name)))
	printDescrition(messages[name])

	if (minLength > 0) && (maxLength > 0) {
		printDescrition(fmt.Sprintf("Length of this field must be between %d and %d", minLength, maxLength))
	} else if maxLength > 0 {
		printDescrition(fmt.Sprintf("Maximum length of this field must be %d", maxLength))
	} else if minLength > 0 {
		printDescrition(fmt.Sprintf("Minimum length of this field must be %d", minLength))
		maxLength = math.MaxInt32
	} else {
		maxLength = math.MaxInt32
	}
	fmt.Println(maxLength)

	if required {
		printDescrition("This field is Required")
	} else {
		printDescrition("This field is Optional")
	}

	if len(values) > 0 && !personalized {

		if acceptExtra {
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

			idx, err := strconv.Atoi(readInput(&reader, wrap))
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

	if (len(values) == 0 || acceptExtra) && (valueInput == extraKeyName || valueInput == "") {
		printInput(fmt.Sprintf("Enter commit's %s:", name))

		valueInput = readInput(&reader, wrap)
		if !required && (valueInput == "") {
			printSkipping(fmt.Sprintf("Skipping %s as no values entered!", name))
			return ""
		}
		for len(valueInput) > maxLength || len(valueInput) < minLength {
			printError(fmt.Sprintf("Input Length must be between %d and %d characters. Please enter again", minLength, maxLength))
			valueInput = readInput(&reader, wrap)
			if !required && (valueInput == "") {
				printSkipping(fmt.Sprintf("Skipping %s as no values entered!", name))
				return ""
			}
		}
	}

	return valueInput
}

func readInput(r *bufio.Reader, wrap int) string {

	text, err := r.ReadString('\n')
	if err != nil {
		printError(fmt.Sprintf("could not read from stdin %v\n", err))
		os.Exit(1)
	}

	if wrap > 0 {
		return wordwrap.WrapString(text, uint(wrap))
	}

	return strings.TrimSuffix(text, "\n")

}

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func createScope(config Config, personalized bool) string {
	printHeader("\n-------SCOPE-------")
	printDescrition("An optional scope MAY be provided after a type. A scope is a phrase describing a section of the codebase.")
	printDescrition(fmt.Sprintf("Type length must be between %d and %d", config.Scope.MinLength, config.Scope.MaxLength))

	scopeConfig := config.Scope
	var scopeValues = scopeConfig.Values
	var scopeInput string = ""
	if scopeConfig.AcceptExtra == true {
		scopeValues = append(scopeValues, extraKeyName)
	}

	if len(scopeValues) > 0 && !reflect.DeepEqual(scopeValues, extraKeyName) && !personalized {
		printInput("Select commit's scope from the below or any other character to skip it:")
		var idx int
		var input string
		for index, element := range scopeValues {
			printInput(fmt.Sprintf("%d - %s", index, element))
		}
		fmt.Scanf("%s", &input)
		idx, err := strconv.Atoi(input)
		if int(idx) < len(scopeValues) && idx >= 0 && err == nil {
			scopeInput = scopeValues[idx]
		} else {
			printSkipping("Skipping body as no values entered!")
			return scopeInput
		}
	}

	if scopeConfig.AcceptExtra && (scopeInput == extraKeyName || scopeInput == "") {
		printInput("Enter a custom commit's scope:")
		fmt.Scanf("%s", &scopeInput)
		for len(scopeInput) > scopeConfig.MaxLength || len(scopeInput) < scopeConfig.MinLength {
			printError(fmt.Sprintf("Wrong Input! Type length must be between %d and %d", scopeConfig.MinLength, scopeConfig.MaxLength))
			fmt.Scanf("%s", &scopeInput)
		}
	}

	return scopeInput
}

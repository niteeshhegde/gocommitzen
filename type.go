package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func createType(config Config, personalized bool) string {
	printHeader("\n-------TYPE-------")
	printDescrition("Commits MUST be prefixed with a type, which consists of a noun, feat, fix, etc.")
	printDescrition(fmt.Sprintf("Type length must be between %d and %d", config.Type.MinLength, config.Type.MaxLength))

	typeConfig := config.Type
	var typeValues = typeConfig.Values
	var typeInput string = ""
	if typeConfig.AcceptExtra == true {
		typeValues = append(typeValues, extraKeyName)
	}

	if len(typeValues) > 0 && !reflect.DeepEqual(typeValues, extraKeyName) && !personalized {
		for typeInput == "" {
			printInput("Select commit's type:")
			var idx int
			var input string
			for index, element := range typeValues {
				printInput(fmt.Sprintf("%d - %s", index, element))
			}
			fmt.Scanf("%s", &input)
			idx, err := strconv.Atoi(input)
			if int(idx) < len(typeValues) && idx >= 0 && err == nil {
				typeInput = typeValues[idx]
			} else {
				printError(fmt.Sprintf("Wrong Inupt! Please enter a number between 0 and %d",
					len(typeValues)-1))
			}
		}
	}

	if typeConfig.AcceptExtra && (typeInput == extraKeyName || typeInput == "") {
		printInput("Enter a custom commit's type:")
		fmt.Scanf("%s", &typeInput)
		for len(typeInput) > typeConfig.MaxLength || len(typeInput) < typeConfig.MinLength {
			printError(fmt.Sprintf("Wrong Input! Type length must be between %d and %d", typeConfig.MinLength, typeConfig.MaxLength))
			fmt.Scanf("%s", &typeInput)
		}
	}

	return typeInput
}

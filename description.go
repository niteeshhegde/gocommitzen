package main

import "fmt"

func createDescription(config Config) string {
	printHeader("\n-------DESCRIPTION-------")
	printDescrition("A description MUST immediately follow the type/scope prefix. The description is a short description of the changes")
	printDescrition(fmt.Sprintf("Description length must be between %d and %d", config.Description.MinLength, config.Description.MaxLength))

	var descriptionInput string = ""
	for descriptionInput == "" {
		printInput("Enter a commit's description:")
		fmt.Scanln(&descriptionInput)
		if len(descriptionInput) < config.Description.MinLength || len(descriptionInput) > config.Description.MaxLength {
			printError("Wrong Inupt! Please Enter Again.")
			descriptionInput = ""
		}
	}
	return descriptionInput
}

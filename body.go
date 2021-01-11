package main

import "fmt"

func createBody(config Config) string {
	printHeader("\n-------BODY-------")
	printDescrition("A longer commit body MAY be provided after the short description.")

	var bodyInput string = ""
	printInput("Enter a commit's body:")
	fmt.Scanln(&bodyInput)
	if bodyInput == "" {
		printSkipping("Skipping body as no values entered!")
	}

	return bodyInput
}

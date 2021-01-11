package main

import "fmt"

func createFooter(config Config) string {
	printHeader("\n-------FOOTER-------")
	printDescrition("A footer MAY be provided one blank line after the body. The footer SHOULD contain additional meta-information about the changes(such as the issues it fixes, e.g., fixes #13, #5).")

	var footerInput string = ""
	printInput("Enter a commit's footer:")
	fmt.Scanln(&footerInput)
	if footerInput == "" {
		printSkipping("Skipping footer as no values entered! ")
	}
	return footerInput
}

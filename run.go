package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var colorReset string = "\033[0m"
var colorRed string = "\033[31m"
var colorGreen string = "\033[32m"
var colorYellow string = "\033[33m"
var colorBlue string = "\033[34m"
var colorPurple string = "\033[35m"
var colorCyan string = "\033[36m"
var colorWhite string = "\033[37m"

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

var extraKeyName = []string{"personalized"}

func createType(config Config) string {
	fmt.Println()
	fmt.Println(string(colorPurple), "-------TYPE-------", string(colorReset))
	fmt.Println(string(colorCyan), "Commits MUST be prefixed with a type, which consists of a noun, feat, fix, etc.")
	fmt.Printf(" Type length must be between %d and %d \n", config.Type.MinLength, config.Type.MaxLength)

	typeConfig := config.Type
	var typeValues = typeConfig.Values
	var typeInput string = ""
	// if typeConfig.AcceptExtra == true {
	//  typeValues = extraKeyName
	// }

	//if !reflect.DeepEqual(typeValues, extraKeyName) {
	if len(typeValues) > 0 {
		fmt.Println(string(colorGreen), "Select commit's type:")
		var idx int
		for index, element := range typeValues {
			fmt.Println(string(colorGreen), index, element)
			// element is the element from someSlice for where we are
		}
		fmt.Println(string(colorReset), "")
		fmt.Scanln(&idx)
		if idx < len(typeValues) && idx >= 0 {
			typeInput = typeValues[idx]
		} else {
			fmt.Println(string(colorRed), "Wrong Inupt!")
		}
	}

	//}

	for typeInput == "" {
		fmt.Println(string(colorGreen), "Enter a custom commit's type:", string(colorReset))
		fmt.Scanln(&typeInput)
	}
	// if typeInput == "" || reflect.DeepEqual(typeValues, extraKeyName) {
	//  fmt.Println(string(colorGreen), "Enter a custom commit's type:")
	//  fmt.Scanln(&typeInput)
	// }
	return typeInput
}

func createScope(config Config) string {
	fmt.Println()
	fmt.Println(string(colorPurple), "-------SCOPE-------", string(colorReset))
	fmt.Println(string(colorCyan), "An optional scope MAY be provided after a type. A scope is a phrase describing a section of the codebase.")
	fmt.Printf(" Scope length must be between %d and %d \n", config.Scope.MinLength, config.Scope.MaxLength)

	scopeConfig := config.Scope
	var scopeValues = scopeConfig.Values
	var scopeInput string = ""
	// if scopeConfig.AcceptExtra == true {
	//  scopeValues = extraKeyName
	// }

	//if !reflect.DeepEqual(scopeValues, extraKeyName) {
	if len(scopeValues) > 0 {
		fmt.Println(string(colorGreen), "Select commit's scope:")
		var idx int
		for index, element := range scopeValues {
			fmt.Println(string(colorGreen), index, element)
		}
		fmt.Scanln(&idx)
		if idx < len(scopeValues) && idx >= 0 {
			scopeInput = scopeValues[idx]
		} else {
			fmt.Println(string(colorRed), "Wrong Inupt!", string(colorReset))
		}
	}

	//}

	if scopeInput == "" && scopeConfig.AcceptExtra {
		fmt.Println(string(colorGreen), "Enter a custom commit's scope:", string(colorReset))
		fmt.Scanln(&scopeInput)
	}
	if scopeInput == "" {
		fmt.Println(string(colorYellow), "Skipping Scope as no values entered! ", string(colorReset))
	}
	// if scopeInput == "" || reflect.DeepEqual(scopeValues, extraKeyName) {
	//  fmt.Println(string(colorGreen), "Enter a custom commit's scope:")
	//  fmt.Scanln(&scopeInput)
	// }
	return scopeInput
}

func createDescription(config Config) string {
	fmt.Println()
	fmt.Println(string(colorPurple), "-------DESCRIPTION-------", string(colorReset))
	fmt.Println(string(colorCyan), "A description MUST immediately follow the type/scope prefix. The description is a short description of the changes")
	fmt.Printf(" Description length must be between %d and %d \n", config.Description.MinLength, config.Description.MaxLength)

	var descriptionInput string = ""
	for descriptionInput == "" {
		fmt.Println(string(colorGreen), "Enter a commit's description:", string(colorReset))
		fmt.Scanln(&descriptionInput)
		if len(descriptionInput) < config.Description.MinLength || len(descriptionInput) > config.Description.MaxLength {
			fmt.Println(string(colorRed), "Wrong Inupt! Please Enter Again.", string(colorReset))
			descriptionInput = ""
		}
	}
	return descriptionInput
}

func createBody(config Config) string {
	fmt.Println()
	fmt.Println(string(colorPurple), "-------BODY-------", string(colorReset))
	fmt.Println(string(colorCyan), "A longer commit body MAY be provided after the short description.")

	var bodyInput string = ""
	fmt.Println(string(colorGreen), "Enter a commit's body:", string(colorReset))
	fmt.Scanln(&bodyInput)
	if bodyInput == "" {
		fmt.Println(string(colorYellow), "Skipping body as no values entered! ", string(colorReset))
	}

	return bodyInput
}

func createFooter(config Config) string {
	fmt.Println()
	fmt.Println(string(colorPurple), "-------FOOTER-------", string(colorReset))
	fmt.Println(string(colorCyan), "A footer MAY be provided one blank line after the body. The footer SHOULD contain additional meta-information about the changes(such as the issues it fixes, e.g., fixes #13, #5).")

	var footerInput string = ""
	fmt.Println(string(colorGreen), "Enter a commit's footer:", string(colorReset))
	fmt.Scanln(&footerInput)
	if footerInput == "" {
		fmt.Println(string(colorYellow), "Skipping footer as no values entered! ", string(colorReset))
	}
	return footerInput
}

func loadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
func createCommit(subject string, body string, footer string) ([]byte, error) {
	cmd := exec.Command("git", "commit", "-m", subject, "-m", body, "-m", footer)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf(err.Error())
	}

	return stdout, nil
}

func main() {
	// Subcommands
	commitCommand := flag.NewFlagSet("commit", flag.ExitOnError)
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	config := loadConfiguration("/home/niteesh/go/src/gocommitzen/config.json")
	fmt.Println(config)
	typeEntered := createType(config)
	scopeEntered := createScope(config)
	descriptionEntered := createDescription(config)
	bodyEntered := createBody(config)
	footerEntered := createFooter(config)
	subject := typeEntered + scopeEntered + ": " + descriptionEntered
	fmt.Println(string(colorReset), typeEntered, scopeEntered, descriptionEntered, bodyEntered, footerEntered)

	// Count subcommand flag pointers
	// Adding a new choice for --metric of 'substring' and a new --substring flag
	addAllFilesPtr := commitCommand.Bool("a", false, "All files to commit")

	// Use flag.Var to create a flag of our new flagType
	// Default value is the current value at countStringListPtr (currently a nil value)
	// List subcommand flag pointers

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("list or count subcommand is required")
		os.Exit(1)
	}

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]

	switch os.Args[1] {
	case "list":
		addCommand.Parse(os.Args[2:])
	case "commit":
		commitCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
	// if listCommand.Parsed() {
	//  // Required Flags
	//  if *listTextPtr == "" {
	//      listCommand.PrintDefaults()
	//      os.Exit(1)
	//  }
	//  //Choice flag
	//  metricChoices := map[string]bool{"chars": true, "words": true, "lines": true}
	//  if _, validChoice := metricChoices[*listMetricPtr]; !validChoice {
	//      listCommand.PrintDefaults()
	//      os.Exit(1)
	//  }
	//  // Print
	//  fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n",
	//      *listTextPtr,
	//      *listMetricPtr,
	//      *listUniquePtr,
	//  )
	// }

	if commitCommand.Parsed() {
		if *addAllFilesPtr {
			//commitCommand.PrintDefaults()
			cmd := exec.Command("git", "add", ".")
			stdout, err := cmd.Output()

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Print(string(stdout))
			//os.Exit(1)
		}
		cmd := exec.Command("git", "add", ".")
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Print(string(stdout))
	}

	otpt, err := createCommit(subject, bodyEntered, footerEntered)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(otpt), "otpt")

	// // If the metric flag is substring, the substring or substringList flag is required
	// if *countMetricPtr == "substring" && *countSubstringPtr == "" && (&countStringList).String() == "[]" {
	//  countCommand.PrintDefaults()
	//  os.Exit(1)
	// }
	// //If the metric flag is not substring, the substring flag must not be used
	// if *countMetricPtr != "substring" && (*countSubstringPtr != "" || (&countStringList).String() != "[]") {
	//  fmt.Println("--substring and --substringList may only be used with --metric=substring.")
	//  countCommand.PrintDefaults()
	//  os.Exit(1)
	// }
	// //Choice flag
	// metricChoices := map[string]bool{"chars": true, "words": true, "lines": true, "substring": true}
	// if _, validChoice := metricChoices[*listMetricPtr]; !validChoice {
	//  countCommand.PrintDefaults()
	//  os.Exit(1)
	// }
	// //Print
	// fmt.Printf("textPtr: %s, metricPtr: %s, substringPtr: %v, substringListPtr: %v, uniquePtr: %t\n",
	//  *countTextPtr,
	//  *countMetricPtr,
	//  *countSubstringPtr,
	//  (&countStringList).String(),
	//  *countUniquePtr,
	// )
}

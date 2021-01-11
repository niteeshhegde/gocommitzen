package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var extraKeyName string = "personalized"

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
	commitCommand := flag.NewFlagSet("commit", flag.ExitOnError)
	addAllFilesPtr := commitCommand.Bool("a", false, "Boolean value refering to stage all files")
	personalizedCommitPtr := commitCommand.Bool("p", false, "Boolean value allowing you to add custom type and scope skipping choises")
	configFilePtr := commitCommand.String("c", "", "String value refering to path of the config file for conventional commits")
	commitCommand.Usage = func() {
		fmt.Println("Description: gocommitzen provides an interface for using conventional commits for your code")
		printHeader(" 1. commit")
		fmt.Println("  This command asks for all components of the conventional commit and commits the code.")
		printInput(" Usage:")
		printSkipping(" gocommitzen commit -a -p -c 'home/user/config.json'")
		printInput(" Args:")
		commitCommand.PrintDefaults()
	}
	switch os.Args[1] {
	case "commit":
		commitCommand.Parse(os.Args[2:])
	default:
		commitCommand.Usage()
		os.Exit(1)
	}
	if commitCommand.Parsed() == false {
		commitCommand.Usage()
		os.Exit(1)
	}

	config, err := newConfig(*configFilePtr)
	typeEntered := createType(config, *personalizedCommitPtr)
	scopeEntered := createScope(config, *personalizedCommitPtr)
	descriptionEntered := createDescription(config)
	bodyEntered := createBody(config)
	footerEntered := createFooter(config)
	subject := typeEntered + scopeEntered + ": " + descriptionEntered

	if len(os.Args) < 2 {
		fmt.Println("list or count subcommand is required")
		os.Exit(1)
	}

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

	otpt, err := createCommit(subject, bodyEntered, footerEntered)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(otpt), "otpt")

}

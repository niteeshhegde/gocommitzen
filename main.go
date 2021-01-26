package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var extraKeyName string = "personalized"
var skipKeyName string = "skip"

func createCommit(subject string, body string, footer string) ([]byte, error) {
	cmd := exec.Command("git", "commit", "-m", subject, "-m", body, "-m", footer)
	stdout, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return stdout, nil
}

func usage() {
	fmt.Println("Description: gocommitzen provides an interface for using conventional commits for your code")
	printHeader(" 1. commit")
	fmt.Println("  This command asks for all components of the conventional commit and commits the code.")
	printInput(" Usage:")
	printSkipping(" gocommitzen commit -a -p -c 'home/user/config.json'")
	printInput(" Args:")
	flag.PrintDefaults()
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	var a, p bool
	var c string

	flag.BoolVar(&a, "a", false, "Boolean value refering to stage all files")
	flag.BoolVar(&p, "p", false, "Boolean value allowing you to add custom type and scope skipping choises")
	flag.StringVar(&c, "c", "", "String value refering to path of the config file for conventional commits")

	flag.Parse()
	if flag.Parsed() == false {
		usage()
		os.Exit(1)
	}
	if c == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			printError(fmt.Sprintf("%s", err))
		}
		if fileExists(currentDir + "/commit.json") {
			c = currentDir + "/commit.json"
		}
	}

	config, err := newConfig(c)
	if err != nil {
		printError(fmt.Sprintf("%s", err))
		return
	}
	reader := bufio.NewReader(os.Stdin)
	//handle err
	// t := reflect.TypeOf(config)
	// reader := bufio.NewReader(os.Stdin)
	// for i := 0; i < t.NumField(); i++ {
	// 	fmt.Printf("%+v\n", t.Field(i))
	// 	createMessage(t.Field(i), p, t.Field(i).Name, *reader)
	// }

	typeEntered := createMessage(config.Type, p, "type", *reader)
	scopeEntered := createMessage(config.Scope, p, "scope", *reader)
	descriptionEntered := createMessage(config.Description, p, "description", *reader)
	bodyEntered := createMessage(config.Body, p, "body", *reader)
	footerEntered := createMessage(config.Footer, p, "footer", *reader)
	fmt.Println(typeEntered, scopeEntered, descriptionEntered, bodyEntered, footerEntered)
	subject := typeEntered + scopeEntered + ": " + descriptionEntered

	if a {
		cmd := exec.Command("git", "add", ".")
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Print(string(stdout))
	}

	stdout, err := createCommit(subject, bodyEntered, footerEntered)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))

}

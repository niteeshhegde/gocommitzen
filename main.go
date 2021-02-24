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

func createCommit(subject string, body string, footer string) {
	cmd := exec.Command("git", "commit", "-m", subject, "-m", body, "-m", footer)
	stdout, err := cmd.Output()

	if err != nil {
		printError(fmt.Sprintf("%s", err))
		os.Exit(1)
	}

	fmt.Println(string(stdout))
}

func addFiles() {
	cmd := exec.Command("git", "add", ".")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(stdout))
}

func usage() {
	printHeader("Description:")
	fmt.Println("\t commitzen provides an interface for using conventional commits for your code.")
	printHeader("Usage:")
	fmt.Println("\t commitzen -a -p -c 'home/user/config.json'")
	printHeader("Args:")
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
	var c, prefix string

	flag.BoolVar(&a, "a", false, "Boolean value refering to stage all files")
	flag.BoolVar(&p, "p", false, "Boolean value allowing you to add custom type and scope skipping choises")
	flag.StringVar(&c, "c", "", "String value refering to path of the config file for conventional commits")
	flag.Usage = usage
	flag.Parse()
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

	typeEntered := createMessage(config.Type, p, "type", *reader)
	scopeEntered := createMessage(config.Scope, p, "scope", *reader)
	if scopeEntered != "" {
		scopeEntered = fmt.Sprintf("(%s)", scopeEntered)
	}
	descriptionEntered := createMessage(config.Description, p, "description", *reader)
	bodyEntered := createMessage(config.Body, p, "body", *reader)
	footerEntered := createMessage(config.Footer, p, "footer", *reader)

	printDescrition("Enter any breaking change description, press enter to skip")
	if s := readInput(reader, 0); s != "" {
		breakingChange := fmt.Sprintf("%s%s", "BREAKING CHANGE: ", s)
		footerEntered = fmt.Sprintf("%s\n%s", breakingChange, footerEntered)
		prefix = "!"
	}

	subject := typeEntered + scopeEntered + prefix + ": " + descriptionEntered

	if a {
		addFiles()
	}
	createCommit(subject, bodyEntered, footerEntered)

}

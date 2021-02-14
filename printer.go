package main

import "fmt"

var colorReset string = "\033[0m"
var colorRed string = "\033[31m"
var colorGreen string = "\033[32m"
var colorYellow string = "\033[33m"
var colorBlue string = "\033[34m"
var colorPurple string = "\033[35m"
var colorCyan string = "\033[36m"
var colorWhite string = "\033[37m"

func printHeader(str string) {
	fmt.Println(string(colorCyan), str, string(colorReset))
}

func printDescrition(str string) {
	fmt.Println(string(colorGreen), str, string(colorReset))
}

func printInput(str string) {
	fmt.Println(string(colorPurple), str, string(colorReset))
}

func printError(str string) {
	fmt.Println(string(colorRed), str, string(colorReset))
}

func printSkipping(str string) {
	fmt.Println(string(colorYellow), str, string(colorReset))
}

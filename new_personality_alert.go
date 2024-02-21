package main

import (
	"fmt"
	"os"
)

const VERSION = "v0.1.0"

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--version" ||
			os.Args[1] == "-v" {
			fmt.Println(VERSION)
		} else if os.Args[1] == "--help" ||
			os.Args[1] == "-h" {
			fmt.Println(
				"-v, --version\tshow the app version",
			)
		} else {
			fmt.Println("unknown option ", os.Args[1])
			fmt.Println("try --help for more information")
		}
		return
	}
}

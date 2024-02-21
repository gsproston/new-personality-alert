package main

import (
	"fmt"
	"os"

	"github.com/gsproston/new-personality-alert/tmdb"
)

const VERSION = "v0.1.0"

func handleArgs(args []string) {
	if args[1] == "--version" ||
		args[1] == "-v" {
		fmt.Println(VERSION)
	} else if args[1] == "--help" ||
		args[1] == "-h" {
		fmt.Println(
			"-v, --version\tshow the app version",
		)
	} else {
		fmt.Println("unknown option ", args[1])
		fmt.Println("try --help for more information")
	}
}

func main() {
	if len(os.Args) > 1 {
		handleArgs(os.Args)
		return
	}

	tmdb.GetRyanGoslingMovies()
}

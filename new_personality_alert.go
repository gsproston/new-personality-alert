package main

import (
	"fmt"
	"os"
	"sort"
	"time"

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

	movies := tmdb.GetRyanGoslingMovies()
	if movies != nil {
		sort.Slice(movies, func(i, j int) bool {
			iRelease, _ := time.Parse("2006-01-02", movies[i].Release_Date)
			jRelease, _ := time.Parse("2006-01-02", movies[j].Release_Date)
			return iRelease.After(jRelease)
		})
		fmt.Println("Latest movie: ", movies[0])
	}
}

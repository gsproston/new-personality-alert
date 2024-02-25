package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gsproston/new-personality-alert/email"
	"github.com/gsproston/new-personality-alert/tmdb"
)

const VERSION = "v1.0.0"

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

func getMovieReleasedYesterday(movies []tmdb.TmdbMovie) *tmdb.TmdbMovie {
	// get yesterday's date
	yesterday := time.Now().UTC().Add(time.Hour * -24).Format("2006-01-02")
	fmt.Println("Looking for movies released on", yesterday)

	// see if any films were released yesterday
	for _, movie := range movies {
		if movie.Release_Date == yesterday &&
			// exclude self credits
			!strings.Contains(strings.ToLower(movie.Character), "self") {
			return &movie
		}
	}

	return nil
}

func main() {
	if len(os.Args) > 1 {
		handleArgs(os.Args)
		return
	}

	fmt.Println("Looking for movies...")
	movies := tmdb.GetRyanGoslingMovies()
	fmt.Println("Found movies:\n", movies)

	if movies != nil {
		movie := getMovieReleasedYesterday(movies)

		if movie != nil {
			fmt.Println("New movie found!")
			email.SendAlert(movie.Title, movie.Overview, movie.Character)
		} else {
			fmt.Println("No new movie :(")
		}
	}
}

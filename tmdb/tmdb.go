package tmdb

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetRyanGoslingMovies() {
	client := http.DefaultClient
	req, err := http.NewRequest(
		"GET",
		"https://api.themoviedb.org/3/person/30614-ryan-gosling/movie_credits",
		nil)
	if err != nil {
		log.Fatal((err))
	}

	req.Header.Add("Authorization", "Bearer "+TMDB_TOKEN)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal((err))
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(bodyBytes))
	} else {
		log.Fatal("Unexpected status code: ", resp.StatusCode)
	}
}

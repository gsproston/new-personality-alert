package tmdb

import (
	"encoding/json"
	"log"
	"net/http"
)

type TmdbMovie struct {
	Overview     string
	Release_Date string
	Title        string
	Character    string
}

type TmdbMovieCreditsResponse struct {
	Cast []TmdbMovie
}

func GetRyanGoslingMovies() []TmdbMovie {
	req, err := http.NewRequest(
		"GET",
		"https://api.themoviedb.org/3/person/30614-ryan-gosling/movie_credits",
		nil)
	if err != nil {
		log.Fatal((err))
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+TMDB_TOKEN)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal((err))
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var credits TmdbMovieCreditsResponse
		dec := json.NewDecoder(resp.Body)
		err := dec.Decode(&credits)
		if err != nil {
			log.Fatal(err)
		}
		return credits.Cast
	} else {
		log.Fatal("Unexpected status code: ", resp.StatusCode)
	}

	return nil
}

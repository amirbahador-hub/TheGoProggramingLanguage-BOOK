package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type Movie struct {
	Title   string
	Year    string
	Runtime string
	Poster  string
}

const APIURL = "http://www.omdbapi.com/?"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (m Movie) createFileName() string {
	format := filepath.Ext(m.Poster)
	return fmt.Sprintf("%s%s%s%s", m.Title, m.Year, randSeq(10), format)
}

func (m Movie) createFileMoviePoster() (int, error) {
	fmt.Println(m)
	response, err := http.Get(m.Poster)
	if err != nil {
		return -1, errors.New("can't get this url")
	}
	defer response.Body.Close()
	file, err := os.Create(m.createFileName())
	if err != nil {
		return -1, errors.New("can't create file")
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return -1, errors.New("can't read file")
	}
	return 0, nil
}

func createMovieModel(url string) (movie Movie, err error) {
	response, err := http.Get(url)
	if err != nil {
		return Movie{}, errors.New("can't get this url")
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&movie)
	return
}

func main() {
	url := fmt.Sprintf("%st=%s", APIURL, url.QueryEscape(os.Args[1]))
	movie, err := createMovieModel(url)
	if err != nil {
		os.Exit(1)
	}
	_, err = movie.createFileMoviePoster()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

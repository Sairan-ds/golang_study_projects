package main

import (
	"github.com/Sairan-ds/golang_study_projects/urlShortener/internal/database"
	"github.com/Sairan-ds/golang_study_projects/urlShortener/internal/shortener"
)

func main() {
	database.StartDb()
	database.StartDb()
	shortener.StartShortener()
}

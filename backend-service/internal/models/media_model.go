package models

import (
	"time"
)

type MediaType string

const (
	MovieType  MediaType = "Movie"
	AnimeType  MediaType = "Anime"
	MangaType  MediaType = "Manga"
	ManhwaType MediaType = "Manhwa"
	BookType   MediaType = "Book"
	TVShowType MediaType = "TV Show"
)

type Rating struct {
	Value float64
}

type Media struct {
	ReleaseDate time.Time
	ID          string
	Title       string
	Type        MediaType
	Genre       []string
}

type Movie struct {
	Director string
	Media
	BoxOffice   float64
	IMDbRating  Rating
	SeucyRating Rating
	RunningTime int
}

type Anime struct {
	Studio      string
	WatchStatus string
	Media
	Episodes    int
	Seasons     int
	IMDbRating  Rating
	SeucyRating Rating
}

type Manga struct {
	Author string
	Media
	Chapters    int
	Volumes     int
	IMDbRating  Rating
	SeucyRating Rating
	IsOngoing   bool
}

type Book struct {
	Author    string
	Publisher string
	ISBN      string
	Media
	Pages       int
	IMDbRating  Rating
	SeucyRating Rating
}

type TVShow struct {
	Network     string
	WatchStatus string
	Media
	Seasons     int
	Episodes    int
	IMDbRating  Rating
	SeucyRating Rating
}

func createDummyData() {
	movie := Movie{
		Media: Media{
			ID:          "1",
			Title:       "Inception",
			ReleaseDate: time.Date(2010, time.July, 8, 0, 0, 0, 0, time.UTC),
			Genre:       []string{"Sci-Fi", "Action", "Thriller"},
			Type:        MovieType,
		},
		Director:    "Christopher Nolan",
		BoxOffice:   829895144,
		IMDbRating:  Rating{Value: 8.8},
		SeucyRating: Rating{Value: 9.5},
		RunningTime: 148,
	}
	anime := Anime{
		Media: Media{
			ID:          "2",
			Title:       "One Punch Man",
			ReleaseDate: time.Date(2015, time.October, 5, 0, 0, 0, 0, time.UTC),
			Genre:       []string{"Action", "Comedy", "Superhero"},
			Type:        AnimeType,
		},
		Episodes:    24,
		Studio:      "Madhouse",
		Seasons:     2,
		WatchStatus: "Watching",
		IMDbRating:  Rating{Value: 8.5},
		SeucyRating: Rating{Value: 9.0},
	}
	_ = movie
	_ = anime
}

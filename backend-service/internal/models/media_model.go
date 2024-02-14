package models

import "database/sql"

// TODO: Update table and model acc to change in requirements
type Movie struct {
	ReleaseDate    sql.NullTime
	Title          string
	Genre          string
	SecurityRating string
	AverageRating  sql.NullFloat64
	PersonalRating sql.NullFloat64
	ID             int
	RuntimeMinutes int
	TMDBID         int
}

type Anime struct {
	ReleaseDate    sql.NullTime
	Title          string
	Genre          string
	SecurityRating string
	AverageRating  sql.NullFloat64
	PersonalRating sql.NullFloat64
	ID             int
	EpisodeCount   int
	TMDBID         int
}

type TVShow struct {
	ReleaseDate    sql.NullTime
	Title          string
	Genre          string
	SecurityRating string
	AverageRating  sql.NullFloat64
	PersonalRating sql.NullFloat64
	ID             int
	EpisodeCount   int
	TMDBID         int
}

type Manga struct {
	ReleaseDate    sql.NullTime
	Title          string
	Genre          string
	SecurityRating string
	AverageRating  sql.NullFloat64
	PersonalRating sql.NullFloat64
	ID             int
	ChapterCount   int
}

type Manhua struct {
	ReleaseDate    sql.NullTime
	Title          string
	Genre          string
	SecurityRating string
	AverageRating  sql.NullFloat64
	PersonalRating sql.NullFloat64
	ID             int
	ChapterCount   int
}

type Book struct {
	ReleaseDate    sql.NullTime
	Title          string
	Genre          string
	BookType       string
	SecurityRating string
	AverageRating  sql.NullFloat64
	PersonalRating sql.NullFloat64
	ID             int
}

type UserPreference struct {
	UserID           string `json:"userid"`
	MediaID          string `json:"mediaid"`
	Review           string `json:"review"`
	MediaType        string `json:"mediatype"`
	UserPreferenceID int    `json:"userpreferenceid"`
	Rating           int    `json:"rating"`
}

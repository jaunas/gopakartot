package main

type BaseResponse struct {
	Kinas        int    `json:"kinas"`
	UserLogin    int    `json:"user_login"`
	CodeConfirm  int    `json:"code_confirm"`
	CheckPhone   int    `json:"checkphone"`
	ErrorMessage string `json:"error_message"`
}

type GenreRaw struct {
	Id         string `json:"genre_id"`
	Name       string `json:"genre_name"`
	Order      string `json:"genre_order"`
	Homepage   string `json:"genre_homepage"`
	AlbumPhoto string `json:"album_photo"`
	Cover      string `json:"genre_cover"`
}

type GenreResponse struct {
	*BaseResponse
	Genres []GenreRaw `json:"genres"`
}

type TrackRaw struct {
	LikeCount     interface{} `json:"like_count"`
	LikeState     string      `json:"like_state"`
	FullPermalink string      `json:"track_full_permalink"`
	Id            string      `json:"tid"`
	Length        string      `json:"length"`
	Title         string      `json:"title"`
	Artist        string      `json:"artist"`
	Filename      string      `json:"filename"`
}

type AlbumTrackRaw struct {
	*TrackRaw
	AlbumTrackId string `json:"album_track_id"`
	AlbumId      string `json:"album_track__album_id"`
	Order        string `json:"album_track_order"`
	IsDisabled   string `json:"album_track_is_disabled"`
	Id           string `json:"track_id"`
	Title        string `json:"track_name"`
	Length       string `json:"track_length"`
	Year         string `json:"track_year"`
	Permalink    string `json:"track_permalink"`
	Artist       string `json:"performers"`
}

type AlbumFilesResponse struct {
	*BaseResponse
	Tracks []TrackRaw `json:"tracks"`
}

type AlbumInfoResponse struct {
	*BaseResponse
	Legal       bool            `json:"legal"`
	TotalLength string          `json:"total_length"`
	Tracks      []AlbumTrackRaw `json:"tracks"`
}

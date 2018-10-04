package main

type BaseResponse struct {
	Kinas        int    `json:"kinas"`
	UserLogin    int    `json:"user_login"`
	CodeConfirm  int    `json:"code_confirm"`
	CheckPhone   int    `json:"checkphone"`
	ErrorMessage string `json:"error_message"`
}

type Genre struct {
	Id         string `json:"genre_id"`
	Name       string `json:"genre_name"`
	Order      string `json:"genre_order"`
	Homepage   string `json:"genre_homepage"`
	AlbumPhoto string `json:"album_photo"`
	Cover      string `json:"genre_cover"`
}

type GenreResponse struct {
	*BaseResponse
	Genres []Genre `json:"genres"`
}

type Like struct {
	Count string `json:"count"`
	State string `json:"state"`
}

type Album struct {
	Id            string `json:"album_id"`
	Name          string `json:"album_name"`
	Year          string `json:"album_year"`
	Photo         string `json:"album_photo"`
	Permalink     string `json:"album_permalink"`
	NameGenerated string `json:"album_name_generated"`
	Performers    string `json:"performers"`
	PhotoPath     string `json:"photo_path"`
	Like          Like   `json:"like"`
}

type AlbumWithLikeTotalCount struct {
	*Album
	LikeTotalCount string `json:"like_total_count"`
}

type NewMusicAlbumResponse struct {
	*BaseResponse
	Action string  `json:"action"`
	Albums []Album `json:"new_music_albums"`
}

type NewestAlbumResponse struct {
	*BaseResponse
	Action string  `json:"action"`
	Albums []Album `json:"newest_albums"`
}

type MostLikedAlbumResponse struct {
	*BaseResponse
	Action string                    `json:"action"`
	Albums []AlbumWithLikeTotalCount `json:"most_liked_albums"`
}

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

type LikeRaw struct {
	Count interface{} `json:"count"`
	State string      `json:"state"`
}

type AlbumRaw struct {
	Id               string      `json:"album_id"`
	UserId           string      `json:"album__user_id"`
	LabelId          string      `json:"album__label_id"`
	LabelFreetext    string      `json:"album_label_freetext"`
	ProjectId        string      `json:"album__project_id"`
	AliasId          string      `json:"album__alias_id"`
	GenreId          string      `json:"album__genre_id"`
	AgataId          string      `json:"album_agata_id"`
	Year             string      `json:"album_year"`
	Name             string      `json:"album_name"`
	NameGenerated    string      `json:"album_name_generated"`
	Photo            string      `json:"album_photo"`
	PhotoOffsetX     string      `json:"album_photo_offset_x"`
	PhotoOffsetY     string      `json:"album_photo_offset_y"`
	PublishTime      string      `json:"album_publish_time"`
	Permalink        string      `json:"album_permalink"`
	IsDisabled       string      `json:"album_is_disabled"`
	IsLocked         string      `json:"album_is_locked"`
	State            string      `json:"album_state"`
	Timestamp        string      `json:"album_timestamp"`
	Main             interface{} `json:"album_main"`
	MainOrder        interface{} `json:"album_main_order"`
	IsAdmin          string      `json:"album_is_admin"`
	IsEditableByUser string      `json:"album_is_editable_by_user"`
	FullInfo         string      `json:"album_full_info"`
	UploadBy         string      `json:"album_upload_by"`
	EditBy           string      `json:"album_edit_by"`
	UpdateTime       string      `json:"album_update_time"`
	Genre            string      `json:"genre"`
	Performers       string      `json:"performers"`
	PhotoPath        string      `json:"photo_path"`
	Like             LikeRaw     `json:"like"`
}

type AlbumWithLikeTotalCountRaw struct {
	*AlbumRaw
	LikeTotalCount string `json:"like_total_count"`
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

type NewMusicAlbumResponse struct {
	*BaseResponse
	Action string     `json:"action"`
	Albums []AlbumRaw `json:"new_music_albums"`
}

type NewestAlbumResponse struct {
	*BaseResponse
	Action string     `json:"action"`
	Albums []AlbumRaw `json:"newest_albums"`
}

type MostLikedAlbumResponse struct {
	*BaseResponse
	Action string                       `json:"action"`
	Albums []AlbumWithLikeTotalCountRaw `json:"most_liked_albums"`
}

type GenreAlbumResponse struct {
	*BaseResponse
	Albums    []AlbumRaw `json:"albums"`
	GenreName string     `json:"genre_name"`
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

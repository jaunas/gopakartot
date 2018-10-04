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
	Count interface{} `json:"count"`
	State string      `json:"state"`
}

type Album struct {
	Id               string      `json:"album_id"`
	UserId           string      `json:"album__user_id"`
	LabelId          string      `json:"album__label_id"`
	LabelFreetext    interface{} `json:"album_label_freetext"`
	ProjectId        interface{} `json:"album__project_id"`
	AliasId          interface{} `json:"album__alias_id"`
	GenreId          interface{} `json:"album__genre_id"`
	AgataId          interface{} `json:"album_agata_id"`
	Year             string      `json:"album_year"`
	Name             string      `json:"album_name"`
	NameGenerated    string      `json:"album_name_generated"`
	Photo            string      `json:"album_photo"`
	PhotoOffsetX     string      `json:"album_photo_offset_x"`
	PhotoOffsetY     string      `json:"album_photo_offset_y"`
	PublishTime      interface{} `json:"album_publish_time"`
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
	Like             Like        `json:"like"`
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

type GenreAlbumResponse struct {
	*BaseResponse
	Albums    []Album `json:"albums"`
	GenreName string  `json:"genre_name"`
}

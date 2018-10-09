package main

import (
	"strconv"
	"time"
)

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

type Album struct {
	Id     int
	UserId int
	Label  struct {
		Id       int
		Freetext string
	}
	ProjectId        int
	AliasId          int
	GenreId          int
	AgataId          int
	Year             int
	Name             string
	NameGenerated    string
	Photo            *Photo
	PublishTime      time.Time
	Permalink        string
	IsDisabled       bool
	IsLocked         bool
	State            string
	Timestamp        time.Time
	Main             interface{}
	MainOrder        interface{}
	IsAdmin          bool
	IsEditableByUser bool
	FullInfo         string
	UploadBy         int
	EditBy           int
	UpdateTime       time.Time
	Genre            string
	Performers       string
	Like             *Like
}

func (raw AlbumRaw) Parse() (*Album, error) {
	album := &Album{}

	album.Id, _ = strconv.Atoi(raw.Id)
	album.UserId, _ = strconv.Atoi(raw.UserId)
	album.ProjectId, _ = strconv.Atoi(raw.ProjectId)
	album.AliasId, _ = strconv.Atoi(raw.AliasId)
	album.GenreId, _ = strconv.Atoi(raw.GenreId)
	album.AgataId, _ = strconv.Atoi(raw.AgataId)
	album.Year, _ = strconv.Atoi(raw.Year)
	album.Name = raw.Name
	album.NameGenerated = raw.NameGenerated
	album.Photo, _ = raw.ParsePhoto()
	album.Permalink = raw.Permalink
	album.State = raw.State
	album.FullInfo = raw.FullInfo
	album.UploadBy, _ = strconv.Atoi(raw.UploadBy)
	album.EditBy, _ = strconv.Atoi(raw.EditBy)
	album.Genre = raw.Genre
	album.Performers = raw.Performers
	album.Like, _ = raw.Like.Parse()

	// Label  struct {
	// 	Id       int
	// 	Freetext string
	// }
	labelId, _ := strconv.Atoi(raw.LabelId)
	album.Label = Label{
		Id:       labelId,
		Freetext: raw.LabelFreetext,
	}

	// PublishTime      time.Time
	if len(raw.PublishTime) > 0 {
		time, err := time.Parse("2006-01-02 15:04:05", raw.PublishTime)
		if err != nil {
			return nil, err
		}
		album.PublishTime = time
	}

	// IsDisabled       bool
	if raw.IsDisabled == "N" {
		album.IsDisabled = false
	} else if raw.IsDisabled == "Y" {
		album.IsDisabled = true
	}

	// IsLocked         bool
	if raw.IsLocked == "0" {
		album.IsLocked = false
	} else if raw.IsLocked == "1" {
		album.IsLocked = true
	}

	// Timestamp        time.Time
	if len(raw.Timestamp) > 0 {
		time, err := time.Parse("2006-01-02 15:04:05", raw.Timestamp)
		if err != nil {
			return nil, err
		}
		album.Timestamp = time
	}

	// Main             interface{}
	if raw.Main != nil {
		panic(raw.Main)
	}

	// MainOrder        interface{}
	if raw.MainOrder != nil {
		panic(raw.MainOrder)
	}

	// IsAdmin          bool
	if raw.IsAdmin == "N" {
		album.IsAdmin = false
	} else if raw.IsAdmin == "Y" {
		album.IsAdmin = true
	}

	// IsEditableByUser bool
	if raw.IsEditableByUser == "N" {
		album.IsEditableByUser = false
	} else if raw.IsEditableByUser == "Y" {
		album.IsEditableByUser = true
	}

	// UpdateTime       time.Time
	if len(raw.UpdateTime) > 0 {
		time, err := time.Parse("2006-01-02 15:04:05", raw.UpdateTime)
		if err != nil {
			return nil, err
		}
		album.UpdateTime = time
	}

	return album, nil
}

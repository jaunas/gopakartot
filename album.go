package main

import (
	"strconv"
	"time"
)

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

func CreateAlbumFromRaw(raw AlbumRaw) (*Album, error) {
	album := &Album{}

	// Id     int
	album.Id, _ = strconv.Atoi(raw.Id)

	// UserId int
	album.UserId, _ = strconv.Atoi(raw.UserId)

	// Label  struct {
	// 	Id       int
	// 	Freetext string
	// }
	labelId, _ := strconv.Atoi(raw.LabelId)
	album.Label = Label{
		Id:       labelId,
		Freetext: raw.LabelFreetext,
	}

	// ProjectId        int
	album.ProjectId, _ = strconv.Atoi(raw.ProjectId)

	// AliasId          int
	album.AliasId, _ = strconv.Atoi(raw.AliasId)

	// GenreId          int
	album.GenreId, _ = strconv.Atoi(raw.GenreId)

	// AgataId          int
	album.AgataId, _ = strconv.Atoi(raw.AgataId)

	// Year             int
	album.Year, _ = strconv.Atoi(raw.Year)

	// Name             string
	album.Name = raw.Name

	// NameGenerated    string
	album.NameGenerated = raw.NameGenerated

	// Photo            Photo
	album.Photo, _ = raw.ExtractPhoto()

	// PublishTime      time.Time
	if len(raw.PublishTime) > 0 {
		time, err := time.Parse("2006-01-02 15:04:05", raw.PublishTime)
		if err != nil {
			return nil, err
		}
		album.PublishTime = time
	}

	// Permalink        string
	album.Permalink = raw.Permalink

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

	// State            string
	album.State = raw.State

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

	// FullInfo         string
	album.FullInfo = raw.FullInfo

	// UploadBy         int
	album.UploadBy, _ = strconv.Atoi(raw.UploadBy)

	// EditBy           int
	album.EditBy, _ = strconv.Atoi(raw.EditBy)

	// UpdateTime       time.Time
	if len(raw.UpdateTime) > 0 {
		time, err := time.Parse("2006-01-02 15:04:05", raw.UpdateTime)
		if err != nil {
			return nil, err
		}
		album.UpdateTime = time
	}

	// Genre            string
	album.Genre = raw.Genre

	// Performers       string
	album.Performers = raw.Performers

	// Like             Like
	album.Like, _ = CreateLikeFromRaw(raw.Like)

	return album, nil
}

package main

import "strconv"

type Photo struct {
	Photo  string
	Offset struct {
		X int
		Y int
	}
	Path string
}

func (raw *AlbumRaw) ExtractPhoto() (*Photo, error) {
	photoOffsetX, _ := strconv.Atoi(raw.PhotoOffsetX)
	photoOffsetY, _ := strconv.Atoi(raw.PhotoOffsetY)

	return &Photo{
		Photo: raw.Photo,
		Offset: struct {
			X int
			Y int
		}{
			X: photoOffsetX,
			Y: photoOffsetY,
		},
		Path: raw.PhotoPath,
	}, nil
}

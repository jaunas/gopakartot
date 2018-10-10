package main

type AlbumsResponse struct {
	*BaseResponse
	Albums []AlbumRaw `json:"albums"`
}

type NewMusicAlbumResponse struct {
	*AlbumsResponse
	Albums []AlbumRaw `json:"new_music_albums"`
	Action string     `json:"action"`
}

type MostLikedAlbumResponse struct {
	*AlbumsResponse
	Albums []AlbumRaw `json:"most_liked_albums"`
	Action string     `json:"action"`
}

type NewestAlbumResponse struct {
	*AlbumsResponse
	Albums []AlbumRaw `json:"newest_albums"`
	Action string     `json:"action"`
}

type GenreAlbumResponse struct {
	*AlbumsResponse
	GenreName string `json:"genre_name"`
}

func (response AlbumsResponse) GetAlbums() ([]*Album, error) {
	var albums []*Album
	for _, rawAlbum := range response.Albums {
		album, err := rawAlbum.Parse()
		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}

	return albums, nil
}

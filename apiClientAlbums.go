package main

import (
	"encoding/json"
	"errors"
)

func (apiClient *ApiClient) getNewMusicAlbums(page int) (*NewMusicAlbumResponse, error) {
	response, err := apiClient.request(map[string]string{
		"action": "new_music_albums",
		"url":    "home",
		"page":   string(page),
	})

	if err != nil {
		return nil, err
	}

	albumResponse := &NewMusicAlbumResponse{}
	json.Unmarshal(response, &albumResponse)

	if len(albumResponse.ErrorMessage) > 0 {
		return nil, errors.New(albumResponse.ErrorMessage)
	}

	return albumResponse, nil
}

func (apiClient *ApiClient) getNewestAlbums(page int) (*NewestAlbumResponse, error) {
	response, err := apiClient.request(map[string]string{
		"action": "newest_albums",
		"url":    "home",
		"page":   string(page),
	})

	if err != nil {
		return nil, err
	}

	albumResponse := &NewestAlbumResponse{}
	json.Unmarshal(response, albumResponse)

	if len(albumResponse.ErrorMessage) > 0 {
		return nil, errors.New(albumResponse.ErrorMessage)
	}

	return albumResponse, nil
}

func (apiClient *ApiClient) getMostLikedAlbums(page int) (*MostLikedAlbumResponse, error) {
	response, err := apiClient.request(map[string]string{
		"action": "most_liked_albums",
		"url":    "home",
		"page":   string(page),
	})

	if err != nil {
		return nil, err
	}

	albumResponse := &MostLikedAlbumResponse{}
	json.Unmarshal(response, albumResponse)

	if len(albumResponse.ErrorMessage) > 0 {
		return nil, errors.New(albumResponse.ErrorMessage)
	}

	return albumResponse, nil
}

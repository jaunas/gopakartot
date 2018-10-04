package main

import (
	"encoding/json"
)

func (apiClient *ApiClient) getNewMusicAlbums(page int) NewMusicAlbumResponse {
	response := apiClient.request(map[string]string{
		"action": "new_music_albums",
		"url":    "home",
		"page":   string(page),
	})

	var albumResponse NewMusicAlbumResponse
	json.Unmarshal(response, &albumResponse)

	return albumResponse
}

func (apiClient *ApiClient) getNewestAlbums(page int) NewestAlbumResponse {
	response := apiClient.request(map[string]string{
		"action": "newest_albums",
		"url":    "home",
		"page":   string(page),
	})

	var albumResponse NewestAlbumResponse
	json.Unmarshal(response, &albumResponse)

	return albumResponse
}

func (apiClient *ApiClient) getMostLikedAlbums(page int) MostLikedAlbumResponse {
	response := apiClient.request(map[string]string{
		"action": "most_liked_albums",
		"url":    "home",
		"page":   string(page),
	})

	var albumResponse MostLikedAlbumResponse
	json.Unmarshal(response, &albumResponse)

	return albumResponse
}

package main

import (
	"encoding/json"
	"log"
)

func (apiClient *ApiClient) getNewMusicAlbums(page int) NewMusicAlbumResponse {
	response, err := apiClient.request(map[string]string{
		"action": "new_music_albums",
		"url":    "home",
		"page":   string(page),
	})

	if err != nil {
		log.Fatal(err)
	}

	var albumResponse NewMusicAlbumResponse
	json.Unmarshal(response, &albumResponse)

	return albumResponse
}

func (apiClient *ApiClient) getNewestAlbums(page int) NewestAlbumResponse {
	response, err := apiClient.request(map[string]string{
		"action": "newest_albums",
		"url":    "home",
		"page":   string(page),
	})

	if err != nil {
		log.Fatal(err)
	}

	var albumResponse NewestAlbumResponse
	json.Unmarshal(response, &albumResponse)

	return albumResponse
}

func (apiClient *ApiClient) getMostLikedAlbums(page int) MostLikedAlbumResponse {
	response, err := apiClient.request(map[string]string{
		"action": "most_liked_albums",
		"url":    "home",
		"page":   string(page),
	})

	if err != nil {
		log.Fatal(err)
	}

	var albumResponse MostLikedAlbumResponse
	json.Unmarshal(response, &albumResponse)

	return albumResponse
}

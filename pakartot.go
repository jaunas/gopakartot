package main

import (
	"fmt"
	"log"
)

func main() {
	apiClient := NewApiClient()

	albumsResponse, err := apiClient.getNewMusicAlbums(1)

	if err != nil {
		log.Fatal(err)
	}

	for _, album := range albumsResponse {
		fmt.Printf("%+v\n", album)
	}
}

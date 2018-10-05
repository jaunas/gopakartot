package main

import (
	"fmt"
	"log"
)

func main() {
	apiClient := NewApiClient()

	albumsResponse, err := apiClient.getAlbumFiles(10788)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", albumsResponse.Tracks)
}

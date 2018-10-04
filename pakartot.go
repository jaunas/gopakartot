package main

import (
	"fmt"
	"log"
)

func main() {
	apiClient := NewApiClient()

	albumsResponse, err := apiClient.getGenreAlbums(1, 2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", albumsResponse.Albums)
}

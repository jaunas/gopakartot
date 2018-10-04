package main

import (
	"fmt"
	"log"
)

func main() {
	apiClient := NewApiClient()

	albumsResponse, err := apiClient.getMostLikedAlbums(0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", albumsResponse)
}

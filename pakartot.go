package main

import (
	"fmt"
	"log"
)

func main() {
	apiClient := NewApiClient()

	albumsResponse, err := apiClient.getMostLikedAlbums(2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", albumsResponse)
}

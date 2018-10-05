package main

import (
	"fmt"
	"log"
)

func main() {
	apiClient := NewApiClient()

	albumsResponse, err := apiClient.getAlbumInfo(10788)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", albumsResponse)
}

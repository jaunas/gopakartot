package main

import "fmt"

func main() {
	apiClient := NewApiClient()
	// request := apiClient.getGengres()
	//
	// for _, genre := range request {
	// 	fmt.Println(genre.Id, genre.Name)
	// }

	fmt.Printf("%+v\n", apiClient.getMostLikedAlbums(0))
}

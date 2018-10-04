package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type ApiClient struct {
	username string
	password string
	url      string
}

func NewApiClient() *ApiClient {
	apiClient := &ApiClient{
		username: "publicUSR",
		password: "vka3qaGHowofKcRdTeiV",
		url:      "https://api.pakartot.lt/",
	}

	return apiClient
}

func (apiClient *ApiClient) request(parameters map[string]string) []byte {
	values := url.Values{
		"username": {apiClient.username},
		"password": {apiClient.password},
	}

	for key, value := range parameters {
		values.Add(key, value)
	}

	resp, err := http.PostForm(apiClient.url, values)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func (apiClient *ApiClient) getGengres() []Genre {
	response := apiClient.request(map[string]string{
		"url": "genres",
	})

	var genreResponse GenreResponse
	json.Unmarshal(response, &genreResponse)

	return genreResponse.Genres
}

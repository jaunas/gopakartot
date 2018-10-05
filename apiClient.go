package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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

func (apiClient *ApiClient) request(parameters map[string]string) ([]byte, error) {
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

	if json.Valid(body) {
		return body, nil
	}

	lines := strings.Split(string(body), "\n")

	for _, line := range lines {
		if json.Valid([]byte(line)) {
			return []byte(line), nil
		}
	}

	return nil, errors.New("Response is not valid JSON.")
}

func (apiClient *ApiClient) getGengres() []GenreRaw {
	response, err := apiClient.request(map[string]string{
		"url": "genres",
	})

	if err != nil {
		log.Fatal(err)
	}

	var genreResponse GenreResponse
	json.Unmarshal(response, &genreResponse)

	return genreResponse.Genres
}

package main

import "strconv"

type LikeRaw struct {
	Count interface{} `json:"count"`
	State string      `json:"state"`
}

type Like struct {
	Count int
	State bool
}

func (raw LikeRaw) Parse() (*Like, error) {
	likeCount := 0
	v, ok := raw.Count.(string)
	if ok {
		likeCount, _ = strconv.Atoi(v)
	}

	likeState := false
	if raw.State == "on" {
		likeState = true
	}

	return &Like{
		Count: likeCount,
		State: likeState,
	}, nil
}

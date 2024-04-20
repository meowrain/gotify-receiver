package client

import "gotify-client/pkg/utils/time"

type GotifyMessage struct {
	Id       int            `json:"id"`
	Appid    int            `json:"appid"`
	Message  string         `json:"message"`
	Title    string         `json:"title"`
	Priority int            `json:"priority"`
	Extras   map[string]any `json:"extras"`
	Date     time.Time      `json:"date"`
}

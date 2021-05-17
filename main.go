package main

import (
	"homeserver/services/radarr"
)

func main() {
	r := radarr.New("127.0.0.1", 7878, "APIKEY", "http")
	r.GetMovies()
}

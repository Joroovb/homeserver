package radarr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type radarr struct {
	host     string
	port     int
	key      string
	protocol string
}

func New(host string, port int, key string, protocol string) radarr {
	return radarr{
		host,
		port,
		key,
		protocol,
	}
}

// MOVIE

func (r radarr) GetMovies() []Movie {
	uri := fmt.Sprintf("%s://%s:%d/api/v3/movie?apikey=%s", r.protocol, r.host, r.port, r.key)
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var movie []Movie
	err = json.Unmarshal(body, &movie)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Print(sb)

	return movie
}

func (r radarr) GetMovieById(id int) Movie {
	uri := fmt.Sprintf("%s://%s:%d/api/v3/movie?tmdbid=%d&apikey=%s", r.protocol, r.host, r.port, id, r.key)
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var movie []Movie
	err = json.Unmarshal(body, &movie)

	if err != nil {
		log.Fatalln(err)
	}
	return movie[0]
}

func (r radarr) LookupMovie(query string) {
	q := url.QueryEscape(query)
	uri := fmt.Sprintf("%s://%s:%d/api/v3/movie/lookup?term=%s&apikey=%s", r.protocol, r.host, r.port, q, r.key)
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var movie []Movie
	err = json.Unmarshal(body, &movie)

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Print(sb)
}

func (r radarr) AddMovie() {
	uri := fmt.Sprintf("%s://%s:%d/api/v3/movie?apikey=%s", r.protocol, r.host, r.port, r.key)
	postBody, _ := json.Marshal(map[string]string{})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(uri, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	fmt.Println(sb)
}

func (r radarr) GetHealth() []Health {
	uri := fmt.Sprintf("%s://%s:%d/api/v3/health?apikey=%s", r.protocol, r.host, r.port, r.key)
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var health []Health
	err = json.Unmarshal(body, &health)

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Print(sb)

	return health
}

// INDEXER

func (r radarr) GetAllIndexers() []Indexer {
	uri := fmt.Sprintf("%s://%s:%d/api/v3/indexer?apikey=%s", r.protocol, r.host, r.port, r.key)
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var indexers []Indexer
	err = json.Unmarshal(body, &indexers)

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Print(sb)

	return indexers
}

func (r radarr) GetIndexerById(id int) Indexer {
	uri := fmt.Sprintf("%s://%s:%d/api/v3/indexer/%d?apikey=%s", r.protocol, r.host, r.port, id, r.key)
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var indexer Indexer
	err = json.Unmarshal(body, &indexer)

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Print(sb)

	return indexer
}

func (r radarr) EditIndexerById(id int) {
	uri := fmt.Sprintf("%s://%s:%d/api/v3/indexer/%d?apikey=%s", r.protocol, r.host, r.port, id, r.key)
	_, err := http.NewRequest("PUT", uri, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func (r radarr) DeleteIndexerById(id int) {
	uri := fmt.Sprintf("%s://%s:%d/api/v3/indexer/%d?apikey=%s", r.protocol, r.host, r.port, id, r.key)
	_, err := http.NewRequest("Delete", uri, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// DISKSPACE

func (r radarr) GetDiskSpace() []Disk {
	uri := fmt.Sprintf("%s://%s:%d/api/v3/diskspace?apikey=%s", r.protocol, r.host, r.port, r.key)
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var disks []Disk
	err = json.Unmarshal(body, &disks)

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Print(sb)

	return disks
}

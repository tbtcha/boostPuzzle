package rpc

import (
	"boostPuzzle/server/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	defaultLimit   = 10
	defaultType    = "image"
	defaultLimitBy = "post"
)

type RPC interface {
	GetAlbum(username, offset string) (*models.MediaAlbum, error)
}

type rpc struct {
	url string
	cl  Doer
}

func NewRpc(url string, cl Doer) RPC {
	return rpc{
		cl:  cl,
		url: url,
	}
}

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// GetAlbum ?type=image&limit=100&limit_by=post?offset=%s
func (r rpc) GetAlbum(username, offset string) (*models.MediaAlbum, error) {
	url := fmt.Sprintf("%s/%s/media_album/", r.url, username)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()                                 // Get a copy of the query values.
	q.Add("type", fmt.Sprintf("%s", defaultType))        // Add a new value to the set.
	q.Add("limit", fmt.Sprintf("%d", defaultLimit))      // Add a new value to the set.
	q.Add("limit_by", fmt.Sprintf("%s", defaultLimitBy)) // Add a new value to the set.
	fmt.Println(offset)
	if len(offset) > 0 {
		q.Add("offset", fmt.Sprintf("%s", offset)) // Add a new value to the set.
	}
	req.URL.RawQuery = q.Encode()
	resp, err := r.cl.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(b))
		return nil, errors.New(string(b))
	}

	var res models.MediaAlbum
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		fmt.Println("2", err.Error())
		return nil, err
	}

	return &res, nil
}

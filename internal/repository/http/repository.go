package http

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ClientProvider interface {
	Do(req *http.Request) (*http.Response, error)
}

type repository struct {
	client ClientProvider
}

func New(cli ClientProvider) *repository {
	return &repository{cli}
}

func (repo *repository) Get(ctx context.Context, urlParam string, queryParam map[string]string, header map[string]string, body []byte) ([]byte, error) {
	baseUrl, err := url.Parse(urlParam)
	if err != nil {
		return nil, err
	}

	// Set Query Params
	params := url.Values{}
	for key, value := range queryParam {
		params.Add(key, value)
	}
	baseUrl.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", baseUrl.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set Headers
	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := repo.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return respBody, nil
}

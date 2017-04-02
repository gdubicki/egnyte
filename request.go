package egnyte

import (
	"io"
	"io/ioutil"
	"net/http"
)

func get(token, domain string) ([]byte, error) {
	return makeRequest(token, domain, "GET", nil)
}

func post(token, domain string, payload io.Reader) ([]byte, error) {
	return makeRequest(token, domain, "POST", payload)
}

func makeRequest(token, endpoint, method string, payload io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, endpoint, payload)
	if len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	if payload != nil {
		req.Header.Add("content-type", "application/x-www-form-urlencoded")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

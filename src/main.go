package main

import "C"

import (
	"net/http"
	"io/ioutil"
)

//export Request
func Request(method, url string) (string) {
	req, err := http.NewRequest(method, url, nil);
	if err != nil {
		return "ERROR"
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "ERROR"
	}

	body, err := ioutil.ReadAll(resp.Body);
	if err != nil {
		return "ERROR"
	}

	return string(body)
}

//export Get
func Get(url string) (string) {
	return Request("GET", url)
}

func main() {}
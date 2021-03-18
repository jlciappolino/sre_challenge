package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

//Get basic http get
func get(url string) (response []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)

	return
}

//post basic http post
func post(url string, body string) (response []byte, err error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)

	return
}

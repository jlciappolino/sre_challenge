package main

import (
	"io/ioutil"
	"net/http"
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

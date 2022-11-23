package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("ERROR: %v", err)
	}
	defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// return string(body[:]), nil
	return resp.Header.Get("Content-Type"), nil
}

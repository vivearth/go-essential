package challenge

import (
	"fmt"
	"net/http"
)

func ContentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("ERROR: %v", err)
	}
	defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// return string(body[:]), nil
	return resp.Header.Get("Content-Type"), nil
}

func ContentTypeP(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		// fmt.Printf("ERROR: %v\n", err)
		out <- fmt.Sprintf("%s -> error: %s", url, err)
		return
	}
	defer resp.Body.Close()
	out <- fmt.Sprintf("Response for url %s is %s", url, resp.Header.Get("Content-Type"))
}

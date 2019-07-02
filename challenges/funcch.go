package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	contype := resp.Header.Get("Content-Type")

	if contype == "" {
		return "", fmt.Errorf("cant find Content-Type header")
	}
	return contype, nil
}
func main() {

	ret1, err := contentType("https://golang.org")

	fmt.Println(ret1, err)
}

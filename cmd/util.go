package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
)

func ResponseData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Panicf("error when fetch random quote: %v\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("error when read data random quote: %v\n", err)
	}

	return body
}

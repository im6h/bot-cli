package cmd

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

// common
var name, page, limit string

// github
var json jsoniter.API
var ctx context.Context
var token, description, dns string
var isPrivate, isAutoInit bool

func init() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}

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

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type CompileResult struct {
	Errors    string
	Events    []Event
	VetErrors string
}

type Event struct {
	Message string
	Kind    string
	Delay   int
}

/*
		curl 'https://go.dev/_/compile' \
		  --data-raw $'version=2&body=%2F%2F+You+can+edit+this+code\u0021%0A%2F%2F+Click+here+and+start+typing.%0Apackage+main%0A%0Aimport+%22fmt%22%0A%0Afunc+main()+%7B%0A%09fmt.Println(%22Hello%2C+%E4%B8%96%E7%95%8C%22)%0A%7D%0A&withVet=true'

		curl 'https://go.dev/_/fmt' \
	  --data-raw $'body=%2F%2F+You+can+edit+this+code\u0021%0A%2F%2F+Click+here+and+start+typing.%0Apackage+main%0A%0Aimport+%22fmt%22%0A%0Afunc+main()+%7B%0A%09fmt.Println(%22Hello%2C+%E4%B8%96%E7%95%8C%22)%0A%7D%0A&imports=true'
*/
func compile(input string) (string, error) {
	const url = "https://go.dev/_/compile"
	res, err := post(input, url, &CompileResult{})
	if err != nil {
		return "", err
	}

	if res.Errors != "" {
		return "", errors.New(res.Errors)
	}
	if len(res.Events) == 0 {
		return "", errors.New("no events")
	}
	str := res.Events[0].Message
	return str, nil
}

type FormatResult struct {
	Error string
	Body  string
}

func format(input string) (string, error) {
	const url = "https://go.dev/_/fmt"
	res, err := post(input, url, &FormatResult{})
	if err != nil {
		return "", err
	}
	if res.Error != "" {
		return "", errors.New(res.Error)
	}
	if res.Body == "" {
		return "", errors.New("syntax error")
	}

	return res.Body, nil
}

func post[T any](body, apiURL string, respInput T) (T, error) {
	data := url.Values{}
	data.Set("body", body)
	data.Set("version", "2")
	data.Set("withVet", "true")
	req, err := http.NewRequest(
		http.MethodPost,
		apiURL,
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return respInput, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return respInput, err
	}
	defer resp.Body.Close()

	res := json.NewDecoder(resp.Body)
	err = res.Decode(respInput)
	return respInput, err
}

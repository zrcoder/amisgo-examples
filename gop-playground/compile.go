package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type CompileResult struct {
	Errors     string
	Events     []Event
	Status     int
	IsTest     bool
	TestFailed int
}

type Event struct {
	Message string
	Kind    string
	Delay   int
}

func compile(input string) (string, error) {
	const url = "https://play.goplus.org/compile"
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
	const url = "https://play.goplus.org/fmt"
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

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
	data := url.Values{}
	data.Set("body", input)
	req, err := http.NewRequest(
		http.MethodPost,
		"https://play.goplus.org/compile",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	res := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	cmRes := &CompileResult{}
	err = res.Decode(cmRes)
	if err != nil {
		return "", err
	}

	if cmRes.Errors != "" {
		return "", errors.New(cmRes.Errors)
	}
	if len(cmRes.Events) == 0 {
		return "", errors.New("no events")
	}
	str := cmRes.Events[0].Message
	return str, nil
}

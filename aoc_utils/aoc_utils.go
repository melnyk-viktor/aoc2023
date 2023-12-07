package aoc_utils

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
)

func GetInputData(url string, nd string, session string) string {
	cookie := &http.Cookie{
		Name:  "session",
		Value: session,
	}

	client := &http.Client{
		Jar:       &cookiejar.Jar{},
		Transport: &http.Transport{},
	}

	var iurl string
	if session == "" {
		iurl = url + "/day/" + nd
	} else {
		iurl = url + "/day/" + nd + "/input"
	}
	req, err := http.NewRequest(http.MethodGet, iurl, nil)
	if err != nil {
		fmt.Println("Unable to create request to AoC website", err)
		os.Exit(1)
	}

	if session != "" {
		req.AddCookie(cookie)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to send request to AoC website", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Unable to read response body", err)
		os.Exit(1)
	}

	if session == "" {
		data := strings.Split(strings.Split(string(resBody), "<pre><code>")[1], "</code>")[0]
		data = strings.TrimSpace(data)
		return data
	}

	return strings.TrimSpace(string(resBody))
}

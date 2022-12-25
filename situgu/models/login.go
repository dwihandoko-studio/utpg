package models

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

type TokenApp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredIn    int64  `json:"expired_in"`
}

func GetTokenApp(username, password string) (TokenApp, error) {
	var err error
	var client = &http.Client{}
	var data TokenApp

	var param = url.Values{}
	param.Set("username", username)
	param.Set("password", password)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", `http://10.90.90.1:1990/login`, payload)
	// request, err := http.NewRequest("POST", os.Getenv("URL_BE")+`/login`, payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/json")
	// request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("X-API-TOKEN", os.Getenv("API_TOKEN"))

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

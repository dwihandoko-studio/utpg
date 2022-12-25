package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FormLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenApp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredIn    int64  `json:"expired_in"`
}

func GetTokenApp(username, password string) (TokenApp, error) {
	var err error
	var client = &http.Client{}
	var data TokenApp

	// var param = url.Values{}
	// param.Set("username", username)
	// param.Set("password", password)
	// var payload = bytes.NewBufferString(param.Encode())
	user := FormLogin{
		Username: username,
		Password: password,
	}
	body, _ := json.Marshal(user)
	var payload = bytes.NewBuffer(body)

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

	log.Println("REQUEST LOGIN: ", response.StatusCode)
	r := Response{
		Data: &TokenApp{},
	}

	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return data, err
	}

	if r.Status != 200 {
		return data, errors.New(r.Message)
	}

	token, _ := r.Data.(*TokenApp)
	log.Println(token.AccessToken)

	if token == nil {
		return data, errors.New("Gagal memuat data.")
	}

	return *token, nil
}

package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dwihandoko-studio/utpg/layanan/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Users struct {
	Data      []models.User `json:"data"`
	Total     int           `json:"total"`
	Page      int           `json:"page"`
	Last_page int           `json:"last_page"`
}

type ReqPengguna struct {
	Start   string `json:"start"`
	Length  string `json:"length"`
	Keyword string `json:"keyword"`
	Role    string `json:"role"`
}

func GetUsers(req interface{}, c *gin.Context) (Users, error) {

	var err error
	var client = &http.Client{}
	var data Users

	session := sessions.Default(c)
	token := session.Get("jwtAct")
	if token == nil || token == "" {
		log.Println("TIDAK ADA SESSION USER ", token)
		session.Clear()
		session.Save()
		return data, errors.New("Session telah habis.")
	}

	body, _ := json.Marshal(req.(ReqPengguna))
	var payload = bytes.NewBuffer(body)

	request, err := http.NewRequest("POST", `http://10.90.90.1:1990/users`, payload)
	// request, err := http.NewRequest("POST", os.Getenv("URL_BE")+`/login`, payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/json")
	// request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("X-API-TOKEN", os.Getenv("API_TOKEN"))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}

	if response.StatusCode == 401 {
		log.Println("TIDAK ADA SESSION USER ", token)
		session.Clear()
		session.Save()
		return data, errors.New("Session telah habis.")
	}
	defer response.Body.Close()

	r := models.Response{
		Data: &Users{},
	}

	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return data, err
	}

	if r.Status != 200 {

		return data, errors.New(r.Message)
	}

	users, _ := r.Data.(*Users)

	log.Println(users.Total)
	if users == nil {
		return data, errors.New("Gagal memuat data.")
	}

	return *users, nil
}

func GetUser(id string, c *gin.Context) (models.User, error) {
	var err error
	var client = &http.Client{}
	var data models.User

	session := sessions.Default(c)
	token := session.Get("jwtAct")
	if token == nil || token == "" {
		log.Println("TIDAK ADA SESSION USER ", token)
		session.Clear()
		session.Save()
		return data, errors.New("Session telah habis.")
	}

	// var param = url.Values{}
	// param.Set("username", username)
	// param.Set("password", password)
	// var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("GET", `http://10.90.90.1:1990/user/`+id, nil)
	// request, err := http.NewRequest("POST", os.Getenv("URL_BE")+`/login`, payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/json")
	// request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("X-API-TOKEN", os.Getenv("API_TOKEN"))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}

	if response.StatusCode == 401 {
		log.Println("TIDAK ADA SESSION USER ", token)
		session.Clear()
		session.Save()
		return data, errors.New("Session telah habis.")
	}
	defer response.Body.Close()

	r := models.Response{
		Data: &models.User{},
	}

	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return data, err
	}

	if r.Status != 200 {
		return data, errors.New(r.Message)
	}

	user, _ := r.Data.(*models.User)

	log.Println(user.Fullname)
	if user == nil {
		return data, errors.New("Gagal memuat data.")
	}

	return *user, nil
}

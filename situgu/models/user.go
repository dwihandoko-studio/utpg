package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	ID             string      `json:"id"`
	Fullname       string      `json:"fullname"`
	Email          string      `json:"email"`
	Nip            string      `json:"nip"`
	NoHp           string      `json:"no_hp"`
	JenisKelamin   string      `json:"jenis_kelamin"`
	Jabatan        string      `json:"jabatan"`
	Npsn           string      `json:"npsn"`
	Kecamatan      interface{} `json:"kecamatan"`
	SuratTugas     string      `json:"surat_tugas"`
	ProfilePicture interface{} `json:"profile_picture"`
	RoleUser       int         `json:"role_user"`
	LastActive     string      `json:"last_active"`
	UpdatedAt      string      `json:"updated_at"`
}

func GetUser(token string) (User, error) {
	var err error
	var client = &http.Client{}
	var data User

	// var param = url.Values{}
	// param.Set("username", username)
	// param.Set("password", password)
	// var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("GET", `http://10.90.90.1:1990/user`, nil)
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
	defer response.Body.Close()

	r := Response{
		Data: &User{},
	}

	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return data, err
	}

	if r.Status != 200 {
		return data, errors.New(r.Message)
	}

	user, _ := r.Data.(*User)

	log.Println(user.Fullname)
	if user == nil {
		return data, errors.New("Gagal memuat data.")
	}

	return *user, nil
}

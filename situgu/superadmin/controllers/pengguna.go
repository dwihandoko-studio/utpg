package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dwihandoko-studio/utpg/layanan/config"
	"github.com/dwihandoko-studio/utpg/layanan/models"
	sm "github.com/dwihandoko-studio/utpg/situgu/superadmin/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PenggunaPage(c *gin.Context) {
	session := sessions.Default(c)
	loggedIn := session.Get("loggedIn")

	if loggedIn != true {
		location := url.URL{Path: "/auth"}
		c.Redirect(http.StatusFound, location.RequestURI())
	} else {
		user := session.Get("user")
		if user == nil || user == "" {
			log.Println("Error get info token", user)
			location := url.URL{Path: "/auth"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}

		if user.(models.User).RoleUser != 1 {
			location := url.URL{Path: "/situgu"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}

		log.Println("REQUEST PATH", c.Request.URL.Path)

		c.HTML(http.StatusOK,
			"Pengguna",
			gin.H{
				"message":  "ERROR",
				"segment":  c.Request.URL.Path,
				"base_url": config.BaseUrl(c),
				"User": gin.H{
					"fullname": "handoko",
				},
			})
	}
}

func ReqPenggunaAll(c *gin.Context) {
	session := sessions.Default(c)
	loggedIn := session.Get("loggedIn")

	if loggedIn != true {
		config.ReturnJsonError(c, 401, errors.New("Session telah habis."), "auth")
		return
	} else {
		user := session.Get("user")
		if user == nil || user == "" {
			config.ReturnJsonError(c, 401, errors.New("Session telah habis."), "auth")
			return
		}

		if user.(models.User).RoleUser != 1 {
			config.ReturnJsonError(c, 303, errors.New("Session telah habis."), "situgu")
			return
		}

		reqBody := c.Request.Body

		body, err := ioutil.ReadAll(reqBody)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Request Body:\n%s\n", body)

		var d sm.DTJsonPengguna
		err = json.Unmarshal(body, &d)
		if err != nil {
			log.Println(err.Error())
		}
		// fmt.Println("Data Convertan: ", spew.Sdump(d))

		// var draw string = "0"
		// // var start string = "0"
		// // var length string = "10"
		// // var keyword string
		// // var role string

		// json_map := make(map[string]interface{})
		// // err := json.NewDecoder(body).Decode(&json_map)

		// log.Println("FORM :", c.Request.FormValue("draw"))

		// if err != nil {
		// 	log.Println("FORM PARAMS")
		// 	log.Println("ERRORNYA: ", err.Error())
		// 	draw = c.PostForm("draw")
		// 	start = c.PostForm("start")
		// 	length = c.PostForm("length")
		// 	keyword = c.PostForm("keyword")
		// 	role = c.PostForm("role")
		// } else {
		// 	//json_map has the JSON Payload decoded into a map
		// 	draw = fmt.Sprintf("%s", json_map["draw"])
		// 	start = fmt.Sprintf("%s", json_map["start"])
		// 	length = fmt.Sprintf("%s", json_map["length"])
		// 	keyword = fmt.Sprintf("%s", json_map["keyword"])
		// 	role = fmt.Sprintf("%s", json_map["role"])
		// }

		res, err := sm.GetUsers(sm.ReqPengguna{
			Start:   strconv.Itoa(d.Start),
			Length:  strconv.Itoa(d.Length),
			Keyword: d.Search.Value,
			Role:    d.Role,
		}, c)

		if err != nil {
			c.JSON(
				http.StatusOK,
				gin.H{
					"draw":            d.Draw,
					"recordsTotal":    0,
					"recordsFiltered": 0,
					"data":            []string{},
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"draw":            d.Draw,
				"recordsTotal":    res.Total,
				"recordsFiltered": res.Total,
				"data":            res.Data,
			},
		)
		return

	}
}

func DetailPengguna(c *gin.Context) {
	session := sessions.Default(c)
	loggedIn := session.Get("loggedIn")

	if loggedIn != true {
		config.ReturnJsonError(c, 401, errors.New("Session telah habis."), "auth")
		return
	} else {
		user := session.Get("user")
		if user == nil || user == "" {
			config.ReturnJsonError(c, 401, errors.New("Session telah habis."), "auth")
			return
		}

		if user.(models.User).RoleUser != 1 {
			config.ReturnJsonError(c, 303, errors.New("Session telah habis."), "situgu")
			return
		}

		id := c.PostForm("id")
		res, err := sm.GetUser(id, c)
		if err != nil {
			log.Println("error get user: ", err.Error())
			config.ReturnJsonError(c, 400, err, "")
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"status":  200,
				"message": "berhasil",
				"data":    template.HTML(GetFormDetail(res)),
			},
		)
		return
	}
}

func GetFormDetail(user models.User) string {
	buffer := &bytes.Buffer{}

	temp, _ := template.ParseFiles("html/detail_pengguna.html")

	data := map[string]interface{}{
		"user": user,
	}

	temp.ExecuteTemplate(buffer, "detail_pengguna.html", data)

	return buffer.String()
}

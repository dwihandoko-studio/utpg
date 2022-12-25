package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/securecookie"
)

type M map[string]interface{}

type JwtCustomClaims struct {
	Id    string `json:"id"`
	Level int64  `json:"level"`
	jwt.StandardClaims
}

var sc = securecookie.New([]byte("com-handoko"), []byte("com-handoko-utpg"))

func SetCookie(c *gin.Context, name string, data M) error {
	encoded, err := sc.Encode(name, data)
	if err != nil {
		return err
	}

	cookie, err := c.Cookie("lyn_cookie")
	if err != nil {
		cookie = "NotSet"
		c.SetCookie("lyn_cookie", encoded, 3600*24, "/", "localhost", false, true)
	}

	log.Printf("Cookie value: %s \n", cookie)

	return err
}

func GetCookie(c *gin.Context, name string) (M, error) {
	cookie, err := c.Cookie("lyn_cookie")
	if err == nil {
		data := M{}
		if err = sc.Decode(name, cookie, &data); err == nil {
			return data, nil
		}
	}

	return nil, err
}

func GetInfoJwt(tokenString string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret key handokowae.my.id"), nil
	})

	if err != nil {
		log.Println("Error decode token: ", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		log.Printf("%v %v", claims.Id, claims.Level)
		return claims, nil
	} else {
		log.Println(err)
		return nil, err
	}
}

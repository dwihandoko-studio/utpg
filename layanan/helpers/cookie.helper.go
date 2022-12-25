package helpers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/securecookie"
)

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret key handokowae.my.id")

type M map[string]interface{}

type JwtCustomClaims struct {
	Id    string `json:"id"`
	Level int64  `json:"level"`
	jwt.StandardClaims
}

var sc = securecookie.New([]byte("com-handoko"), []byte("com-handoko-utpg"))

func SetCookie(c *gin.Context, name string, data string) error {
	// encoded, err := sc.Encode(name, data)
	// if err != nil {
	// 	return err
	// }

	// cookie, err := c.Cookie("lyn_cookie")
	// if err != nil {
	// 	cookie = "NotSet"
	// c.SetCookie("label", "ok", 30, "/", "localhost", false, true)
	c.SetCookie("lyn_cookie", data, 3600, "/", "localhost", false, true)
	// }

	// log.Printf("Cookie value: %s \n", encoded)

	return nil
}

func GetCookie(c *gin.Context, name string) (*string, error) {
	cookie, err := c.Cookie(name)
	// cookie, err: = c.Request.Cookie(name)
	if err != nil {
		return nil, err
		// data := M{}
		// if err = sc.Decode(name, cookie, &data); err == nil {
		// 	return data, nil
		// }
	}

	return &cookie, err
}

func GetInfoJwt(tokenString string) (jwt.MapClaims, error) {
	// token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte("secret key handokowae.my.id"), nil
	// })

	// if err != nil {
	// 	log.Println("Error decode token: ", err)
	// 	return nil, err
	// }

	// if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
	// 	log.Printf("%v %v", claims.Id, claims.Level)
	// 	return claims, nil
	// } else {
	// 	log.Println(err)
	// 	return nil, err
	// }

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

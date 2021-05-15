package security

import (
	"fmt"
	"net/http"
	"time"
	"trackly-backend/app/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/matthewhartstonge/argon2"
)

var mySigningKey = []byte("supersecretboogaloo")

type Security struct {
	AuthorizationEnabled bool
}

type PasswordConfig struct {
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Camille Feghali"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	utils.CheckError(err)
	return tokenString, err
}

func (security *Security) IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if security.AuthorizationEnabled {
			if r.Header["Token"] != nil {
				token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an error")
					}
					return mySigningKey, nil
				})
				if err != nil {
					fmt.Println(err.Error())
					utils.ErrorResponse(w, 403, err.Error())
				}
				if token.Valid {
					endpoint(w, r)
				}
			} else {
				utils.ErrorResponse(w, 403, "Not Authorized")
			}
		} else {
			endpoint(w, r)
		}
	})
}

func Encrypt(password string) (string, error) {
	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(password))
	utils.LogError(err)
	return string(encoded), err
}

func PasswordMatches(password, hash string) (bool, error) {
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(hash))
	utils.LogError(err)
	return ok, err
}

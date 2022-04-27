package Service

import (
	"encoding/json"
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginMethod interface {
	LoginResponse()
}

// LoginResponse is a function that returns a response to the user
func (l LoginUser) LoginResponse(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if l.Email == "" || l.Password == "" {
		Helpers.ResponseMessage(w, http.StatusBadRequest, "Invalid Input")
		return
	}

	hashPass := ExistsUserPassword(w, l.Email)
	if len(hashPass) != 0 && ValidatePassword(w, hashPass, l.Password) {
		token, err := MakeJwtToken(w, l.Email)
		if err != nil {
			Helpers.ResponseMessage(w, http.StatusBadGateway, "Error while generating token")
		}
		loginResponse := struct {
			Token   string `json:"token"`
			Message string `json:"message"`
		}{
			Token:   token,
			Message: "Login Successfully",
		}
		r.Header.Add("User-Id", UserId(w, l.Email))
		fmt.Println(r.Header)

		Helpers.ResponseMessage(w, http.StatusOK, loginResponse)

	}

}

// MakeJwtToken make a jwt token for user
func MakeJwtToken(w http.ResponseWriter, data interface{}) (string, error) {
	secret, err := MakeJwtSecret()
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusInternalServerError, "Error while making jwt secret")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["data"] = data
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusInternalServerError, "Error while making jwt token")
	}
	return tokenString, nil
}

// MakeJwtSecret make jwt secret
func MakeJwtSecret() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	secret := os.Getenv("JWT_SECRET")
	if len(secret) == 0 {
		secret = "01307997323"
	}
	return secret, nil
}

// DecodeJwtToken decode JWT token and jwt token validation and token expiration check and token signature check
func DecodeJwtToken(w http.ResponseWriter, tokenString string) (jwt.MapClaims, error) {
	secret, err := MakeJwtSecret()
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusInternalServerError, "Error while making jwt secret")
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	err = claims.Valid()
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusUnauthorized, "Invalid JWT token")
		return nil, err
	}
	if !token.Valid {
		Helpers.ResponseMessage(w, http.StatusUnauthorized, "Invalid JWT token")
		return nil, err
	}

	return claims, nil
}

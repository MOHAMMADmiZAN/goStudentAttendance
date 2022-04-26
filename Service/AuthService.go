package Service

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Utils"
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

// MakeJwtToken make a jwt token for user
func MakeJwtToken(w http.ResponseWriter, data interface{}) (string, error) {
	secret, err := MakeJwtSecret()
	if err != nil {
		Utils.ResponseMessage(w, http.StatusInternalServerError, "Error while making jwt secret")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["data"] = data
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		Utils.ResponseMessage(w, http.StatusInternalServerError, "Error while making jwt token")
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
		Utils.ResponseMessage(w, http.StatusInternalServerError, "Error while making jwt secret")
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
		Utils.ResponseMessage(w, http.StatusUnauthorized, "Invalid JWT token")
		return nil, err
	}
	if !token.Valid {
		Utils.ResponseMessage(w, http.StatusUnauthorized, "Invalid JWT token")
		return nil, err
	}

	return claims, nil
}

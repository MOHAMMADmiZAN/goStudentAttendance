package Service

import (
	"encoding/json"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strconv"
	"time"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginMethod interface {
	LoginResponse()
}

// VerifyRequestUser verify Request User struct
type VerifyRequestUser struct {
	Id         string `json:"id"`
	ExpireTime int64  `json:"expireTime"`
}

var LogVerify VerifyRequestUser

// verify request interface
type verifyRequestUserMethod interface {
	StoreVerifyRequest(id string, exp int64) *VerifyRequestUser
	GetIdFromVerifyRequest() string
	GetExpireTimeFromVerifyRequest() int64
}

//StoreVerifyRequest is a Method to store the verify request
func (v VerifyRequestUser) StoreVerifyRequest(id string, exp int64) VerifyRequestUser {
	return VerifyRequestUser{
		Id:         id,
		ExpireTime: exp,
	}
}

// GetIdFromVerifyRequest get id form the verify request
func (v VerifyRequestUser) GetIdFromVerifyRequest() string {
	return v.Id
}

// GetExpireTimeFromVerifyRequest get expire time form the verify request
func (v VerifyRequestUser) GetExpireTimeFromVerifyRequest() int64 {
	return v.ExpireTime
}

// LoginResponse is a function that returns a response to the user
func (l LoginUser) LoginResponse(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		Helper.ResponseMessage(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if l.Email == "" || l.Password == "" {
		Helper.ResponseMessage(w, http.StatusBadRequest, "Invalid Input")
		return
	}

	hashPass := ExistsUserPassword(w, l.Email)
	if len(hashPass) != 0 && ValidatePassword(w, hashPass, l.Password) {
		token, err := MakeJwtToken(w, l.Email)
		if err != nil {
			Helper.ResponseMessage(w, http.StatusBadGateway, "Error while generating token")
		}
		loginResponse := struct {
			Token   string `json:"token"`
			Message string `json:"message"`
		}{
			Token:   token,
			Message: "Login Successfully",
		}
		jwtToken, err := DecodeJwtToken(w, token)
		if err != nil {
			Helper.ResponseMessage(w, http.StatusBadGateway, "Error while decoding token")
		}
		id := UserId(w, l.Email)
		exp := int64(jwtToken["exp"].(float64))
		LogVerify = VerifyRequestUser.StoreVerifyRequest(LogVerify, id, exp)
		Helper.MakeSecureCookie(w, r, "UserData", map[string]string{
			"Id":         id,
			"Email":      l.Email,
			"ExpireTime": strconv.FormatInt(exp, 10),
		}, exp)
		Helper.ResponseMessage(w, http.StatusOK, loginResponse)

	}

}

// MakeJwtToken make a jwt token for user
func MakeJwtToken(w http.ResponseWriter, data interface{}) (string, error) {
	secret, err := MakeJwtSecret()
	if err != nil {
		Helper.ResponseMessage(w, http.StatusInternalServerError, "Error while making jwt secret")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["data"] = data
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		Helper.ResponseMessage(w, http.StatusInternalServerError, "Error while making jwt token")
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
		Helper.ResponseMessage(w, http.StatusInternalServerError, "Error while making jwt secret")
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
		Helper.ResponseMessage(w, http.StatusUnauthorized, "Invalid JWT token")
		return nil, err
	}
	if !token.Valid {
		Helper.ResponseMessage(w, http.StatusUnauthorized, "Invalid JWT token")
		return nil, err
	}

	return claims, nil
}

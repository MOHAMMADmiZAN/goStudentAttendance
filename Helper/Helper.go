package Helper

import (
	"encoding/json"
	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
	"math/rand"
	"net/http"
	"os"
)

var hashKey = securecookie.GenerateRandomKey(32)
var blockKey = securecookie.GenerateRandomKey(32)
var C *securecookie.SecureCookie = securecookie.New(hashKey, blockKey)

type JsonResponse struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}
type JsonResponseMethod interface {
	GetResponse(w http.ResponseWriter)
}

// GetResponse get response method
func (j JsonResponse) GetResponse(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(j)
	if err != nil {
		ResponseMessage(w, http.StatusInternalServerError, "Error encoding response")
	}

}

// ResponseMessage is a helper function that takes in a status code and a message and returns a JSON response
func ResponseMessage(w http.ResponseWriter, status int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	j := JsonResponse{
		Status:  status,
		Message: message,
	}
	JsonResponse.GetResponse(j, w)

}

// LoadEnv load env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		os.Exit(1)
	}
}

// Contains array contains  function
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// MyArrayMap My Array Map function  takes in an array and a callback and returns a new  array
func MyArrayMap(arr []string, callback func(v string, i int, arr []string)) []string {
	for i := 0; i < len(arr); i++ {
		callback(arr[i], i, arr)
	}
	return arr
}

// GetCookie get cockie function
func GetCookie(w http.ResponseWriter, r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		ResponseMessage(w, http.StatusBadRequest, "Cookie not found")
		return ""

	} else {
		ResponseMessage(w, http.StatusOK, "Cookie found")
	}

	return c.Value
}

// MakeSecureCookie make secure cookie
/**
  Todo: make a secure cookie function with gorilla secure cookie
*/
func MakeSecureCookie(w http.ResponseWriter, r *http.Request, name string, value map[string]string, maxAge int64) {

	encoded, err := C.Encode(name, value)
	if err != nil {
		ResponseMessage(w, http.StatusInternalServerError, "Error encoding cookie")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    encoded,
		MaxAge:   int(maxAge),
		SameSite: http.SameSiteStrictMode,
	})

}

// DecodeSecureCookie decodeSecureCookie decode
func DecodeSecureCookie(w http.ResponseWriter, r *http.Request, name string) (map[string]string, error) {
	c, err := r.Cookie(name)
	if err != nil {
		ResponseMessage(w, http.StatusBadRequest, "Cookie not found")
		return nil, err
	}
	value := make(map[string]string)
	err = C.Decode(name, c.Value, &value)
	if err != nil {
		ResponseMessage(w, http.StatusInternalServerError, "Error decoding cookie")
		return nil, err
	}
	return value, nil

}

// DeleteAllCookies all cookies delete

// RandomString randomString random string
func RandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

package Helpers

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

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

// AddCookie for global Request
func AddCookie(w http.ResponseWriter, r *http.Request, name string, value string) {
	c := &http.Cookie{
		Name:  name,
		Value: value,
	}
	http.SetCookie(w, c)
}

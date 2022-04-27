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

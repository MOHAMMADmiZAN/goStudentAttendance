package Utils

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

// ResponseMessage is a helper function that takes in a status code and a message and returns a JSON response
func ResponseMessage(w http.ResponseWriter, status int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	j := JsonResponse{
		Status:  status,
		Message: message,
	}
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(j)
	if err != nil {
		ResponseMessage(w, http.StatusInternalServerError, "Error encoding response")
	}

}

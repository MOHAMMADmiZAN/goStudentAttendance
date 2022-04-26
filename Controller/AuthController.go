package Controller

import (
	"encoding/json"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// login user
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var loginUser Service.LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		Utils.ResponseMessage(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if loginUser.Email == "" || loginUser.Password == "" {
		Utils.ResponseMessage(w, http.StatusBadRequest, "Invalid Input")
		return
	}

	hashPass := Service.ExistsUserPassword(w, loginUser.Email)
	if Service.ValidatePassword(w, hashPass, loginUser.Password) {
		token, err := Service.MakeJwtToken(w, loginUser.Email)
		if err != nil {
			Utils.ResponseMessage(w, http.StatusBadGateway, "Error while generating token")
		}
		loginResponse := struct {
			Token   string `json:"token"`
			Message string `json:"message"`
		}{
			Token:   token,
			Message: "Login Successfully",
		}
		Utils.ResponseMessage(w, http.StatusOK, loginResponse)

	}

}

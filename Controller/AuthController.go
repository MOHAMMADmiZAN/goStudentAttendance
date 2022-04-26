package Controller

import (
	"encoding/json"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helpers"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
)

var wg2 sync.WaitGroup

// Login login user
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		var loginUser Service.LoginUser
		err := json.NewDecoder(r.Body).Decode(&loginUser)
		if err != nil {
			Helpers.ResponseMessage(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if loginUser.Email == "" || loginUser.Password == "" {
			Helpers.ResponseMessage(w, http.StatusBadRequest, "Invalid Input")
			return
		}

		hashPass := Service.ExistsUserPassword(w, loginUser.Email)
		if Service.ValidatePassword(w, hashPass, loginUser.Password) {
			token, err := Service.MakeJwtToken(w, loginUser.Email)
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
			Helpers.ResponseMessage(w, http.StatusOK, loginResponse)

		}
	}()
	wg2.Wait()

}

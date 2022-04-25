package Controller

import (
	"encoding/json"
	"fmt"
	"github.com/Kamva/mgm"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Model"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// CreateUser create new user //
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newUser Service.CreateRequestUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(newUser.Roles) == 0 {
		newUser.Roles = []string{"User"}
	}
	if newUser.AccountStatus == "" {
		newUser.AccountStatus = "Pending"
	}
	hashedPassword := Service.PasswordHash(newUser.Password)
	user := Model.UserModel(newUser.Name, newUser.Email, hashedPassword, newUser.Roles, newUser.AccountStatus)
	err = mgm.Coll(user).Create(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

// get user by id
func GetUser(w http.ResponseWriter, _, p httprouter.Params) {
	id := p.ByName("id")
	user := &Model.User{}
	coll := mgm.Coll(user)
	err := coll.FindByID(id, user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

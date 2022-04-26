package Controller

import (
	"encoding/json"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Model"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Utils"
	"github.com/julienschmidt/httprouter"
	"github.com/kamva/mgm/v3"
	"log"
	"net/http"
)

// CreateUser create new user //
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newUser Service.CreateRequestUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		Utils.ResponseMessage(w, http.StatusBadRequest, "Invalid request body")
	}
	if len(newUser.Roles) == 0 {
		newUser.Roles = []string{"User"}
	}
	if newUser.AccountStatus == "" {
		newUser.AccountStatus = "Pending"
	}
	if Service.DuplicateUser(w, newUser.Email) {
		hashedPassword := Service.PasswordHash(newUser.Password)
		user := Model.UserModel(newUser.Name, newUser.Email, hashedPassword, newUser.Roles, newUser.AccountStatus)
		err = mgm.Coll(user).Create(user)
		if err != nil {
			Utils.ResponseMessage(w, http.StatusInternalServerError, "Internal server error")
		}
		Utils.ResponseMessage(w, http.StatusCreated, user)
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

// delete user by id
func DeleteUser(w http.ResponseWriter, _, p httprouter.Params) {
	id := p.ByName("id")
	user := &Model.User{}
	coll := mgm.Coll(user)
	err := coll.FindByID(id, user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = coll.Delete(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// update user by id
func UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	var updateUser Service.CreateRequestUser
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := &Model.User{}
	coll := mgm.Coll(user)
	err = coll.FindByID(id, user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}
	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}
	if updateUser.Password != "" {
		user.Password = Service.PasswordHash(updateUser.Password)
	}
	if len(updateUser.Roles) != 0 {
		user.Roles = updateUser.Roles
	}
	if updateUser.AccountStatus != "" {
		user.AccountStatus = updateUser.AccountStatus
	}
	err = coll.Update(user)
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

// user find by email

package Controller

import (
	"encoding/json"
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helpers"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Model"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/julienschmidt/httprouter"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

// CreateUser create new user //
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newUser Service.CreateRequestUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if len(newUser.Roles) == 0 {
		newUser.Roles = []string{"USER"}
	}
	if len(newUser.Roles) > 0 {
		for _, role := range newUser.Roles {
			if Helpers.Contains(Service.UserRoles, role) {
				strings.ToUpper(role)
				continue
			} else {
				roleErr := fmt.Sprintf("Role %s is not allowed", role)
				Helpers.ResponseMessage(w, http.StatusBadRequest, roleErr)
				return
			}

		}
	}

	if newUser.AccountStatus == "" {
		newUser.AccountStatus = "PENDING"
	}
	if !Service.DuplicateUser(w, newUser.Email) {
		hashedPassword := Service.PasswordHash(newUser.Password)
		user := Model.UserModel(newUser.Name, newUser.Email, hashedPassword, newUser.Roles, newUser.AccountStatus)
		err = mgm.Coll(user).Create(user)
		if err != nil {
			Helpers.ResponseMessage(w, http.StatusInternalServerError, "Internal server error")
			return
		}
		Helpers.ResponseMessage(w, http.StatusCreated, "User created successfully")

	}

}

// GetUser get user by id
func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	user := &Model.User{}
	coll := mgm.Coll(user)
	err := coll.FindByID(id, user)
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusNotFound, "User not found")
		return
	}
	Helpers.ResponseMessage(w, http.StatusOK, user)

}

// DeleteUser delete user by id
func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	user := &Model.User{}
	coll := mgm.Coll(user)
	err := coll.FindByID(id, user)
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusNotFound, "User not found")
		return
	}
	err = coll.Delete(user)
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	Helpers.ResponseMessage(w, http.StatusOK, "User deleted Successfully")

}

// UpdateUser update user by id
func UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	var updateUser Service.CreateRequestUser
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	user := &Model.User{}
	coll := mgm.Coll(user)
	err = coll.FindByID(id, user)
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusNotFound, "User not found")
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
		Helpers.ResponseMessage(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	Helpers.ResponseMessage(w, http.StatusOK, "User updated Successfully")

}

// GetAllUsers get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	user := &Model.User{}
	coll := mgm.Coll(user)
	var users []Model.User
	err := coll.SimpleFind(&users, bson.M{})
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	Helpers.ResponseMessage(w, http.StatusOK, users)

}

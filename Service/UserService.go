package Service

import (
	"encoding/json"
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helpers"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Model"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

// CreateRequestUser create new user struct
type CreateRequestUser struct {
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	Password      string   `json:"password"`
	Roles         []string `json:"roles"`
	AccountStatus string   `json:"account_status"`
}
type CreateRequestUserMethod interface {
	CreateUser()
}

// CreateUser create new user
func (u CreateRequestUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		Helpers.ResponseMessage(w, http.StatusBadRequest, "Invalid request Input")
		return
	}
	if len(u.Roles) == 0 {
		u.Roles = []string{"USER"}
	}
	if len(u.Roles) > 0 {
		for _, role := range u.Roles {
			if Helpers.Contains(UserRoles, role) {
				strings.ToUpper(role)
				continue
			} else {
				roleErr := fmt.Sprintf("Role %s is not allowed", role)
				Helpers.ResponseMessage(w, http.StatusBadRequest, roleErr)
				return
			}

		}
	}

	if u.AccountStatus == "" {
		u.AccountStatus = "PENDING"
	}
	if !DuplicateUser(w, u.Email) {
		hashedPassword := PasswordHash(u.Password)
		user := Model.UserModel(u.Name, u.Email, hashedPassword, u.Roles, u.AccountStatus)
		err = mgm.Coll(user).Create(user)
		if err != nil {
			Helpers.ResponseMessage(w, http.StatusInternalServerError, "Internal server error")
			return
		}
		Helpers.ResponseMessage(w, http.StatusCreated, "User created successfully")

	}
}

// UserRoles user Role
var UserRoles = []string{"ADMIN", "USER", "student"}

// PasswordHash Password hashing
func PasswordHash(pass string) string {
	pw := []byte(pass)
	password, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(password)
}

// DuplicateUser Duplicate User Find
func DuplicateUser(w http.ResponseWriter, email string) bool {
	if ExistsUser(w, email) {
		Helpers.ResponseMessage(w, http.StatusBadRequest, "User Already Exists")
		return true
	}
	return false
}

// ExistsUserPassword Exits User Password
func ExistsUserPassword(w http.ResponseWriter, email string) string {
	user := &Model.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusNotFound, "User Not Exists")
		return ""
	}

	return user.Password
}

// ValidatePassword bcrypt password validation
func ValidatePassword(w http.ResponseWriter, hashedPassword string, password string) bool {
	byteHash := []byte(hashedPassword)
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusBadRequest, "Password Not Match")
		return false
	}
	return true
}

// UserId  find by email
func UserId(w http.ResponseWriter, email string) string {
	user := &Model.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		Helpers.ResponseMessage(w, http.StatusNotFound, "User Not Exists")
	}
	return user.ID.Hex()
}

//ExistsUser Exists User
func ExistsUser(w http.ResponseWriter, email string) bool {
	user := &Model.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {

		return false
	}
	return true
}

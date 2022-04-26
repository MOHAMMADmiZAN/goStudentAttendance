package Service

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Model"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Utils"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// CreateRequestUser create new user struct
type CreateRequestUser struct {
	Name          string   `json:"username"`
	Email         string   `json:"email"`
	Password      string   `json:"password"`
	Roles         []string `json:"roles"`
	AccountStatus string   `json:"account_status"`
}

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
		Utils.ResponseMessage(w, http.StatusBadRequest, "User Already Exists")
		return true
	}
	return false
}

// ExistsUserPassword Exits User Password
func ExistsUserPassword(w http.ResponseWriter, email string) string {
	user := &Model.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		Utils.ResponseMessage(w, http.StatusBadRequest, "User Not Exists")
	}

	return user.Password
}

// ValidatePassword becript password validation
func ValidatePassword(w http.ResponseWriter, hashedPassword string, password string) bool {
	byteHash := []byte(hashedPassword)
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		Utils.ResponseMessage(w, http.StatusBadRequest, "Password Not Match")
	}
	return true
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

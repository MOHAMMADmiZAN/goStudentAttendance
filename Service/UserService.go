package Service

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Model"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Utils"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type CreateRequestUser struct {
	Name          string   `json:"username"`
	Email         string   `json:"email"`
	Password      string   `json:"password"`
	Roles         []string `json:"roles"`
	AccountStatus string   `json:"account_status"`
}

func PasswordHash(pass string) string {
	pw := []byte(pass)
	password, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(password)
}
func DuplicateUser(w http.ResponseWriter, email string) bool {
	user := &Model.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err == nil {
		Utils.ResponseMessage(w, http.StatusBadRequest, "User Already Exists")
		return false
	}
	return true
}

// create a jwt token for the user

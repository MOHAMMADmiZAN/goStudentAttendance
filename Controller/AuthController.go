package Controller

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Register  user
func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var registerUser Service.CreateRequestUser
	registerUser.CreateUser(w, r)
}

//Login login user
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var loginUser Service.LoginUser
	loginUser.LoginResponse(w, r)

}

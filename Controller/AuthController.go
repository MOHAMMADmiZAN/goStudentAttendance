package Controller

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Login login user
func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var registerUser Service.CreateRequestUser
	Service.CreateRequestUser.CreateUser(registerUser, w, r)
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var loginUser Service.LoginUser
	Service.LoginUser.LoginResponse(loginUser, w, r)

}

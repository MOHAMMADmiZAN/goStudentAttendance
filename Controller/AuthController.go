package Controller

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Login login user
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var loginUser Service.LoginUser
	Service.LoginUser.LoginResponse(loginUser, w, r)

}

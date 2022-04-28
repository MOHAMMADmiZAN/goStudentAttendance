package Middleware

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helpers"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

// Auth Middleware with JWT Token Validation and Authorization Check for all routes except login and register routes with httprouter
func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authToken := r.Header.Get("Authorization")

		if len(authToken) != 0 {
			authToken = strings.Split(authToken, " ")[1]
			token, err := Service.DecodeJwtToken(w, authToken)
			if err != nil {
				Helpers.ResponseMessage(w, http.StatusUnauthorized, "Unauthorized")
				return
			}
			email := token["data"].(string)

			if !Service.ExistsUser(w, email) {
				Helpers.ResponseMessage(w, http.StatusUnauthorized, "please login first")
				return
			}
			/**
			TODO: active  id and exp check in auth middleware

			*/
			/*
				exp := int64(token["exp"].(float64))
						id := Service.UserId(w, email)
				if id != Service.VerifyRequestUser.GetIdFromVerifyRequest(Service.LogVerify) || exp != Service.VerifyRequestUser.GetExpireTimeFromVerifyRequest(Service.LogVerify) {
						Helpers.ResponseMessage(w, http.StatusUnauthorized, "Token is not valid")
						return

					}*/

		}
		next(w, r, ps)

	}

}

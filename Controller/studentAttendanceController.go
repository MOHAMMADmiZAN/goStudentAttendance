package Controller

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helper"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Model"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/julienschmidt/httprouter"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetRunningStatusForStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := Service.FindActiveAttendance()
	if len(result) == 0 {
		Helper.ResponseMessage(w, http.StatusNotFound, "No active attendance found")
		return
	}
	Helper.ResponseMessage(w, http.StatusOK, result)
	return

}
func SubmitAttendance(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var aId = ps.ByName("id")
	res := Service.FindAttendanceByUserId(aId)

	if len(res) == 0 {
		Helper.ResponseMessage(w, http.StatusNotFound, "No active attendance found")
		return
	}
	cookie, err := Helper.DecodeSecureCookie(w, r, "UserData")
	if err != nil {
		Helper.ResponseMessage(w, http.StatusUnauthorized, "Cookie not found")
		return
	}
	var userId = cookie["Id"]
	if !primitive.IsValidObjectID(userId) {
		Helper.ResponseMessage(w, http.StatusBadRequest, "Invalid user id")
		return
	}
	if Service.IfAttend(aId, userId) {
		Helper.ResponseMessage(w, http.StatusBadRequest, "You have already submitted attendance")
		return
	}
	user := Model.StudentAttendanceModel(Service.HexToObjectId(userId), Service.HexToObjectId(aId))
	err = mgm.Coll(user).Create(user)
	if err != nil {
		Helper.ResponseMessage(w, http.StatusInternalServerError, "Error while creating attendance")
		return

	}
	Helper.ResponseMessage(w, http.StatusOK, "Attendance submitted successfully")
	return

}

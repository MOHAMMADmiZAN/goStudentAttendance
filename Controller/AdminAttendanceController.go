package Controller

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helper"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Model"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/julienschmidt/httprouter"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Enable(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var e Service.AdminAttendanceEnable

	//find is Running Attendance Against this id
	k, uid := Service.FindRunningAttendance(w, r)
	if !primitive.IsValidObjectID(uid) {
		return
	}
	if len(k) == 0 {
		e.EnableAttendance(w, r, uid)
		return
	}

	Helper.ResponseMessage(w, http.StatusBadRequest, "Attendance is already running")
}

// DisableAttendance disable attendance
func DisableAttendance(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := &Model.AdminAttendance{}
	err := mgm.Coll(a).FindByID(ps.ByName("id"), a)
	if err != nil {
		Helper.ResponseMessage(w, http.StatusNotFound, "Admin Attendance Not Found")
		return
	}
	a.Status = "COMPLETED"
	err = mgm.Coll(a).Update(a)
	if err != nil {
		Helper.ResponseMessage(w, http.StatusBadRequest, "Attendance Not Disabled")
		return
	}
	Helper.ResponseMessage(w, http.StatusOK, "Attendance Completed")

}

// GetAllAttendance get all attendance
func GetAllAttendance(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var a = &Model.AdminAttendance{}
	var attendance []Model.AdminAttendance
	err := mgm.Coll(a).SimpleFind(&attendance, bson.M{})
	if err != nil {
		Helper.ResponseMessage(w, http.StatusNotFound, "Attendance Not Found")
		return
	}
	Helper.ResponseMessage(w, http.StatusOK, attendance)

}

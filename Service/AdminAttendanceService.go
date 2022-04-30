package Service

import (
	"encoding/json"
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helper"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Model"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

type AdminAttendanceEnable struct {
	TimeLimit int `json:"timeLimit"`
}
type AdminAttendanceMethod interface {
	EnableAttendance(AdminAttendanceEnable)
}

// EnableAttendance Method
func (aas AdminAttendanceEnable) EnableAttendance(w http.ResponseWriter, r *http.Request, uid string) {
	err := json.NewDecoder(r.Body).Decode(&aas)
	if err != nil {
		Helper.ResponseMessage(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	if uid == "" || uid == " " || len(uid) == 0 {
		Helper.ResponseMessage(w, http.StatusBadRequest, "Invalid uId")
	}

	enable := Model.AdminAttendanceModel("RUNNING", aas.TimeLimit, HexToObjectId(uid))
	err = mgm.Coll(enable).Create(enable)
	if err != nil {
		Helper.ResponseMessage(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	Helper.ResponseMessage(w, http.StatusOK, "Attendance Enabled")
	return

}

// FindRunningAttendance find running attendance
func FindRunningAttendance(w http.ResponseWriter, r *http.Request) ([]Model.AdminAttendance, string) {
	cookie, err := Helper.DecodeSecureCookie(w, r, "UserData")
	if err != nil {
		Helper.ResponseMessage(w, http.StatusNotFound, "Cookie not found")
		return nil, ""
	}
	userId := cookie["Id"]
	a := Model.AdminAttendance{}
	var result []Model.AdminAttendance
	err = mgm.Coll(&a).SimpleFind(&result, bson.M{"status": "RUNNING", "userID": HexToObjectId(userId)})
	if err != nil {
		Helper.ResponseMessage(w, http.StatusBadRequest, err.Error())
		return nil, ""
	}
	return result, userId
}

// DisableAttendanceWhenTimeOut DisableAttendance When Time out
func DisableAttendanceWhenTimeOut(res []Model.AdminAttendance) {
	var w http.ResponseWriter
	result := FindActiveAttendance()
	if len(result) > 0 {
		for _, v := range result {
			timeout := v.CreatedAt.Add(time.Duration(v.TimeLimit) * time.Minute)
			if timeout.Before(time.Now()) {
				v.Status = "COMPLETED"
				err := mgm.Coll(&v).Update(&v)
				if err != nil {
					Helper.ResponseMessage(w, http.StatusBadRequest, err.Error())
				}
				fmt.Println("Attendance Completed")

			}

		}
	} else {
		fmt.Println("No Attendance Running")
		Helper.ResponseMessage(w, http.StatusNotFound, "Not Running Attendance Here")

	}
	return

}

//FindActiveAttendance Find Active Attendance
func FindActiveAttendance() []Model.AdminAttendance {
	var w http.ResponseWriter
	var a Model.AdminAttendance
	var result []Model.AdminAttendance
	err := mgm.Coll(&a).SimpleFind(&result, bson.M{"status": "RUNNING"})
	if err != nil {
		Helper.ResponseMessage(w, http.StatusBadRequest, err.Error())
		return nil
	}
	return result
}

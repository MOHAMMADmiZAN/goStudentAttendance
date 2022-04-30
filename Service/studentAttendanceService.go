package Service

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Model"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func IfAttend(aid string, uid string) bool {
	var sa Model.StudentAttendance
	coll := mgm.Coll(&sa)
	err := coll.First(bson.M{"studentId": HexToObjectId(uid), "adminAttendanceID": HexToObjectId(aid)}, &sa)
	if err != nil {
		return false
	}
	return true

}

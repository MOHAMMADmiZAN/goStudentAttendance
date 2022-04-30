package Model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StudentAttendance /*- StudentId
type StudentAttendance struct {
	mgm.DefaultModel  `bson:",inline"`
	StudentId         primitive.ObjectID `json:"studentId" bson:"studentId"`
	AdminAttendanceID primitive.ObjectID `json:"adminAttendanceID" bson:"adminAttendanceID"`
}

func StudentAttendanceModel(studentID primitive.ObjectID, adminAttendanceId primitive.ObjectID) *StudentAttendance {
	return &StudentAttendance{
		StudentId:         studentID,
		AdminAttendanceID: adminAttendanceId,
	}
}

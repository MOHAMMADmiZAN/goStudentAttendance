package Model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*

- Status
- TimeLimit
- UserID

*/
type AdminAttendance struct {
	mgm.DefaultModel `bson:",inline"`
	Status           string             ` json:"status" bson:"status"`
	TimeLimit        int                `json:"timeLimit" bson:"timeLimit"`
	UserID           primitive.ObjectID `json:"userID" bson:"userID,omitempty"`
}

func AdminAttendanceModel(status string, timeLimit int, userId primitive.ObjectID) *AdminAttendance {
	return &AdminAttendance{
		Status:    status,
		TimeLimit: timeLimit,
		UserID:    userId,
	}
}

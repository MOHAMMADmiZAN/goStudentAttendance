package Model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
- First Name
- Last Name
- Phone No
- Profile Picture

*/
type Profile struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string             `json:"firstName" bson:"firstName"`
	LastName         string             `json:"lastName" bson:"lastName"`
	PhoneNo          string             `json:"phoneNo" bson:"phoneNo"`
	ProfilePic       string             `json:"profilePic" bson:"profilePic"`
	UserId           primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
}

func ProfileModel(firstName string, lastName string, phoneNumber string, profilePic string, userId primitive.ObjectID) *Profile {
	return &Profile{
		FirstName:  firstName,
		LastName:   lastName,
		PhoneNo:    phoneNumber,
		ProfilePic: profilePic,
	}
}

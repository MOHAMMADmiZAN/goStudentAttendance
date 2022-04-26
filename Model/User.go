package Model

import "github.com/kamva/mgm/v3"

/**
TODO : User model Validation
*/

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string   `json:"name" bson:"name"`
	Email            string   `json:"email" bson:"email"`
	Password         string   `json:"password" bson:"password"`
	Roles            []string `json:"roles" bson:"roles"`
	AccountStatus    string   `json:"account_status" bson:"account_status"`
}

func UserModel(name string, email string, password string, roles []string, account_status string) *User {
	return &User{
		Name:          name,
		Email:         email,
		Password:      password,
		Roles:         roles,
		AccountStatus: account_status,
	}
}

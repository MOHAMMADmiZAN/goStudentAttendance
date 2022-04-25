package Model

import "github.com/Kamva/mgm"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string   `json:"username" bson:"username"`
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

// make a uuid function for user

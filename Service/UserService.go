package Service

import "golang.org/x/crypto/bcrypt"

type CreateRequestUser struct {
	Name          string   `json:"username"`
	Email         string   `json:"email"`
	Password      string   `json:"password"`
	Roles         []string `json:"roles"`
	AccountStatus string   `json:"account_status"`
}

func PasswordHash(pass string) string {
	pw := []byte(pass)
	password, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(password)
}

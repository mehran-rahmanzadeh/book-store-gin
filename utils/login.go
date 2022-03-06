package utils

import "gin-tutorial/models"

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "mrhnz13@gmail.com",
		password: "1234",
	}
}
func (info *loginInformation) LoginUser(email string, password string) bool {
	count := int64(0)
	models.DB.Model(&models.User{}).Where(
		"email = ? AND password = ?", email, password).Count(&count)
	return count > 0
}

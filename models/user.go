package models

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserRegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserEditInput struct {
	FirstName string `json:"firstName,alpha"`
	LastName  string `json:"lastName,alpha"`
}

package user

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Email	string	`gorm:"unique"`
	Password string
}

type LoginUserForm struct{
	Email string `json:"email"`
	Password string `json:"password"`

}


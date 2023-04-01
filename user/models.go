package user

import "gorm.io/gorm"

type SignUpUser struct{
	gorm.Model
	Email	string	`gorm:"unique"`
	Password string
        FirstName string
        LastName  string
}

type User struct{
     gorm.Model
     Email string
     FirstName string
     LastName  string
     
}

type LoginUserForm struct{
	Email string `json:"email"`
	Password string `json:"password"`

}


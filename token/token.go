package token

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "13gbfd,qlqfqqf,kqvk34vjay3"

func GenerateJWT()(string, error){
	token := jwt.New(jwt.SigningMethodHS256)

	tokenString, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil{
		log.Println("Error in generating Token")
		return "",err
	}
	
	return tokenString, nil
}
package token

import (
	// "fmt"
	"log"

	"github.com/dgrijalva/jwt-go"

	"time"
)

const SECRET_KEY = "13gbfd,qlqfqqf,kqvk34vjay3"

func GenerateJWT(email string)(string, error){
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = email // email is the only unique field in this project at the moment
	claims["exp"] = time.Now().Add(10 * time.Minute)
	

	tokenString, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil{
		log.Println("Error in generating Token")
		return "",err
	}
	
	return tokenString, nil
}

// func ValidateToken(tokenString string){
// 	token, err := jwt.Parse(tokenString,func(token *jwt.Token)(interface{}, error){
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return []byte(SECRET_KEY), nil
// 	})
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid{
// 		fmt.Println(claims)
// 	} else{
// 		fmt.Println("Validate Token Error >>")
// 		fmt.Println(err)
// 	}
	
	
// }
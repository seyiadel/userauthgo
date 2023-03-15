package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	// Own Import
	"github.com/dgrijalva/jwt-go"
	// "github.com/seyiadel/userauthgo/token"
)

const SECRET_KEY = "13gbfd,qlqfqqf,kqvk34vjay3"

func IsAuthorized(handler http.HandlerFunc) http.Handler{
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request){
		headerToken := request.Header.Get("Authorization")
		cleanedToken := strings.TrimPrefix(headerToken, "Bearer ")
		fmt.Println(cleanedToken)

		token, err := jwt.Parse(cleanedToken,func(token *jwt.Token)(interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
	
			return []byte(SECRET_KEY), nil
		})
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid{
			fmt.Println(claims)
		} else{
			fmt.Println("Validate Token Error >>")
			fmt.Println(err)
		}
		ctx := context.WithValue(request.Context(), "loggedInUser", claims)

		handler.ServeHTTP(response, request.WithContext(ctx))



	})

}
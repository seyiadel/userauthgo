package main

import (
	"net/http"
	"github.com/seyiadel/userauthgo/initializers"
	"github.com/seyiadel/userauthgo/user"
	"github.com/seyiadel/userauthgo/middleware"
)


func init(){
	// Commands to run before main function
	//Connect to Database
	initializer.ConnectDB()
}

func main(){

	router := http.NewServeMux()
	router.HandleFunc("/signup", user.SignUpHandler)
	router.HandleFunc("/login", user.LoginHandler)
	router.Handle("/profile", middleware.IsAuthorized(user.UserProfileHandler))

	server := &http.Server{
		Addr:	"0.0.0.0:8080",
		Handler: router,
	}
	server.ListenAndServe()
}
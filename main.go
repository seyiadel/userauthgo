package main

import (
	"net/http"
	"github.com/seyiadel/userauthgo/initializers"
	"github.com/seyiadel/userauthgo/user"
)


func init(){
	// Commands to run before main function
	//Connect to Database
	initializer.ConnectDB()
}

func main(){

	router := http.NewServeMux()
	router.HandleFunc("/signup", user.SignUpHandler)

	server := &http.Server{
		Addr:	"0.0.0.0:8080",
		Handler: router,
	}
	server.ListenAndServe()
}
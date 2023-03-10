package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/seyiadel/userauthgo/initializers"
	"github.com/seyiadel/userauthgo/token"
	"golang.org/x/crypto/bcrypt"
)


func SignUpHandler(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type", "application/json")
	
	var body struct{
		Email string
		Password string
	}

	_= json.NewDecoder(request.Body).Decode(&body)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil{
		panic(err)
	}

	user := User{
		Email:body.Email,
		Password: string(hashPassword),
	}
	
	//Add user Info to "gorm" database
	initializer.DB.Create(&user)
	

	err = json.NewEncoder(response).Encode(user)
	if err != nil{ 
		response.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	response.WriteHeader(http.StatusCreated)
}


func LoginHandler(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type", "application/json")
	var form *LoginUserForm

	_ = json.NewDecoder(request.Body).Decode(&form)

	var user User

	findUserEmail := initializer.DB.First(&user, "email = ?", form.Email)
	if findUserEmail.Error != nil{
		log.Println(findUserEmail)
		response.Write([]byte(`"response":"Incorrect Email"`))
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	
	comparePasswordUnhashed := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if comparePasswordUnhashed != nil{
		log.Println(comparePasswordUnhashed)
		response.Write([]byte(`{"response":"Incorrect Password"`))
		return
	}

	token, err := token.GenerateJWT()
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"response":"`+ err.Error()+`"}`))
		return
	}

	response.Write([]byte(`{"status": "true", "token": "`+token+`"}`))
	if err != nil{ 
		response.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	response.WriteHeader(http.StatusOK)
	
}
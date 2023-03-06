package user

import 	(
	"net/http"
	"github.com/seyiadel/userauthgo/initializers"
	"encoding/json"
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
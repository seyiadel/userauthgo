package initializer

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDB(){
	var err error
	//Connect to Database through Gorm
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
	  panic("failed to connect database")
	}

}

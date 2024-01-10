package databases

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/Models"
)

var UDB *gorm.DB

func connectUser() {
	connectUserAuth, err := gorm.Open(mysql.Open("root@tcp(localhost:3306)/barengin"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	UDB = connectUserAuth
	connectUserAuth.AutoMigrate(&Models.AuthUsers{})
}

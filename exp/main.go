package main

import (
	"fmt"

	"github.com/lenslocked/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "lenslocked_test"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()
	//us.AutoMigrate()

	user := models.User{
		Name:     "Stefan",
		Email:    "stefan@mail.com",
		Password: "stefan",
		Remember: "abc123",
	}
	err = us.Create(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", user)
	user2, err := us.ByRemember("abc123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *user2)
}

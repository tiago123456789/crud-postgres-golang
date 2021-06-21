package main

import (
	"crud-postgres-golang/models"
	"crud-postgres-golang/repositories"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	repositories.UpdateUser(2, models.User{
		Id:       0,
		Name:     "John updated",
		Email:    "johnupdated@gmail.com",
		Password: "johnpasswordupdated",
	})
	// userCreated := createUser(User{
	// 	Id:       0,
	// 	Name:     "John",
	// 	Email:    "john@gmail.com",
	// 	Password: "johnteste",
	// })
	users := repositories.GetUsers()
	// user := GetUserById("1")
	fmt.Println(users)

}

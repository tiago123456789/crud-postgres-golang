package repositories

import (
	"crud-postgres-golang/database"
	"crud-postgres-golang/models"

	"fmt"
	"log"
)

func GetUsers() []models.User {
	db := database.CreateConnection()

	rows, err := db.Query("SELECT * from users")
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		users = append(users, user)
	}

	return users
}

func GetUserById(id string) models.User {
	db := database.CreateConnection()
	defer db.Close()

	row := db.QueryRow("SELECT * from users where id = $1", id)
	var user models.User

	row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	return user
}

func UpdateUser(id int64, user models.User) bool {
	db := database.CreateConnection()
	defer db.Close()

	_, err := db.Exec(
		"UPDATE users SET name=$2, email=$3, password=$4 WHERE id=$1",
		id, user.Name, user.Email, user.Password,
	)

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func deleteUser(id int64) bool {
	db := database.CreateConnection()
	defer db.Close()

	_, err := db.Exec("DELETE FROM users where id = $1", id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func createUser(user models.User) models.User {
	db := database.CreateConnection()
	defer db.Close()

	var id int64

	err := db.QueryRow(
		"INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id",
		user.Name, user.Email, user.Password,
	).Scan(&id)

	if err != nil {
		fmt.Println("Fail when try insert new register")
	}

	user.Id = id
	return user
}

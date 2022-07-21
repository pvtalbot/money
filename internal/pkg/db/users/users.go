package users

import (
	database "back_go/internal/pkg/db/mysql"
	"log"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (user User) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Users(Name, Password) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(user.Name, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	log.Print("Done!")
	return id
}

func GetAll() []User {
	stmt, err := database.Db.Prepare("select id, name from Users")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

package users

import (
	"database/sql"
	"log"

	database "back_go/internal/pkg/db/mysql"

	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, _ := HashPassword(user.Password)
	res, err := stmt.Exec(user.Name, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	return id
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Authenticate(username, claimedPassword string) bool {
	stmt, err := database.Db.Prepare("select Password from Users where Name = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(username)
	var hashedPassword string

	err = row.Scan(&hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	return checkPasswordHash(claimedPassword, hashedPassword)
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

func GetUserByName(username string) (*User, error) {
	stmt, err := database.Db.Prepare("select ID from Users where Name = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(username)

	var user User
	err = row.Scan(&user.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return nil, err
	}

	user.Name = username

	return &user, nil
}

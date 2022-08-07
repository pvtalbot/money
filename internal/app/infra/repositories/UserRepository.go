package repositories

import (
	"strconv"

	"github.com/pvtalbot/money/app/domain/model"

	"database/sql"
	"log"
)

type UserMariaRepository struct {
	db *sql.DB
}

func NewUserMariaRepository(db *sql.DB) UserMariaRepository {
	return UserMariaRepository{
		db: db,
	}
}

func (r UserMariaRepository) Create(user *model.User) (*model.User, error) {
	stmt, err := r.db.Prepare("insert into users(name, password, first_name, last_name) values (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(user.Name, user.GetHashedPassword(), user.FirstName, user.LastName)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
		return nil, err
	}

	user.ID = strconv.FormatInt(id, 10)

	return user, nil
}

func (r UserMariaRepository) FindAll() []*model.User {
	stmt, err := r.db.Prepare("select id, name, first_name, last_name from users")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.FirstName, &user.LastName)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func (r UserMariaRepository) FindByName(username string) (*model.User, error) {
	stmt, err := r.db.Prepare("select id, first_name, last_name from users where Name = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(username)

	var user model.User
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return nil, err
	}

	user.Name = username

	return &user, nil
}

func (r UserMariaRepository) FindPasswordByName(username string) (string, error) {
	stmt, err := r.db.Prepare("select password from users where Name = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return "", err
	}

	return hashedPassword, nil
}

package repositories

import (
	"strconv"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"github.com/pvtalbot/money/domain/models"
	"github.com/pvtalbot/money/errors"

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

func (r UserMariaRepository) Create(user *models.User) (*models.User, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO users(name, password, first_name, last_name) 
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(user.Name, user.GetHashedPassword(), user.FirstName, user.LastName)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == mysqlerr.ER_DUP_ENTRY {
				return nil, errors.DuplicateEntityError{}
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = strconv.FormatInt(id, 10)

	return user, nil
}

func (r UserMariaRepository) FindByName(username string) (*models.User, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, first_name, last_name 
		FROM users 
		WHERE name = ?
	`)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(username)

	var user models.User
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
	stmt, err := r.db.Prepare(`
		SELECT password
		FROM users 
		WHERE name = ?
	`)
	if err != nil {
		return "", err
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

func (r UserMariaRepository) Find(id string) (*models.User, error) {
	stmt, err := r.db.Prepare(`
		SELECT name, first_name, last_name 
		FROM users 
		WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var user models.User
	err = row.Scan(&user.Name, &user.FirstName, &user.LastName)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return nil, err
	}

	user.ID = id

	return &user, nil
}

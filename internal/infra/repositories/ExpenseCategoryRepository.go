package repositories

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/pvtalbot/money/domain/models"
)

type ExpenseCategoryMariaRepository struct {
	db *sql.DB
}

func NewExpenseCategoryMariaRepository(db *sql.DB) ExpenseCategoryMariaRepository {
	return ExpenseCategoryMariaRepository{
		db: db,
	}
}

func (r ExpenseCategoryMariaRepository) Create(ec *models.ExpenseCategory) (*models.ExpenseCategory, error) {
	stmt, err := r.db.Prepare("INSERT INTO expenses_categories(name, user_id) values (?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(ec.Name, ec.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
		return nil, err
	}

	ec.ID = strconv.FormatInt(id, 10)

	return ec, nil
}

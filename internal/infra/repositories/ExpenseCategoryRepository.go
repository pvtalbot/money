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

func (r ExpenseCategoryMariaRepository) FindAll(user *models.User) ([]*models.ExpenseCategory, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, name
		FROM expenses_categories
		WHERE user_id = ?
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []*models.ExpenseCategory
	for rows.Next() {
		var ec models.ExpenseCategory
		err := rows.Scan(&ec.ID, &ec.Name)

		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &ec)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}

func (r ExpenseCategoryMariaRepository) Find(id string) (*models.ExpenseCategory, error) {
	stmt, err := r.db.Prepare(`
		SELECT name, user_id
		FROM expenses_categories
		WHERE id = ?	
	`)
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)

	var userId string
	var expenseCategory models.ExpenseCategory
	err = row.Scan(&expenseCategory.Name, &userId)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return nil, err
	}

	expenseCategory.User = &models.User{ID: userId}
	expenseCategory.ID = id

	return &expenseCategory, nil
}

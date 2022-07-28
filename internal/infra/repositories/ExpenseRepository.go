package repositories

import (
	"back_go/internal/domain/model"
	"database/sql"
	"log"
	"strconv"
)

type ExpenseMariaRepository struct {
	db *sql.DB
}

func NewExpenseMariaRepository(db *sql.DB) ExpenseMariaRepository {
	return ExpenseMariaRepository{
		db: db,
	}
}

func (r ExpenseMariaRepository) GetAllExpensesFromUser(user *model.User) []*model.Expense {
	stmt, err := r.db.Prepare(`
		SELECT expenses.id, expenses.amount
		FROM expenses
		INNER JOIN users on users.id = expenses.user_id
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

	var expenses []*model.Expense
	for rows.Next() {
		var expense model.Expense
		err := rows.Scan(&expense.ID, &expense.Amount)
		if err != nil {
			log.Fatal(err)
		}
		expenses = append(expenses, &expense)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return expenses
}

func (r ExpenseMariaRepository) Create(expense *model.Expense, user *model.User) *model.Expense {
	stmt, err := r.db.Prepare("insert into expenses(amount, user_id) values (?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(expense.Amount, user.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	expense.ID = strconv.FormatInt(id, 10)

	return expense
}

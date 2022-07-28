package repositories

import (
	"back_go/internal/domain/model"
	"database/sql"
	"log"
	"strconv"
	"time"
)

type ExpenseMariaRepository struct {
	db *sql.DB
}

func NewExpenseMariaRepository(db *sql.DB) ExpenseMariaRepository {
	return ExpenseMariaRepository{
		db: db,
	}
}

func (r ExpenseMariaRepository) GetAllExpensesFromUserBetweenDates(user *model.User, startDate, endDate time.Time) []*model.Expense {
	stmt, err := r.db.Prepare(`
		SELECT expenses.id, expenses.amount, expenses.date
		FROM expenses
		INNER JOIN users on users.id = expenses.user_id
		WHERE user_id = ?
		AND expenses.date > ?
		AND expenses.date < ?
	`)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.ID, startDate, endDate)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var expenses []*model.Expense
	for rows.Next() {
		var expense model.Expense
		err := rows.Scan(&expense.ID, &expense.Amount, &expense.Date)
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
	stmt, err := r.db.Prepare("insert into expenses(amount, date, user_id) values (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(expense.Amount, expense.Date, user.ID)
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

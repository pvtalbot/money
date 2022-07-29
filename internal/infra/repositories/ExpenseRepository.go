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
		SELECT id, amount, date
		FROM expenses
		WHERE user_id = ?
		AND date > ?
		AND date < ?
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

func (r ExpenseMariaRepository) SumAllExpensesFromUserBetweenDatesByMonth(user *model.User, startDate, endDate time.Time) []*model.ExpenseSum {
	stmt, err := r.db.Prepare(`
		SELECT SUM(amount), MONTH(date), YEAR(date)
		FROM expenses
		WHERE user_id = ?
		AND date > ?
		AND date < ?
		GROUP BY MONTH(date), YEAR(date)
		ORDER BY YEAR(date), MONTH(date)
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

	var expensesSum []*model.ExpenseSum
	for rows.Next() {
		var amount, month, year int
		err := rows.Scan(&amount, &month, &year)
		if err != nil {
			log.Fatal(err)
		}

		expenseSum := model.ExpenseSum{
			Amount:    amount,
			StartDate: time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location()),
			EndDate:   time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Now().Location()),
		}

		expensesSum = append(expensesSum, &expenseSum)
	}

	return expensesSum
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

func (r ExpenseMariaRepository) Find(id int64) (*model.Expense, error) {
	stmt, err := r.db.Prepare("SELECT id, amount, date, user_id FROM expenses WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)

	var expense model.Expense
	var userId string
	err = row.Scan(&expense.ID, &expense.Amount, &expense.Date, &userId)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return nil, err
	}

	expense.User = model.User{ID: userId}

	return &expense, nil
}

func (r ExpenseMariaRepository) Delete(id int64) error {
	stmt, err := r.db.Prepare("DELETE FROM expenses WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

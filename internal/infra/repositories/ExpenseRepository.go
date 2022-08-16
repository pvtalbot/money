package repositories

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/pvtalbot/money/domain/models"
)

type ExpenseMariaRepository struct {
	db *sql.DB
}

func NewExpenseMariaRepository(db *sql.DB) ExpenseMariaRepository {
	return ExpenseMariaRepository{
		db: db,
	}
}

func (r ExpenseMariaRepository) GetAllExpensesFromUserBetweenDates(userId string, startDate, endDate time.Time) ([]*models.Expense, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, amount, date, expense_category_id
		FROM expenses
		WHERE user_id = ?
		AND date > ?
		AND date < ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []*models.Expense
	for rows.Next() {
		var expense models.Expense
		var expenseDate time.Time
		var expenseCategoryId string
		err := rows.Scan(&expense.ID, &expense.Amount, &expenseDate, &expenseCategoryId)
		if err != nil {
			return nil, err
		}
		expense.SetDate(expenseDate)
		expense.Category = models.ExpenseCategory{ID: expenseCategoryId}
		expense.User = models.User{ID: userId}
		expenses = append(expenses, &expense)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r ExpenseMariaRepository) SumAllExpensesFromUserBetweenDatesByMonth(userId string, startDate, endDate time.Time) ([]*models.ExpenseSum, error) {
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
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expensesSum []*models.ExpenseSum
	for rows.Next() {
		var amount, month, year int
		err := rows.Scan(&amount, &month, &year)
		if err != nil {
			return nil, err
		}

		expenseSum := models.ExpenseSum{
			Amount:    amount,
			StartDate: time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location()),
			EndDate:   time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Now().Location()),
		}

		expensesSum = append(expensesSum, &expenseSum)
	}

	return expensesSum, nil
}

func (r ExpenseMariaRepository) Create(expense *models.Expense, userId, categoryId string) (*models.Expense, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO expenses(amount, date, user_id, expense_category_id) 
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(expense.Amount, expense.GetDate(), userId, categoryId)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	expense.ID = strconv.FormatInt(id, 10)

	return expense, nil
}

func (r ExpenseMariaRepository) Update(expense *models.Expense) (*models.Expense, error) {
	stmt, err := r.db.Prepare(`
		UPDATE expenses 
		SET amount = ?, date = ?, expense_category_id = ? 
		WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(expense.Amount, expense.GetDate(), expense.Category.ID, expense.ID)
	if err != nil {
		return nil, err
	}

	return r.Find(expense.ID)
}

func (r ExpenseMariaRepository) Find(id string) (*models.Expense, error) {
	stmt, err := r.db.Prepare(`
		SELECT amount, date, user_id, expense_category_id 
		FROM expenses 
		WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var expense models.Expense
	var expenseDate time.Time
	var userId string
	var categoryId string
	err = row.Scan(&expense.Amount, &expenseDate, &userId, &categoryId)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return nil, err
	}

	expense.SetDate(expenseDate)
	expense.ID = id
	expense.User = models.User{ID: userId}
	expense.Category = models.ExpenseCategory{ID: categoryId}

	return &expense, nil
}

func (r ExpenseMariaRepository) Delete(id string) error {
	stmt, err := r.db.Prepare(`
		DELETE FROM expenses 
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

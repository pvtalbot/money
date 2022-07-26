package repositories

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/pvtalbot/money/domain/models"
)

type RevenueMariaRepository struct {
	db *sql.DB
}

func NewRevenueMariaRepository(db *sql.DB) RevenueMariaRepository {
	return RevenueMariaRepository{
		db: db,
	}
}

func (r RevenueMariaRepository) GetAllRevenuesOfUserBetweenDates(userId string, startDate, endDate time.Time) ([]*models.Revenue, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, amount, date
		FROM revenues
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

	var revenues []*models.Revenue
	for rows.Next() {
		var revenue models.Revenue
		var revenueDate time.Time
		err := rows.Scan(&revenue.ID, &revenue.Amount, &revenueDate)
		if err != nil {
			return nil, err
		}
		revenue.SetDate(revenueDate)
		revenue.User = models.User{ID: userId}
		revenues = append(revenues, &revenue)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return revenues, nil
}

func (r RevenueMariaRepository) Create(revenue *models.Revenue, userId string) (*models.Revenue, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO revenues(amount, date, user_id)
		VALUES (?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(revenue.Amount, revenue.GetDate(), userId)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	revenue.ID = strconv.FormatInt(id, 10)

	return revenue, nil
}

func (r RevenueMariaRepository) Update(revenue *models.Revenue) (*models.Revenue, error) {
	log.Println("revenue repository", revenue.ID, revenue.Amount, revenue.GetDate())
	stmt, err := r.db.Prepare(`
		UPDATE revenues
		SET amount = ?, date = ?
		WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(revenue.Amount, revenue.GetDate(), revenue.ID)
	if err != nil {
		return nil, err
	}

	return r.Find(revenue.ID)
}

func (r RevenueMariaRepository) Find(id string) (*models.Revenue, error) {
	stmt, err := r.db.Prepare(`
		SELECT amount, date, user_id
		FROM revenues
		WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var revenue models.Revenue
	var revenueDate time.Time
	var userId string
	err = row.Scan(&revenue.Amount, &revenueDate, &userId)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return nil, err
	}
	revenue.ID = id
	revenue.SetDate(revenueDate)
	revenue.User = models.User{ID: userId}

	return &revenue, nil
}

func (r RevenueMariaRepository) Delete(id string) error {
	stmt, err := r.db.Prepare(`
		DELETE FROM revenues
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

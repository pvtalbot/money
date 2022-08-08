package repositories

import (
	"database/sql"

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

func (r ExpenseCategoryMariaRepository) Create(ec *models.ExpenseCategoryRepository, user *models.User)

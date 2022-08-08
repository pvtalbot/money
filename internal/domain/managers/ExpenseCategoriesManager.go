package managers

import "github.com/pvtalbot/money/domain/models"

func GetDefaultCategories(u *models.User) []models.ExpenseCategory {
	names := []string{"Basic", "Pleasure"}

	var result []models.ExpenseCategory
	for _, v := range names {
		result = append(result, models.ExpenseCategory{Name: v, User: u})
	}

	return result
}

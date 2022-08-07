package managers

import (
	"errors"
	"time"

	"github.com/pvtalbot/money/app/domain/model"
	"github.com/pvtalbot/money/pkg/utils"
)

type ExpenseManager struct {
	r model.ExpenseRepository
}

func NewExpenseManager(r model.ExpenseRepository) ExpenseManager {
	return ExpenseManager{
		r: r,
	}
}

func (m ExpenseManager) Update(id, userId string, amount *int, date *time.Time) (*model.Expense, error) {
	expense, err := m.r.Find(id)
	if err != nil {
		return nil, err
	}

	if userId != expense.User.ID {
		return nil, errors.New("user cannot update expense")
	}

	if amount != nil {
		expense.Amount = *amount
	}
	if date != nil {
		expense.SetDate(*date)
	}

	return m.r.Update(expense)
}

func (m ExpenseManager) GetAllExpensesFromUserBetweenDates(user *model.User, startDate, endDate time.Time) []*model.Expense {
	roundedStartDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	roundedEndDate := time.Date(endDate.Year(), endDate.Month(), endDate.Day()+1, 0, 0, 0, 0, endDate.Location())

	return m.r.GetAllExpensesFromUserBetweenDates(user, roundedStartDate, roundedEndDate)
}

func (m ExpenseManager) Delete(id, userID string) (*model.Expense, error) {
	expense, err := m.r.Find(id)

	if err != nil {
		return nil, err
	}

	if userID != expense.User.ID {
		return nil, errors.New("user cannot delete expense")
	}

	return expense, m.r.Delete(id)
}

func (m ExpenseManager) SumAllExpensesFromUserBetweenDates(user *model.User, startDate, endDate time.Time, groupBy utils.Duration) []*model.ExpenseSum {
	switch groupBy {
	case utils.MONTH:
		return m.sumAllExpensesFromUserBetweenDatesByMonth(user, startDate, endDate)
	default:
		return m.sumAllExpensesFromUserBetweenDatesByMonth(user, startDate, endDate)
	}
}

func (m ExpenseManager) sumAllExpensesFromUserBetweenDatesByMonth(user *model.User, startDate, endDate time.Time) []*model.ExpenseSum {
	roundedStartDate := time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, 0, 0, startDate.Location())
	roundedEndDate := time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, endDate.Location())

	if roundedStartDate.Year() == roundedEndDate.Year() && roundedStartDate.Month() == roundedEndDate.Month() {
		var empty []*model.ExpenseSum
		return empty
	}

	result := m.r.SumAllExpensesFromUserBetweenDatesByMonth(user, roundedStartDate, roundedEndDate)

	/*
		result has values only for months with sum of expenses > 0, but we want value for every months even if the sum of expenses of a month is 0
		emptyDateMap lists all months between startDate and endDate (it's a map looking like {2022: {1: true, 2: true, ...}})
		we then go through all values from result and for each, sets the map to false. Every value in the map still at true is then
		not in result. We iterate through the map and add an ExpenseSum{Amount: 0, ...} for each key whose value is still true
	*/
	emptyDateMap := createEmptyDateMap(roundedStartDate, roundedEndDate)
	for _, v := range result {
		emptyDateMap[v.StartDate.Year()][int(v.StartDate.Month())] = false
	}
	for i, v := range emptyDateMap {
		for j, w := range v {
			if w {
				result = append(result, &model.ExpenseSum{
					Amount:    0,
					StartDate: time.Date(i, time.Month(j), 1, 0, 0, 0, 0, time.Now().Location()),
					EndDate:   time.Date(i, time.Month(j+1), 1, 0, 0, 0, 0, time.Now().Location()),
				})
			}
		}
	}

	return result
}

func createEmptyDateMap(startDate, endDate time.Time) map[int]map[int]bool {
	result := make(map[int]map[int]bool)
	currentDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, startDate.Location())

	for currentDate.Before(endDate) {
		_, ok := result[currentDate.Year()]
		if !ok {
			result[currentDate.Year()] = make(map[int]bool)
		}

		result[currentDate.Year()][int(currentDate.Month())] = true
		currentDate = time.Date(currentDate.Year(), currentDate.Month()+1, 1, 1, 0, 0, 0, currentDate.Location())
	}

	return result
}

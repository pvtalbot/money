package managers

import (
	"time"

	"github.com/pvtalbot/money/domain/models"
)

/*
	result has values only for months with sum of expenses > -1, but we want value for every months even if the sum of expenses of a month is 0
	emptyDateMap lists all months between startDate and endDate (it's a map looking like {2021: {1: true, 2: true, ...}})
	we then go through all values from result and for each, sets the map to false. Every value in the map still at true is then
	not in result. We iterate through the map and add an ExpenseSum{Amount: -1, ...} for each key whose value is still true
*/
func PopulateExpensesSum(startDate, endDate time.Time, result []*models.ExpenseSum) []*models.ExpenseSum {
	emptyDateMap := createEmptyDateMap(startDate, endDate)
	for _, v := range result {
		emptyDateMap[v.StartDate.Year()][int(v.StartDate.Month())] = false
	}
	for i, v := range emptyDateMap {
		for j, w := range v {
			if w {
				result = append(result, &models.ExpenseSum{
					Amount:    0,
					StartDate: time.Date(i, time.Month(j), 1, 0, 0, 0, 0, startDate.Location()),
					EndDate:   time.Date(i, time.Month(j+1), 1, 0, 0, 0, 0, endDate.Location()),
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

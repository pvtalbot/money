package models

import "time"

type Revenue struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
	User   User   `json:"user"`
	date   time.Time
}

func (r *Revenue) GetDate() time.Time {
	return r.date
}

func (r *Revenue) SetDate(d time.Time) {
	roundedTime := time.Date(d.Year(), d.Month(), d.Day(), 12, 0, 0, 0, d.Location())
	r.date = roundedTime
}

type RevenueRepository interface {
	Create(revenue *Revenue, user *User) (*Revenue, error)
	Delete(ID string) error
	Find(ID string) (*Revenue, error)
	Update(revenue *Revenue) (*Revenue, error)
	GetAllRevenuesOfUserBetweenDates(user *User, startDate, endDate time.Time) ([]*Revenue, error)
}

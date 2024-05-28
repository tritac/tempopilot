package worklogs

import "time"

type WorkDay struct {
	Date         time.Time
	WeekDay      time.Weekday
	IsWorkingDay bool
}

type WorkLogStore struct {
	days []WorkDay
}

func NewWorkLogStore() *WorkLogStore {
	return &WorkLogStore{}
}

func (wl *WorkLogStore) GenerateMonthDays(year int, month time.Month) []WorkDay {

	workDayList := []WorkDay{}
	days := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC).Day()

	for i := 1; i <= days; i++ {
		date := time.Date(year, month, i, 12, 00, 00, 00, time.UTC)
		weekDay := date.Weekday()
		isWorkingDay := true
		if weekDay == time.Saturday || weekDay == time.Sunday {
			isWorkingDay = false
		}
		workday := WorkDay{Date: date, WeekDay: weekDay, IsWorkingDay: isWorkingDay}
		workDayList = append(workDayList, workday)

	}
	return workDayList

}

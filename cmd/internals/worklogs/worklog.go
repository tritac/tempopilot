package worklog

import "time"

type WorkDay struct {
	Date         time.Time `json:"date"`
	WeekDay      string    `json:"day"`
	IsWorkingDay bool      `json:"isWorking"`
}

type WorkLogStore struct {
	Mdays []WorkDay
	Today time.Time
}

func NewWorkLogStore() *WorkLogStore {
	return &WorkLogStore{Today: time.Now()}
}

func (wl *WorkLogStore) GenerateMothDays(year int, month time.Month) []WorkDay {

	workdayslist := []WorkDay{}

	days := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	for i := 1; i <= days; i++ {

		date := time.Date(year, month, i, 12, 00, 00, 00, time.UTC)
		weekDay := date.Weekday()

		isWorking := true
		if weekDay == time.Saturday || weekDay == time.Sunday {
			isWorking = false
		}
		workday := WorkDay{Date: date, WeekDay: weekDay.String(), IsWorkingDay: isWorking}
		workdayslist = append(workdayslist, workday)
	}
	return workdayslist

}

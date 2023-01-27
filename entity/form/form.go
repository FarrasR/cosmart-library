package form

import "time"

type FormCreateBook struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Edition int    `json:"edition"`
}

type FormCreateSchedule struct {
	Name       string    `json:"name"`
	BookId     int       `json:"book_id"`
	PickupTime time.Time `json:"pickup_time"`
}

type FormReturnBook struct {
	ScheduleId int       `json:"schedule_id"`
	ReturnTime time.Time `json:"return_time"`
}

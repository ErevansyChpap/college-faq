package models

type Service struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Room             *string `json:"room,omitempty"`
	Floor            *int    `json:"floor,omitempty"`
	Phone            *string `json:"phone,omitempty"`
	Manager          *string `json:"manager,omitempty"`
	ScheduleWeekdays *string `json:"schedule_weekdays,omitempty"`
	ScheduleSaturday *string `json:"schedule_saturday,omitempty"`
}

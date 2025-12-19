package model

import "time"

type Schedule struct {
	ID        int       `db:"id"`
	GroupName string    `db:"group_name"`
	DayOfWeek string    `db:"day_of_week"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
	Subject   string    `db:"subject"`
	Room      string    `db:"room"`
	Teacher   string    `db:"teacher"`
	CreatedAt time.Time `db:"created_at"`
}

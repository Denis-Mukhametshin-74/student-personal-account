package repository

import (
	"database/sql"
	"fmt"
	"student-portal/internal/app/model"
)

type ScheduleRepository struct {
	db *sql.DB
}

func NewScheduleRepository(db *sql.DB) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

func (r *ScheduleRepository) FindByGroup(groupName string) ([]model.Schedule, error) {
	query := `
        SELECT id, group_name, day_of_week, start_time, end_time, subject, room, teacher, created_at
        FROM schedules 
        WHERE group_name = $1 
        ORDER BY 
            CASE day_of_week
                WHEN 'Понедельник' THEN 1
                WHEN 'Вторник' THEN 2
                WHEN 'Среда' THEN 3
                WHEN 'Четверг' THEN 4
                WHEN 'Пятница' THEN 5
                WHEN 'Суббота' THEN 6
                WHEN 'Воскресенье' THEN 7
            END, start_time
    `

	rows, err := r.db.Query(query, groupName)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса расписания: %w", err)
	}
	defer rows.Close()

	var schedules []model.Schedule
	for rows.Next() {
		var s model.Schedule
		if err := rows.Scan(
			&s.ID, &s.GroupName, &s.DayOfWeek, &s.StartTime, &s.EndTime,
			&s.Subject, &s.Room, &s.Teacher, &s.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("ошибка сканирования расписания: %w", err)
		}
		schedules = append(schedules, s)
	}

	return schedules, nil
}

package service

import (
	"errors"
	"fmt"
	"student-portal/internal/app/dto"
	"student-portal/internal/app/repository"
)

type ScheduleService struct {
	scheduleRepo *repository.ScheduleRepository
	studentRepo  *repository.StudentRepository
}

func NewScheduleService(scheduleRepo *repository.ScheduleRepository,
	studentRepo *repository.StudentRepository) *ScheduleService {
	return &ScheduleService{
		scheduleRepo: scheduleRepo,
		studentRepo:  studentRepo,
	}
}

func (s *ScheduleService) GetSchedule(email string) (*dto.ScheduleResponse, error) {
	student, err := s.studentRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if student == nil {
		return nil, errors.New("пользователь не найден")
	}

	groupName := student.GroupName
	schedules, err := s.scheduleRepo.FindByGroup(groupName)
	if err != nil {
		return nil, err
	}

	if schedules == nil {
		return nil, errors.New("расписание не найдено")
	}

	scheduleMap := make(map[string][]dto.ScheduleItem)
	for _, schedule := range schedules {
		timeStr := fmt.Sprintf(
			"%s-%s",
			schedule.StartTime.Format("15:04"),
			schedule.EndTime.Format("15:04"),
		)

		item := dto.ScheduleItem{
			Time:    timeStr,
			Subject: schedule.Subject,
			Room:    schedule.Room,
			Teacher: schedule.Teacher,
		}

		scheduleMap[schedule.DayOfWeek] = append(
			scheduleMap[schedule.DayOfWeek], item,
		)
	}

	var days []dto.ScheduleDay
	dayOrder := []string{
		"Понедельник", "Вторник", "Среда",
		"Четверг", "Пятница", "Суббота",
	}

	for _, day := range dayOrder {
		if lessons, ok := scheduleMap[day]; ok && len(lessons) > 0 {
			days = append(days, dto.ScheduleDay{
				Day:     day,
				Lessons: lessons,
			})
		}
	}

	return &dto.ScheduleResponse{
		Group: groupName,
		Days:  days,
	}, nil
}

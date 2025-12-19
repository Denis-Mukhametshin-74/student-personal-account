package service

import (
	"errors"
	"student-portal/internal/app/dto"
	"student-portal/internal/app/repository"
)

type ProfileService struct {
	studentRepo *repository.StudentRepository
	subjectRepo *repository.SubjectRepository
}

func NewProfileService(studentRepo *repository.StudentRepository,
	subjectRepo *repository.SubjectRepository) *ProfileService {
	return &ProfileService{
		studentRepo: studentRepo,
		subjectRepo: subjectRepo,
	}
}

func (s *ProfileService) GetProfile(email string) (*dto.ProfileResponse, error) {
	student, err := s.studentRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if student == nil {
		return nil, errors.New("пользователь не найден")
	}

	subjects, err := s.subjectRepo.FindByStudentID(student.ID)
	if err != nil {
		return nil, err
	}

	if subjects == nil {
		return nil, errors.New("предметы не найдены")
	}

	studentResponse := dto.StudentResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email,
		Group: student.GroupName,
	}

	var subjectsResponse []dto.SubjectResponse
	for _, subject := range subjects {
		subjectsResponse = append(subjectsResponse, dto.SubjectResponse{
			ID:      subject.ID,
			Name:    subject.Name,
			Teacher: subject.Teacher,
		})
	}

	return &dto.ProfileResponse{
		Student:  studentResponse,
		Subjects: subjectsResponse,
	}, nil
}

package service

import (
	"errors"
	"student-portal/internal/app/dto"
	"student-portal/internal/app/model"
	"student-portal/internal/app/repository"
	"student-portal/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	studentRepo *repository.StudentRepository
}

func NewAuthService(studentRepo *repository.StudentRepository) *AuthService {
	return &AuthService{studentRepo: studentRepo}
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	student, err := s.studentRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if student == nil {
		return nil, errors.New("пользователь не найден")
	}

	err = bcrypt.CompareHashAndPassword([]byte(student.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("неверный пароль")
	}

	token, err := jwt.GenerateToken(student.ID, student.Email)
	if err != nil {
		return nil, errors.New("ошибка генерации токена")
	}

	return &dto.LoginResponse{
		Token: token,
		Student: dto.StudentResponse{
			ID:    student.ID,
			Name:  student.Name,
			Email: student.Email,
			Group: student.GroupName,
		},
	}, nil
}

func (s *AuthService) Register(req dto.RegisterRequest) (*dto.StudentResponse, error) {
	exists, err := s.studentRepo.Exists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("пользователь с таким email уже существует")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("ошибка хеширования пароля")
	}

	student := &model.Student{
		Name:         req.Name,
		Email:        req.Email,
		GroupName:    req.Group,
		PasswordHash: string(hashedPassword),
	}

	err = s.studentRepo.Create(student)
	if err != nil {
		return nil, err
	}

	return &dto.StudentResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email,
		Group: student.GroupName,
	}, nil
}

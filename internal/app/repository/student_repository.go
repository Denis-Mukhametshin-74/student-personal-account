package repository

import (
	"database/sql"
	"fmt"
	"student-portal/internal/app/model"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) FindByEmail(email string) (*model.Student, error) {
	query := `
        SELECT id, name, email, group_name, password_hash, created_at, updated_at 
        FROM students WHERE email = $1
    `

	var student model.Student
	err := r.db.QueryRow(query, email).Scan(
		&student.ID, &student.Name, &student.Email, &student.GroupName,
		&student.PasswordHash, &student.CreatedAt, &student.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("ошибка поиска пользователя: %w", err)
	}

	return &student, nil
}

func (r *StudentRepository) Create(student *model.Student) error {
	query := `
        INSERT INTO students (name, email, group_name, password_hash) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id, created_at, updated_at
    `

	err := r.db.QueryRow(
		query, student.Name, student.Email, student.GroupName, student.PasswordHash,
	).Scan(&student.ID, &student.CreatedAt, &student.UpdatedAt)

	if err != nil {
		return fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	return nil
}

func (r *StudentRepository) Exists(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM students WHERE email = $1)`
	var exists bool
	err := r.db.QueryRow(query, email).Scan(&exists)
	return exists, err
}

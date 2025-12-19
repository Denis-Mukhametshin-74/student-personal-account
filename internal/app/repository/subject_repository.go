package repository

import (
	"database/sql"
	"fmt"
	"student-personal-account/internal/app/model"
)

type SubjectRepository struct {
	db *sql.DB
}

func NewSubjectRepository(db *sql.DB) *SubjectRepository {
	return &SubjectRepository{db: db}
}

func (r *SubjectRepository) FindByStudentID(studentID int) ([]model.Subject, error) {
	query := `
        SELECT s.id, s.name, s.teacher
        FROM subjects s
        JOIN student_subjects ss ON s.id = ss.subject_id
        WHERE ss.student_id = $1
        ORDER BY s.name
    `

	rows, err := r.db.Query(query, studentID)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса предметов: %w", err)
	}
	defer rows.Close()

	var subjects []model.Subject
	for rows.Next() {
		var s model.Subject
		if err := rows.Scan(&s.ID, &s.Name, &s.Teacher); err != nil {
			return nil, fmt.Errorf("ошибка сканирования предмета: %w", err)
		}
		subjects = append(subjects, s)
	}

	return subjects, nil
}

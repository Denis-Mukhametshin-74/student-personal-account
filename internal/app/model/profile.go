package model

type Profile struct {
	StudentID int `db:"student_id"`
	SubjectID int `db:"subject_id"`
}

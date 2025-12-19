package dto

type StudentResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Group string `json:"group"`
}

type SubjectResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Teacher string `json:"teacher"`
}

type ProfileResponse struct {
	Student  StudentResponse   `json:"student"`
	Subjects []SubjectResponse `json:"subjects,omitempty"`
}

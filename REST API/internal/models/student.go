package models

type Student struct {
	ID        int    `json:"id,omitempty" db:"id,omitempty"`
	FirstName string `json:"first_name,omitempty" db:"FirstName,omitempty"`
	LastName  string `json:"last_name,omitempty" db:"LastName,omitempty"`
	Email     string `json:"email,omitempty" db:"email,omitempty"`
	Class     string `json:"class,omitempty" db:"class,omitempty"`
	TeacherID int    `json:"teacher_id,omitempty" db:"teacher_id,omitempty"`
}

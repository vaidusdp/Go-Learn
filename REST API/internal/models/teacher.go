package models

type Teacher struct {
	ID        int    `json:"id,omitempty" db:"id,omitempty"`
	FirstName string `json:"first_name,omitempty" db:"FirstName,omitempty"`
	LastName  string `json:"last_name,omitempty" db:"LastName,omitempty"`
	Email     string `json:"email,omitempty" db:"email,omitempty"`
	Subject   string `json:"subject,omitempty" db:"subject,omitempty"`
	Class     string `json:"class,omitempty" db:"class,omitempty"`
}

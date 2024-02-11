package models

type UserRole string

const (
	Admin   UserRole = "ADMIN"
	Regular UserRole = "REGULAR"
	Guest   UserRole = "GUEST"
)

type User struct {
	UUID      string   `json:"uuid"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Email     string   `json:"email"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Role      UserRole `json:"role"`
}

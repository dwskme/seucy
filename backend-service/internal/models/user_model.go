package models

import (
	"github.com/dwskme/seucy/backend-service/internal/utils/db"
	"github.com/dwskme/seucy/backend-service/internal/utils/uuid"
)

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
	Password  string   `json:"password"`
	Role      UserRole `json:"role"`
}

func CreateUser(user User) error {
	user.UUID = uuid.New().Generate()
	_, err := db.DB.Exec(
		"INSERT INTO users (uuid, firstname, lastname, email, password, role) VALUES ($1, $2, $3, $4, $5, $6)",
		user.UUID, user.Firstname, user.Lastname, user.Email, user.Password, user.Role,
	)
	return err
}

func GetUser(uuid string) (User, error) {
	var user User
	err := db.DB.QueryRow(
		"SELECT uuid, firstname, lastname, email, password, role FROM users WHERE uuid = $1",
		uuid,
	).Scan(&user.UUID, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Role)
	return user, err
}

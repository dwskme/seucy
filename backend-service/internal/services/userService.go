package services

import (
	"database/sql"

	models "github.com/dwskme/seucy/backend-service/internal/models"
	db "github.com/dwskme/seucy/backend-service/internal/utils/db"
	uuid "github.com/dwskme/seucy/backend-service/internal/utils/uuid"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) CreateUser(user models.User) error {
	user.UUID = uuid.New().Generate()
	_, err := db.DB.Exec(
		"INSERT INTO users (uuid, firstname, lastname, email, password, role,username) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		user.UUID, user.Firstname, user.Lastname, user.Email, user.Password, user.Role, user.Username,
	)
	return err
}

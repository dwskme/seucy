package services

import (
	"database/sql"

	"github.com/dwskme/seucy/backend-service/internal/models"
)

type MediaService struct {
	DB *sql.DB
}

func NewMediaService(db *sql.DB) *MediaService {
	return &MediaService{DB: db}
}

// TODO: better way to get paramerter to check media from url only
func (ms *MediaService) GetUserMedia(userid, mediatype string) ([]models.UserPreference, error) {
	var result []models.UserPreference
	query := "SELECT * FROM userpreferences WHERE userid = $1 AND mediatype= $2"
	rows, err := ms.DB.Query(query, userid, mediatype)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var up models.UserPreference
		if err := rows.Scan(
			&up.UserPreferenceID,
			&up.UserID,
			&up.MediaID,
			&up.MediaType,
			&up.Rating,
			&up.Review); err != nil {
			return nil, err
		}
		result = append(result, up)
	}
	return result, nil
}

func (ms *MediaService) AddUserMedia(up models.UserPreference) error {
	query := "INSERT INTO UserPreferences (userid, mediaid, mediatype, rating, review,userpreferenceid ) VALUES ($1, $2, $3, $4, $5, $6 )"
	_, err := ms.DB.Exec(
		query,
		up.UserID,
		up.MediaID,
		up.MediaType,
		up.Rating,
		up.Review,
		up.UserPreferenceID)
	if err != nil {
		return err
	}
	return nil
}

// TODO: get where param form url only
func (ms *MediaService) UpdateUserMedia(up models.UserPreference) error {
	query := "UPDATE UserPreferences SET rating=$1, review=$2 WHERE userpreferenceid=$3"
	_, err := ms.DB.Exec(query,
		up.Rating,
		up.Review,
		up.UserPreferenceID)
	if err != nil {
		return err
	}
	return nil
}

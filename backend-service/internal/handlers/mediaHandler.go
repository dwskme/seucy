package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dwskme/seucy/backend-service/internal/models"
)

type RequestData struct {
	UserID    string `json:"userid"`
	MediaType string `json:"mediatype"`
}

// TODO:validation and better handling or error return err msg and so on
// Todo: better way of mapping in body
// TODO: add validator for request fields
func (h *NewHandler) AddPreference(w http.ResponseWriter, r *http.Request) {
	var up models.UserPreference
	if err := json.NewDecoder(r.Body).Decode(&up); err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}
	err := h.MediaService.AddUserMedia(up)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	JsonResponse(w, http.StatusCreated, "Created successfully", nil)
}

func (h *NewHandler) UpdatePreference(w http.ResponseWriter, r *http.Request) {
	var up models.UserPreference
	if err := json.NewDecoder(r.Body).Decode(&up); err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}
	err := h.MediaService.UpdateUserMedia(up)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	JsonResponse(w, http.StatusNoContent, "", nil)
}

func (h *NewHandler) GetPreferences(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		JsonResponse(w, http.StatusBadRequest, "Invalid JSON request", nil)
		return
	}
	userPreferences, err := h.MediaService.GetUserMedia(requestData.UserID, requestData.MediaType)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	JsonResponse(w, http.StatusOK, "User media retrieved successfully", userPreferences)
}

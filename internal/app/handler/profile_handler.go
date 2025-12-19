package handler

import (
	"encoding/json"
	"net/http"

	"student-personal-account/internal/app/service"
)

type ProfileHandler struct {
	profileService *service.ProfileService
}

func NewProfileHandler(profileService *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email, ok := ctx.Value("email").(string)
	if !ok || email == "" {
		http.Error(w, "не удалось получить данные пользователя", http.StatusUnauthorized)
		return
	}

	response, err := h.profileService.GetProfile(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

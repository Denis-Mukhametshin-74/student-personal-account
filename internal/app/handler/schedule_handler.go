package handler

import (
	"encoding/json"
	"net/http"

	"student-personal-account/internal/app/service"
)

type ScheduleHandler struct {
	scheduleService *service.ScheduleService
}

func NewScheduleHandler(scheduleService *service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{scheduleService: scheduleService}
}

func (h *ScheduleHandler) GetSchedule(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email, ok := ctx.Value("email").(string)
	if !ok || email == "" {
		http.Error(w, "не удалось получить данные пользователя", http.StatusUnauthorized)
		return
	}

	schedule, err := h.scheduleService.GetSchedule(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedule)
}
